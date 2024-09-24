package program

import (
	"fmt"
	"github.com/xelaj/tl"
)

// Argument represents TL formal described in https://core.telegram.org/mtproto/TL-formal#combinator-declarations
type Argument struct {
	Ident ArgIdent `parser:"@@ colon"` // var:
	Flag  *Flag    `parser:"@@?"`      // flags.1?
	Type  Type     `parser:"@@"`       // type
}

// ArgIdent is identifier of argument name
// var-ident ::= lc-ident | uc-ident
// var-ident-opt ::= var-ident | _
type ArgIdent struct {
	Ident string `parser:"@lc_ident | @uc_ident | @underscore"`
}

func (f *ArgIdent) String() string {
	return f.Ident
}

func (f *ArgIdent) Empty() bool {
	return f.Ident == "_"
}

// Flag is flag of type
type Flag struct {
	Ident string `parser:"@lc_ident"`
	Index int    `parser:"dot @nat_const question_mark"`
}

func ParseArgument(arg *Argument, comment string) (tl.Param, error) {
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
