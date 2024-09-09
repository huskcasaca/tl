// Copyright (c) 2022-2024 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

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
	"github.com/quenbyako/ext/set"
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

func normalizeIdent(i *declaration.ArgType) (TLType, error) {
	if i.Extension != nil {
		if len(i.Extension.Inner) > 1 {
			return nil, errors.New(i.String() + ": too many modificators")
		}
		if i.Simple.String() != "Vector" {
			return nil, errors.New(i.String() + ": only Vector is allowed to modify")
		}

		return TLTypeVector{TLName: GetTLNameFromString(i.Extension.Inner[0].String())}, nil
	}

	return TLTypeCommon{TLName: GetTLNameFromString(i.Simple.String())}, nil
}

func normalizeArgument(arg *declaration.Argument, comment string) (TLParam, error) {
	typ, err := normalizeIdent(&arg.Term)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", arg.Ident.String(), err)
	}

	if arg.Flag == nil {
		if typ, ok := typ.(TLTypeCommon); ok && typ.Key == "#" {
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

	if v, ok := typ.(TLTypeCommon); ok && v.Key == "true" {
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

func normalizeCombinator(
	decl *declaration.Declaration,
	constructorComment string,
	argsComments map[string]string,
	functionsMode bool,
) (*TLObject, error) {
	parts := strings.Split(decl.Combinator, "#") // guaranteed to split by two parts, lexer handles it
	name := GetTLNameFromString(parts[0])
	crcStr := parts[1]

	// same: lexer handles everything already
	crc, err := strconv.ParseUint(crcStr, 16, 32)
	if err != nil {
		panic(err)
	}

	// todo: Matches(Name, CRC)

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

	polytypes := make(TLParams, len(decl.OptArgs))

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
		polytypes[i], argErr = normalizeArgument(&arg, comment)
		if argErr != nil {
			return nil, fmt.Errorf("%v: %w", decl.Combinator, argErr)
		}
	}

	typ, err := normalizeIdent(&declaration.ArgType{
		Simple:    decl.Result.Simple,
		Extension: decl.Result.Extension,
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

	return &TLObject{
		Comment:    constructorComment,
		Name:       name,
		CRC:        uint32(crc),
		Params:     params,
		PolyParams: polytypes,
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

func normalizeEntries(items []declaration.ProgramEntry, functionsMode bool) ([]*TLObject, map[TLName]string, error) {
	var (
		objects      = []*TLObject{}
		typeComments = map[TLName]string{}

		currentTypeComment     string
		constructorComment     string
		constructorCommentType string
		argumentComments       = map[string]string{}
	)
	for _, item := range items {
		switch {
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
				if functionsMode {
					err := fmt.Errorf("@%v %v: impossible on functions", tag, comment)
					return nil, nil, err
				}

				// special case, declaration of type apply to first available type (= SomeType;).
				// when type comment applied to type, currentTypeComment is clearing, so we checking
				// is this type comment declared twice.
				if currentTypeComment != "" {
					err := fmt.Errorf("@%v %v: type comment declared twice", tag, comment)
					return nil, nil, err
				}
				currentTypeComment = comment

			case tagEnum, tagConstructor, tagMethod:
				if functionsMode && tag != tagMethod {
					return nil, nil, errors.New("@" + tag + ": works only in type definitions")
				} else if !functionsMode && comment == tagMethod {
					return nil, nil, errors.New("@" + tag + ": works only in functions definitions")
				}

				if constructorCommentType != "" {
					err := fmt.Errorf("@%v %v: constructor comment declared twice", tag, comment)
					return nil, nil, err
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
					return nil, nil, errors.New("@param " + paramName + ": comment declared twice")
				}
				if realComment != "" {
					argumentComments[paramName] = realComment
				}

			default:
				return nil, nil, errors.New("@" + commentTag + ": invalid comment tag")
			}

		case item.Declaration != nil:
			obj, err := normalizeCombinator(item.Declaration, constructorComment, argumentComments, functionsMode)
			if err != nil {
				return nil, nil, err
			}
			objects = append(objects, obj)

			if currentTypeComment != "" {
				v, ok := obj.Type.(TLTypeCommon)
				if !ok {
					typStr := obj.Type.String()
					err := errors.New("@type: comment set to " + typStr + ", which is impossible")
					return nil, nil, err
				}

				if _, ok := typeComments[TLName(v.TLName)]; ok {
					err := errors.New("@type: for " + TLName(v.TLName).String() + ", type comment defined twice")
					return nil, nil, err
				}

				typeComments[TLName(v.TLName)] = currentTypeComment
				currentTypeComment = ""
			}

			constructorComment = ""
			constructorCommentType = ""
			argumentComments = map[string]string{}
		}
	}

	return objects, typeComments, nil
}

func normalizeProgram(program *declaration.Program) (*TLSchema, error) {
	typesRaw, comments, err := normalizeEntries(program.Constraints, false)
	if err != nil {
		return nil, err
	}

	typeOrder := []TLName{}
	typeSortedOjbects := map[TLName][]TLObject{}
	anyTypeHasField := set.New[TLName]()
	for _, obj := range typesRaw {
		objTypeRaw, ok := obj.Type.(TLTypeCommon)
		if !ok {
			return nil, fmt.Errorf("object %#v: type is not interface", obj.Name)
		}
		objType := TLName(objTypeRaw.TLName)

		if !slices.Contains(typeOrder, objType) {
			typeOrder = append(typeOrder, objType)
		}

		typeSortedOjbects[objType] = append(typeSortedOjbects[objType], *obj)
		if len(obj.Params) > 0 {
			anyTypeHasField = anyTypeHasField.Add(objType)
		}
	}

	types := map[TLName]TypeTLObjects{}
	enums := map[TLName]EnumTLObjects{}
	for typ, objs := range typeSortedOjbects {
		var comment string
		if v, ok := comments[typ]; ok {
			comment = v
		}

		if anyTypeHasField.Has(typ) {
			types[typ] = TypeTLObjects{
				Comment: comment,
				Objects: objs,
			}
		} else {
			enums[typ] = EnumTLObjects{
				Comment: comment,
				Objects: objs,
			}
		}
	}

	methods, _, err := normalizeEntries(program.Methods, true)
	if err != nil {
		return nil, err
	}

	methodGroupOrder := []string{}
	methodGroupSortedOjbects := map[string][]TLObject{}
	for _, obj := range methods {
		if !slices.Contains(methodGroupOrder, obj.Name.Namespace) {
			methodGroupOrder = append(methodGroupOrder, obj.Name.Namespace)
		}

		methodGroupSortedOjbects[obj.Name.Namespace] = append(methodGroupSortedOjbects[obj.Name.Namespace], *obj)
	}

	return &TLSchema{
		ObjSeq:      typeOrder,
		TypeObjMap:  types,
		EnumObjMap:  enums,
		FunctionSeq: methodGroupOrder,
		FunctionMap: methodGroupSortedOjbects,
	}, nil
}
