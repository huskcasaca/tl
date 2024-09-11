package schema

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type TLTypes []TLType

type TLType interface {
	Name() TLName
	fmt.Stringer
}

var (
	TLAnyType TLType = TLTypeCommon{TLName{Key: "Type"}}
)

type TLTypeCommon struct {
	TLName
}

func (t TLTypeCommon) String() string { return t.TLName.String() }
func (t TLTypeCommon) Name() TLName   { return t.TLName }

type TLTypeVector struct {
	TLName
}

func (t TLTypeVector) String() string { return "Vector<" + t.TLName.String() + ">" }
func (t TLTypeVector) Name() TLName   { return t.TLName }

func isFirstRuneUpper(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsUpper(r)
}
