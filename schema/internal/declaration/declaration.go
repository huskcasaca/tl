package declaration

import (
	"strings"
)

// Program represents TL formal described in https://core.telegram.org/mtproto/TL-formal#syntax
// Syntactically, a TL program consists of a stream of tokens (separated by spaces, which are ignored at this stage). General program structure:
//
// TL-program ::= constr-declarations { --- functions --- fun-declarations | --- types --- constr-declarations }
type Program struct {
	Entries []ProgramEntry `parser:"@@*"`
}

// ProgramEntry represents TL formal described in https://core.telegram.org/mtproto/TL-formal#syntax
//
// StartTypeDecl is used to indicate if the following declarations are type declarations
// constr-declarations ::= { declaration }
//
// StartFuncDecl is used to indicate if the following declarations are function declarations
// Newline is used to represent if the line is empty
// Declaration is used to represent functions and types
// declaration ::= combinator-decl | partial-app-decl | final-decl
// constr-declarations ::= { declaration }
// fun-declarations ::= { declaration }
//
// Comment is used to represent comments with `//`
type ProgramEntry struct {
	TypeDecl    bool         `parser:"@(constraints_decl newline) |"`
	FuncDecl    bool         `parser:"@(functions_decl newline) |"`
	Newline     bool         `parser:"@newline |"`
	Declaration *Declaration `parser:"@@ ws? semicolon ( newline | EOF ) |"`
	Comment     *string      `parser:"@comment ( newline | EOF ) "`
}

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

// Argument represents TL formal described in https://core.telegram.org/mtproto/TL-formal#combinator-declarations
type Argument struct {
	Ident ArgIdent `parser:"@@ colon"` // var:
	Flag  *Flag    `parser:"@@?"`      // flags.1?
	Term  Type     `parser:"@@"`       // type
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

// Type is type of argument and result
type Type struct {
	Modifier bool      `parser:"@excl_mark?"`
	Ident    TypeIdent `parser:"@@"`
	SubTypes []Type    `parser:"(langle @@ (ws @@)* rangle)?"`
}

func (i *Type) String() string {
	res := i.Ident.String()
	if len(i.SubTypes) > 0 {
		items := make([]string, len(i.SubTypes))
		for i, item := range i.SubTypes {
			items[i] = item.String()
		}
		res += "<" + strings.Join(items, " ") + ">"
	}
	return res
}

// TypeIdent is identifier of argument type
type TypeIdent struct {
	Ident string `parser:"@uc_ident_ns | @lc_ident_ns | @lc_ident | @uc_ident | @hash"`
}

func (s *TypeIdent) String() string {
	return s.Ident
}
