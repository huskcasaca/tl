package typelang

import (
	"errors"
	"fmt"
	"github.com/xelaj/tl"
	"github.com/xelaj/tl/parser/typelang/declaration"
	"strconv"
	"strings"
)

func ParseType(i *declaration.Type) (tl.Type, error) {
	name := tl.NameFromString(i.Ident.String())
	if len(i.SubTypes) == 0 {
		return tl.Type{Name: name}, nil
	}
	if len(i.SubTypes) == 1 {
		typ, err := ParseType(&i.SubTypes[0])
		if err != nil {
			return typ, err
		}
		return tl.Type{Name: name, SubType: &typ}, nil
	}
	return tl.Type{}, errors.New("too many subtypes")
}

func ParseArgument(arg *declaration.Argument, comment string) (tl.Param, error) {
	typ, err := ParseType(&arg.Type)
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

func ParseDeclaration(decl *declaration.Declaration, comments []Annotation) (*tl.Declaration, error) {
	parts := strings.Split(decl.Combinator, "#")
	if len(parts) == 0 || len(parts) > 2 {
		return nil, errors.New(decl.Combinator + ": invalid combinator")
	}
	name := tl.NameFromString(parts[0])
	crc := uint64(0)
	if len(parts) == 2 {
		crc1, err := strconv.ParseUint(parts[1], 16, 32)
		if err != nil {
			return nil, fmt.Errorf("%v: %w", decl.Combinator, err)
		}
		crc = crc1
	}
	var constComment string
	var returnComment string
	argsComments := make(map[string]string)
	for _, c := range comments {
		switch {
		case c.Tag == AnnotationParamTag:
			param, comment := c.SplitParam()
			argsComments[param] = comment
		case c.Tag == AnnotationConstructorTag:
			constComment = c.Value
		case c.Tag == AnnotationReturnTag:
			returnComment = c.Value
		}
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
		params[i], argErr = ParseArgument(&arg, comment)
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
		optParams[i], argErr = ParseArgument(&arg, comment)
		if argErr != nil {
			return nil, fmt.Errorf("%v: %w", decl.Combinator, argErr)
		}
	}

	typ, err := ParseType(&decl.Result)
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
		Name:       name,
		CRC:        uint32(crc),
		Params:     params,
		OptParams:  optParams,
		Type:       typ,
		Comment:    constComment,
		RetComment: returnComment,
	}, nil
}

func ParseProgram(program *declaration.Program) (*tl.Schema, error) {
	var (
		decls    = []tl.Declaration{}
		comments = []Annotation{}
		funcMode = false
	)
	for _, e := range program.Entries {
		switch {
		case e.TypeDecl:
			funcMode = false
		case e.FuncDecl:
			funcMode = true
		case e.Newline:
			comments = []Annotation{}

		case e.Comment != nil:
			a, err := ParseAnnotation(*e.Comment)
			if err != nil {
				return nil, err
			}
			if a != nil {
				comments = append(comments, *a)
			}

		case e.Declaration != nil:
			d, err := ParseDeclaration(e.Declaration, comments)
			if err != nil {
				return nil, err
			}
			if funcMode {
				d.Category = tl.CategoryFunction
			} else {
				d.Category = tl.CategoryPredict
			}
			// todo check crc
			//if decl.getCRC() != d.CRC {
			//	return nil, nil, nil, errors.New(decl.Name.String() + ": CRC mismatch! Described: " + fmt.Sprintf("0x%08x", decl.CRC) + " Calculated: " + fmt.Sprintf("0x%08x", decl.getCRC()))
			//}
			decls = append(decls, *d)
			comments = []Annotation{}
		}
	}
	return &tl.Schema{
		Declarations: decls,
	}, nil
}
