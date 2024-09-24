package tl

import (
	"fmt"
	"strings"
)

type Types []Type

type Type struct {
	Name
	SubType *Type
	fmt.Stringer
}

func (t Type) String() string {
	if t.SubType == nil {
		return t.Name.String()
	}
	return fmt.Sprintf("%s<%s>", t.Name.String(), (t.SubType).String())
}

func SortType(a, b Type) int { return strings.Compare(a.String(), b.String()) }

var (
	AnyType = Type{Name: Name{Key: "Type"}}
)
