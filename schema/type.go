// Copyright (c) 2022-2024 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package schema

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type TLType interface {
	_type()

	fmt.Stringer
}

var (
	_ TLType = TLTypeCommon(TLName{})
	_ TLType = TLTypeVector(TLName{})
)

type TLTypeCommon TLName

func (_ TLTypeCommon) _type()         {}
func (t TLTypeCommon) String() string { return TLName(t).String() }

type TLTypeVector TLName

func (_ TLTypeVector) _type()         {}
func (t TLTypeVector) String() string { return "Vector<" + TLName(t).String() + ">" }

func isFirstRuneUpper(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsUpper(r)
}
