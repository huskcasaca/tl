// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package declaration

import (
	"strings"
)

type Program struct {
	Constraints []ProgramEntry `parser:"@@*"`
	Methods     []ProgramEntry `parser:"(functions_decl newline @@*)?"`
}

type ProgramEntry struct {
	Newline bool            `parser:"@newline |"`
	Decl    *CombinatorDecl `parser:"@@ ws? semicolon ( newline | EOF ) |"`
	Comment *string         `parser:"@comment ( newline | EOF ) "`
}

type CombinatorDecl struct {
	ID       string     `parser:"@lc_ident_full ws"`
	PolyType *PolyType  `parser:"(@@ ws)?"` // generic parameter, e.g., {X:Type}
	Args     []Argument `parser:"(@@ (ws @@)* ws)? equals ws?"`
	Result   ResultType `parser:"@@"`
}

type PolyType struct {
	Name FieldName `parser:"'{' @@ colon 'Type' '}'"` // captures X:Type
}

type ResultType struct {
	Simple *string    `parser:"( @uc_ident_ns | @uc_ident )"`
	Expr   *Extension `parser:"@@?"`
}

type Argument struct {
	Ident       ArgumentName  `parser:"@@ colon"` // var_name:
	Conditional *ArgumentFlag `parser:"@@?"`      // flags.1?
	Term        ArgumentType  `parser:"@@"`       // type
}

type ArgumentType struct {
	Simple    Field      `parser:"@@"`
	Extension *Extension `parser:"@@?"`
}

func (i *ArgumentType) String() string {
	res := i.Simple.String()
	if i.Extension != nil {
		res += i.Extension.String()
	}
	return res
}

type Field struct {
	Var  *FieldName `parser:"@@ |"`
	Type *FieldType `parser:"@@"`
}

func (s *Field) String() string {
	switch {
	case s.Var != nil:
		return s.Var.Value
	case s.Type != nil:
		return s.Type.String()
	default:
		panic("impossible situation")
	}
}

type FieldName struct {
	Value string `parser:"@lc_ident | @uc_ident"`
}

type FieldType struct {
	Ident *string `parser:"@uc_ident_ns | @lc_ident_ns |"`
	Empty bool    `parser:"@hash"`
}

func (t *FieldType) String() string {
	if t.Empty {
		return "#"
	}

	return *t.Ident
}

type Extension struct {
	Inner []Field `parser:"langle @@ (ws @@)* rangle"`
}

func (i *Extension) String() string {
	items := make([]string, len(i.Inner))
	for i, item := range i.Inner {
		items[i] = item.String()
	}

	return "<" + strings.Join(items, " ") + ">"
}

type ArgumentName struct {
	Ident *FieldName `parser:"@@ |"`
	Empty bool       `parser:"@underscore"`
}

func (i *ArgumentName) String() string {
	if i.Empty {
		return "_"
	}

	return i.Ident.Value
}

type ArgumentFlag struct {
	Ident FieldName `parser:"@@"`
	Index int       `parser:"dot @nat_const question_mark"`
}
