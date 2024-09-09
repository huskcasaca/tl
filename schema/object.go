// Copyright (c) 2022-2024 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package schema

import (
	"cmp"
	"fmt"
	"hash/crc32"
	"strings"
)

type TLObject struct {
	Comment   string
	Name      TLName
	CRC       uint32
	Params    TLParams
	PolyTypes TLTypes
	Type      TLType
}

type TLName struct {
	Namespace string
	Key       string
}

func GetTLNameFromString(s string) TLName {
	groups := strings.Split(s, ".")
	var namespace string
	key := groups[0]
	if len(groups) > 1 {
		namespace = groups[0]
		key = groups[1]
	}

	return TLName{Namespace: namespace, Key: key}
}

func (o TLName) String() string {
	if o.Namespace != "" {
		return o.Namespace + "." + o.Key
	}

	return o.Key
}

func (o TLName) IsInterface() bool { return isFirstRuneUpper(o.Key) }

func (o TLName) Cmp(b TLName) int { return cmpObjName(o, b) }

func cmpObjName(a, b TLName) int {
	if c := cmp.Compare(a.Namespace, b.Namespace); c != 0 {
		return c
	} else if c := cmp.Compare(a.Key, b.Key); c != 0 {
		return c
	} else {
		return 0
	}
}

func sortObject(a, b TLObject) int { return cmpObjName(a.Name, b.Name) }

type TLObjectType uint8

const (
	TLObjectTypeUnknown TLObjectType = iota
	TLObjectTypeConstructor
	TLObjectTypeEnum
	TLObjectTypeMethod
)

func (o TLObjectType) String() string {
	switch o {
	case TLObjectTypeConstructor:
		return "constructor"
	case TLObjectTypeEnum:
		return "enum"
	case TLObjectTypeMethod:
		return "method"
	default:
		return "<UNKNOWN>"
	}
}

// how CRC is calculated:
// first of all, constructor formats to canonical state, e.g.
// `user#12325 id:int         first_name:string    last_name:string = User`
// ↓↓↓
// `user id:int first_name:string last_name:string = User`
//
// all bit triggers are removing, e.g.
// `messages.clearRecentStickers#8999602d flags:# attached:flags.0?true = Bool`
// ↓↓↓
// `messages.clearRecentStickers flags:# = Bool`
//
// For vectors like `getSmthn items:Vector<int> = Bool` i still don't understand
// how to generate, cause it fails in real mtproto schema.
func (o *TLObject) getCRC() uint32 {
	if o.CRC != 0 {
		return o.CRC
	}

	filtered := make(TLParams, 0, len(o.Params))
	for _, item := range o.Params {
		if _, ok := item.(TLTriggerParam); !ok {
			filtered = append(filtered, item)
		}
	}

	fieldsStr := ""
	if len(filtered) > 0 {
		fieldsStr = " " + filtered.String()
	}

	return crc32.ChecksumIEEE([]byte(fmt.Sprintf("%v%v = %v;", o.Name.String(), fieldsStr, o.Type)))
}

func (o *TLObject) String() string {
	fields := ""
	if len(o.Params) > 0 {
		fields = " " + o.Params.String()
	}

	return fmt.Sprintf("%v#%08x%v = %v;", o.Name.String(), o.getCRC(), fields, o.Type)
}

func (o *TLObject) Comments(typ TLObjectType) []string {
	var res []string
	if o.Comment != "" {
		res = append(res, "// @"+typ.String()+" "+o.Comment)
	}

	return append(res, o.Params.Comments()...)
}
