package typelang

import (
	"errors"
	"fmt"
	"github.com/xelaj/tl"
	"github.com/xelaj/tl/parser/typelang/declaration"
	"github.com/xelaj/tl/parser/typelang/lexer"
	"io"
	"io/fs"
	"os"
	"strconv"
	"strings"

	"github.com/alecthomas/participle/v2"
)

//nolint:gochecknoglobals // obviously parser must be global
var parser = participle.MustBuild[declaration.Program](
	participle.Lexer(lexer.NewDefinition()),
)

func ParseFile(filename string) (*tl.Schema, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Parse(filename, f)
}

func ParseFS(fsys fs.FS, filename string) (*tl.Schema, error) {
	f, err := fsys.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Parse(filename, f)
}

func ParseString(filename string, content string) (*tl.Schema, error) {
	return Parse(filename, strings.NewReader(content))
}

func Parse(filename string, content io.Reader) (*tl.Schema, error) {
	res, err := parser.Parse(filename, content)
	if err != nil {
		return nil, err //nolint:wrapcheck // it's important to keep error
	}

	normalized, err := normalizeProgram(res)

	normalized.Layer = 0

	if err != nil {
		return nil, err
	}

	return normalized, nil
}

func normalizeType(i *declaration.Type) (tl.Type, error) {
	name := tl.ParseNameFromString(i.Ident.String())
	if len(i.SubTypes) == 0 {
		return tl.Type{Name: name}, nil
	}
	if len(i.SubTypes) == 1 {
		typ, err := normalizeType(&i.SubTypes[0])
		if err != nil {
			return typ, err
		}
		return tl.Type{Name: name, SubType: &typ}, nil
	}
	return tl.Type{}, errors.New("too many subtypes")
}

func normalizeArgument(arg *declaration.Argument, comment string) (tl.Param, error) {
	typ, err := normalizeType(&arg.Type)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", arg.Ident.String(), err)
	}

	if arg.Flag == nil {
		if typ.Key == "#" {
			return tl.BitflagParam{
				Comment: comment,
				Name:    arg.Ident.String(),
			}, nil
		}

		return tl.RequiredParam{
			Comment: comment,
			Name:    arg.Ident.String(),
			Type:    typ,
		}, nil
	}

	if typ.Key == "true" {
		return tl.TriggerParam{
			Comment:     comment,
			Name:        arg.Ident.String(),
			FlagTrigger: arg.Flag.Ident,
			BitTrigger:  arg.Flag.Index,
		}, nil
	}

	return tl.OptionalParam{
		Comment:     comment,
		Name:        arg.Ident.String(),
		Type:        typ,
		FlagTrigger: arg.Flag.Ident,
		BitTrigger:  arg.Flag.Index,
	}, nil
}

func normalizeDeclaration(decl *declaration.Declaration, constructorComment string, argsComments map[string]string) (*tl.Declaration, error) {
	parts := strings.Split(decl.Combinator, "#") // guaranteed to split by two parts, lexer handles it
	if len(parts) == 0 || len(parts) > 2 {
		return nil, errors.New(decl.Combinator + ": invalid combinator")
	}
	name := tl.ParseNameFromString(parts[0])
	crc := uint64(0)
	if len(parts) == 2 {
		crc1, err := strconv.ParseUint(parts[1], 16, 32)
		if err != nil {
			return nil, fmt.Errorf("%v: %w", decl.Combinator, err)
		}
		crc = crc1
	}

	params := make([]tl.Param, len(decl.Args))
	for i, arg := range decl.Args {
		arg := arg

		var comment string
		if !arg.Ident.Empty() {
			comment = argsComments[arg.Ident.String()]
			delete(argsComments, arg.Ident.String())
		}

		var argErr error
		params[i], argErr = normalizeArgument(&arg, comment)
		if argErr != nil {
			return nil, fmt.Errorf("%v: %w", decl.Combinator, argErr)
		}
	}

	optParams := make(tl.Params, len(decl.OptArgs))

	for i, arg := range decl.OptArgs {
		arg := arg
		var comment string
		if !arg.Ident.Empty() {
			if !arg.Ident.Empty() {
				comment = argsComments[arg.Ident.String()]
				delete(argsComments, arg.Ident.String())
			}
		}

		var argErr error
		optParams[i], argErr = normalizeArgument(&arg, comment)
		if argErr != nil {
			return nil, fmt.Errorf("%v: %w", decl.Combinator, argErr)
		}
	}

	typ, err := normalizeType(&declaration.Type{
		Ident:    decl.Result.Ident,
		SubTypes: decl.Result.SubTypes,
	})
	if err != nil {
		return nil, fmt.Errorf(decl.Combinator+": parsing return type: %w", err)
	}

	if len(argsComments) != 0 {
		keys := make([]string, 0, len(argsComments))
		for k := range argsComments {
			keys = append(keys, k)
		}
		keysStr := strings.Join(keys, ", ")
		return nil, fmt.Errorf("%v: unknown params in comment tags: %v", decl.Combinator, keysStr)
	}

	return &tl.Declaration{
		Comment:   constructorComment,
		Name:      name,
		CRC:       uint32(crc),
		Params:    params,
		OptParams: optParams,
		Type:      typ,
	}, nil
}

