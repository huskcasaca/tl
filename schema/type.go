package schema

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type TLTypes []TLType

type TLType interface {
	_type()

	Name() TLName
	IsInList() bool
	fmt.Stringer
}

var (
	_ TLType = TLTypeCommon{TLName{}, false}
	_ TLType = TLTypeVector{TLName{}, false}
)

var (
	TLAnyType TLType = TLTypeCommon{TLName{Key: "Type"}, false}
)

type TLTypeCommon struct {
	TLName
	IsInterface bool
}

func (_ TLTypeCommon) _type()         {}
func (t TLTypeCommon) String() string { return t.TLName.String() }
func (t TLTypeCommon) Name() TLName   { return t.TLName }
func (t TLTypeCommon) IsInList() bool { return t.IsInterface }

type TLTypeVector struct {
	TLName
	IsInterface bool
}

func (_ TLTypeVector) _type()         {}
func (t TLTypeVector) String() string { return "Vector<" + t.TLName.String() + ">" }
func (t TLTypeVector) Name() TLName   { return t.TLName }
func (t TLTypeVector) IsInList() bool { return t.IsInterface }

func isFirstRuneUpper(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsUpper(r)
}
