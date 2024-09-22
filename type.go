package tl

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type Types []Type

type Type struct {
	Name
	Types
	fmt.Stringer
}

func (t Type) String() string { return t.Name.String() }

var (
	AnyType = Type{Name: Name{Key: "Type"}}
)

func isFirstRuneUpper(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsUpper(r)
}
