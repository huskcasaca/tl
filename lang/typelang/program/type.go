package program

import (
	"errors"
	"github.com/xelaj/tl"
	"strings"
)

// Type is type of argument and result
type Type struct {
	Optional bool      `parser:"@excl_mark?"`
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

func (i *Type) Normalize() (tl.Type, error) {
	name := tl.NameFromString(i.Ident.String())
	if len(i.SubTypes) == 0 {
		return tl.Type{Name: name}, nil
	}
	if len(i.SubTypes) == 1 {
		typ, err := i.SubTypes[0].Normalize()
		if err != nil {
			return typ, err
		}
		return tl.Type{Name: name, SubType: &typ}, nil
	}
	return tl.Type{}, errors.New("too many subtypes")
}
