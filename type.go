package tl

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type TLTypes []TLType

type TLType struct {
	TLName
	TLTypes
	fmt.Stringer
}

func (t TLType) Name() TLName   { return t.TLName }
func (t TLType) Types() TLTypes { return t.TLTypes }
func (t TLType) String() string { return t.TLName.String() }

var (
	TLAnyType = TLType{TLName: TLName{Key: "Type"}}
)

func isFirstRuneUpper(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsUpper(r)
}
