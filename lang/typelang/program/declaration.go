package program

import (
	"errors"
	"fmt"
	"github.com/xelaj/tl"
	"strconv"
	"strings"
)

// Declaration represents TL formal described in https://core.telegram.org/mtproto/TL-formal#combinator-declarations
//
// Combinator is combination of identifiers separated by `#`
// combinator-decl ::= full-combinator-id { opt-args } { args } = result-type ;
// full-combinator-id ::= lc-ident-full | _
// combinator-id ::= lc-ident-ns | _
//
// OptArgs are optional arguments (e.g., {X:Type})
// opt-args ::= { var-ident { var-ident } : [excl-mark] type-expr }
//
// Args are required arguments (e.g., var:string)
// args ::= var-ident-opt : [ conditional-def ] [ ! ] type-term
// args ::= [ var-ident-opt : ] [ multiplicity *] [ { args } ]
// args ::= ( var-ident-opt { var-ident-opt } : [!] type-term )
// args ::= [ ! ] type-term
// multiplicity ::= nat-term
// var-ident-opt ::= var-ident | _
// conditional-def ::= var-ident [ . nat-const ] ?
//
// Result is return type
// result-type ::= boxed-type-ident { subexpr }
// result-type ::= boxed-type-ident < subexpr { , subexpr } >
type Declaration struct {
	Combinator string     `parser:"@(lc_ident_full | (lc_ident (dot lc_ident)?)) ws"`
	OptArgs    []Argument `parser:"(open_brace @@ close_brace ws)*"`
	Args       []Argument `parser:"(@@ (ws @@)* ws)? equals ws?"`
	Result     Type       `parser:"@@"`
}

func (decl *Declaration) Normalize(comments []Annotation) (*tl.Definition, error) {
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
		params[i], argErr = arg.Normalize()
		if argErr != nil {
			return nil, fmt.Errorf("%v: %w", decl.Combinator, argErr)
		}
		params[i].SetComment(comment)
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
		optParams[i], argErr = arg.Normalize()
		if argErr != nil {
			return nil, fmt.Errorf("%v: %w", decl.Combinator, argErr)
		}
		optParams[i].SetComment(comment)
	}

	typ, err := decl.Result.Normalize()
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

	return &tl.Definition{
		Name:       name,
		CRC:        uint32(crc),
		Params:     params,
		OptParams:  optParams,
		Type:       typ,
		Comment:    constComment,
		RetComment: returnComment,
	}, nil
}