const (
	tagType    = "type"
	tagPredict = "predict"
	tagMethod  = "method"
	tagParam   = "param"
)

func normalizeEntries(items []declaration.ProgramEntry) (typeDecls []*tl.Declaration, funcDecls []*tl.Declaration, typeComments map[tl.Name]string, err error) {
	var (
		currentTypeComment     string
		constructorComment     string
		constructorCommentType string
		argumentComments       = map[string]string{}
		funcMode               = false
	)
	for _, item := range items {
		switch {
		case item.TypeDecl:
			funcMode = false
		case item.FuncDecl:
			funcMode = true
		case item.Newline:
			constructorComment = ""
			constructorCommentType = ""
			argumentComments = map[string]string{}

		case item.Comment != nil:
			comment := strings.TrimSpace(strings.TrimPrefix(*item.Comment, "//"))
			var commentTag string
			if strings.HasPrefix(comment, "@") {
				const commentParts = 2

				comment = strings.TrimPrefix(comment, "@")
				parts := strings.SplitN(comment, " ", commentParts)

				commentTag = parts[0]
				if len(parts) >= commentParts {
					comment = strings.TrimSpace(parts[1])
				} else {
					comment = ""
				}
			}

			switch tag := commentTag; tag {
			case "":
				// pass

			case tagType:
				if funcMode {
					err := fmt.Errorf("@%v %v: impossible on functions", tag, comment)
					return nil, nil, nil, err
				}

				// special case, declaration of type apply to first available type (= SomeType;).
				// when type comment applied to type, currentTypeComment is clearing, so we checking
				// is this type comment declared twice.
				if currentTypeComment != "" {
					err := fmt.Errorf("@%v %v: type comment declared twice", tag, comment)
					return nil, nil, nil, err
				}
				currentTypeComment = comment

			case tagPredict, tagMethod:
				if funcMode && tag != tagMethod {
					return nil, nil, nil, errors.New("@" + tag + ": works only in type definitions")
				} else if !funcMode && comment == tagMethod {
					return nil, nil, nil, errors.New("@" + tag + ": works only in functions definitions")
				}

				if constructorCommentType != "" {
					err := fmt.Errorf("@%v %v: constructor comment declared twice", tag, comment)
					return nil, nil, nil, err
				}
				constructorCommentType = tag
				constructorComment = comment

			case tagParam:
				const commentParts = 2

				parts := strings.SplitN(comment, " ", commentParts)
				realComment := ""
				paramName := parts[0]
				if len(parts) >= commentParts {
					realComment = strings.TrimSpace(parts[1])
				}

				if _, ok := argumentComments[paramName]; ok {
					return nil, nil, nil, errors.New("@param " + paramName + ": comment declared twice")
				}
				if realComment != "" {
					argumentComments[paramName] = realComment
				}

			default:
				return nil, nil, nil, errors.New("@" + commentTag + ": invalid comment tag")
			}

		case item.Declaration != nil:
			decl, err := normalizeDeclaration(item.Declaration, constructorComment, argumentComments)
			if err != nil {
				return nil, nil, nil, err
			}
			// todo check crc
			//if decl.getCRC() != decl.CRC {
			//	return nil, nil, nil, errors.New(decl.Name.String() + ": CRC mismatch! Described: " + fmt.Sprintf("0x%08x", decl.CRC) + " Calculated: " + fmt.Sprintf("0x%08x", decl.getCRC()))
			//}
			if funcMode {
				decl.Category = tl.CategoryFunction
				funcDecls = append(funcDecls, decl)
			} else {
				decl.Category = tl.CategoryPredict
				typeDecls = append(typeDecls, decl)
			}

			if currentTypeComment != "" {
				v, _ := decl.Type, true
				//if !ok {
				//	typStr := decl.Type.String()
				//	err := errors.New("@type: comment set to " + typStr + ", which is impossible")
				//	return nil, nil, nil, err
				//}

				if typeComments == nil {
					typeComments = map[tl.Name]string{}
				}

				if _, ok := typeComments[v.Name]; ok {
					err := errors.New("@type: for " + tl.Name(v.Name).String() + ", type comment defined twice")
					return nil, nil, nil, err
				}

				typeComments[v.Name] = currentTypeComment
				currentTypeComment = ""
			}

			constructorComment = ""
			constructorCommentType = ""
			argumentComments = map[string]string{}
		}
	}

	return typeDecls, funcDecls, typeComments, nil
}

func normalizeProgram(program *declaration.Program) (*tl.Schema, error) {
	typeDecls, funcDecls, _, err := normalizeEntries(program.Entries)
	if err != nil {
		return nil, err
	}

	decls := make([]tl.Declaration, 0)

	for _, decl := range typeDecls {
		decls = append(decls, *decl)
	}

	for _, decl := range funcDecls {
		decls = append(decls, *decl)
	}

	//typeDeclMap := map[tl.Name]tl.TypeDeclaration{}
	//for typ, decl := range sortedTypeDecls {
	//	var comment string
	//	if v, ok := comments[typ]; ok {
	//		comment = v
	//	}
	//
	//	typeDeclMap[typ] = tl.TypeDeclaration{
	//		Comment:      comment,
	//		Declarations: decl,
	//	}
	//}

	return &tl.Schema{
		Declarations: decls,
	}, nil
}
