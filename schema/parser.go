package schema

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/quenbyako/ext/slices"

	"github.com/xelaj/tl/schema/internal/declaration"
	"github.com/xelaj/tl/schema/internal/lexer"
)

//nolint:gochecknoglobals // obviously parser must be global
var parser = participle.MustBuild[declaration.Program](
	participle.Lexer(lexer.NewDefinition()),
)

func ParseFile(filename string) (*TLSchema, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Parse(filename, f)
}

func ParseFS(fsys fs.FS, filename string) (*TLSchema, error) {
	f, err := fsys.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Parse(filename, f)
}

func ParseString(filename string, content string) (*TLSchema, error) {
	return Parse(filename, strings.NewReader(content))
}

func Parse(filename string, content io.Reader) (*TLSchema, error) {
	res, err := parser.Parse(filename, content)
	if err != nil {
		return nil, err //nolint:wrapcheck // it's important to keep error
	}

	normalized, err := normalizeProgram(res)
	if err != nil {
		return nil, err
	}

	return normalized, nil
}

func normalizeIdent(i *declaration.Type) (TLType, error) {
	name := GetTLNameFromString(i.Ident.String())
	if len(i.SubTypes) == 0 {
		return TLType{TLName: name}, nil
	}
	types := make([]TLType, len(i.SubTypes))

	for i, item := range i.SubTypes {
		typ, err := normalizeIdent(&item)
		if err != nil {
			return typ, err
		}
		types[i] = typ
	}

	return TLType{
		TLName:  name,
		TLTypes: types,
	}, nil
}

func normalizeArgument(arg *declaration.Argument, comment string) (TLParam, error) {
	typ, err := normalizeIdent(&arg.Term)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", arg.Ident.String(), err)
	}

	if arg.Flag == nil {
		if typ.Key == "#" {
			return TLBitflagParam{
				Comment: comment,
				Name:    arg.Ident.String(),
			}, nil
		}

		return TLRequiredParam{
			Comment: comment,
			Name:    arg.Ident.String(),
			Type:    typ,
		}, nil
	}

	if typ.Key == "true" {
		return TLTriggerParam{
			Comment:     comment,
			Name:        arg.Ident.String(),
			FlagTrigger: arg.Flag.Ident,
			BitTrigger:  arg.Flag.Index,
		}, nil
	}

	return TLOptionalParam{
		Comment:     comment,
		Name:        arg.Ident.String(),
		Type:        typ,
		FlagTrigger: arg.Flag.Ident,
		BitTrigger:  arg.Flag.Index,
	}, nil
}

func normalizeCombinator(decl *declaration.Declaration, constructorComment string, argsComments map[string]string) (*TLDeclaration, error) {
	parts := strings.Split(decl.Combinator, "#") // guaranteed to split by two parts, lexer handles it
	if len(parts) == 0 || len(parts) > 2 {
		return nil, errors.New(decl.Combinator + ": invalid combinator")
	}
	name := GetTLNameFromString(parts[0])
	crc := uint64(0)
	if len(parts) == 2 {
		crc1, err := strconv.ParseUint(parts[1], 16, 32)
		if err != nil {
			return nil, fmt.Errorf("%v: %w", decl.Combinator, err)
		}
		crc = crc1
	}

	params := make([]TLParam, len(decl.Args))
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

	polyParams := make(TLParams, len(decl.OptArgs))

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
		polyParams[i], argErr = normalizeArgument(&arg, comment)
		if argErr != nil {
			return nil, fmt.Errorf("%v: %w", decl.Combinator, argErr)
		}
	}

	typ, err := normalizeIdent(&declaration.Type{
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

	return &TLDeclaration{
		Comment:    constructorComment,
		Name:       name,
		CRC:        uint32(crc),
		Params:     params,
		PolyParams: polyParams,
		Type:       typ,
	}, nil
}

const (
	tagType        = "type"
	tagEnum        = "enum"
	tagConstructor = "constructor"
	tagMethod      = "method"
	tagParam       = "param"
)

func normalizeEntries(items []declaration.ProgramEntry) (typeDecls []*TLDeclaration, funcDecls []*TLDeclaration, typeComments map[TLName]string, err error) {
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

			case tagEnum, tagConstructor, tagMethod:
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
			decl, err := normalizeCombinator(item.Declaration, constructorComment, argumentComments)
			if err != nil {
				return nil, nil, nil, err
			}
			// todo check crc
			//if decl.getCRC() != decl.CRC {
			//	return nil, nil, nil, errors.New(decl.Name.String() + ": CRC mismatch! Described: " + fmt.Sprintf("0x%08x", decl.CRC) + " Calculated: " + fmt.Sprintf("0x%08x", decl.getCRC()))
			//}
			if funcMode {
				funcDecls = append(funcDecls, decl)
			} else {
				typeDecls = append(typeDecls, decl)
			}

			if currentTypeComment != "" {
				v, _ := decl.Type, true
				//if !ok {
				//	typStr := decl.Type.String()
				//	err := errors.New("@type: comment set to " + typStr + ", which is impossible")
				//	return nil, nil, nil, err
				//}

				if _, ok := typeComments[TLName(v.TLName)]; ok {
					err := errors.New("@type: for " + TLName(v.TLName).String() + ", type comment defined twice")
					return nil, nil, nil, err
				}

				typeComments[TLName(v.TLName)] = currentTypeComment
				currentTypeComment = ""
			}

			constructorComment = ""
			constructorCommentType = ""
			argumentComments = map[string]string{}
		}
	}

	return typeDecls, funcDecls, typeComments, nil
}

func normalizeProgram(program *declaration.Program) (*TLSchema, error) {
	typeDecls, funcDecls, comments, err := normalizeEntries(program.Entries)
	if err != nil {
		return nil, err
	}

	typeSeq := []TLName{}
	sortedTypeDecls := map[TLName][]TLDeclaration{}
	for _, decl := range typeDecls {
		declType := decl.Type.TLName

		if !slices.Contains(typeSeq, declType) {
			typeSeq = append(typeSeq, declType)
		}

		sortedTypeDecls[declType] = append(sortedTypeDecls[declType], *decl)
	}

	typeDeclMap := map[TLName]TLTypeDeclaration{}
	for typ, decl := range sortedTypeDecls {
		var comment string
		if v, ok := comments[typ]; ok {
			comment = v
		}

		typeDeclMap[typ] = TLTypeDeclaration{
			Comment:      comment,
			Declarations: decl,
		}
	}

	funcSeq := []TLName{}
	funcDeclMap := map[TLName]TLDeclaration{}
	for _, decl := range funcDecls {
		funcDeclMap[decl.Name] = *decl
	}

	return &TLSchema{
		TypeSeq:     typeSeq,
		TypeDeclMap: typeDeclMap,
		FuncSeq:     funcSeq,
		FuncDeclMap: funcDeclMap,
	}, nil
}
