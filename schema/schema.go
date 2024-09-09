// Copyright (c) 2022-2024 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package schema

import (
	"fmt"
	"strings"

	"github.com/quenbyako/ext/slices"
)

type TLSchema struct {
	ObjSeq     []TLName
	TypeObjMap map[TLName]TypeTLObjects // key is type name
	EnumObjMap map[TLName]EnumTLObjects // key is enum name

	FunctionSeq []string
	FunctionMap map[string][]TLObject // methods must be sorted by name
}

func (s *TLSchema) String() string {
	var parts []string
	for _, typ := range s.ObjSeq {
		if obj, ok := s.TypeObjMap[typ]; ok {
			parts = append(parts, obj.String())
		} else if enum, ok := s.EnumObjMap[typ]; ok {
			parts = append(parts, enum.String())
		} else {
			panic(fmt.Sprintf("missed type %#v", typ))
		}
	}

	if len(s.FunctionSeq) == 0 {
		return strings.Join(parts, "\n\n") + "\n"
	}

	parts = append(parts, "---functions---")

	for _, group := range s.FunctionSeq {
		obj, ok := s.FunctionMap[group]
		if !ok {
			panic(fmt.Sprintf("missed group %#v", group))
		}

		parts = append(parts, methodsString(obj, group))
	}

	return strings.Join(parts, "\n\n") + "\n"
}

type CRCIndex struct {
	Type        TLName
	ObjectIndex int
}

func (s *TLSchema) MakeCRCIndex() map[uint32]CRCIndex {
	res := make(map[uint32]CRCIndex, len(s.TypeObjMap))
	for typ, obj := range s.TypeObjMap {
		for i, o := range obj.Objects {
			res[o.CRC] = CRCIndex{
				Type:        typ,
				ObjectIndex: i,
			}
		}
	}

	return res
}

type TypeTLObjects struct {
	Comment string
	Objects []TLObject // must be sorted by name
}

func (s TypeTLObjects) String() string {
	var parts []string
	if s.Comment != "" {
		parts = append(parts, "// @type "+s.Comment)
	}

	for _, obj := range slices.SortFunc(s.Objects, sortObject) {
		parts = append(parts, obj.Comments(TLObjectTypeConstructor)...)
		parts = append(parts, obj.String())
	}

	return strings.Join(parts, "\n")
}

type EnumTLObjects struct {
	Comment string
	Objects []TLObject // must be sorted by name
}

func (s EnumTLObjects) String() (res string) {
	var parts []string
	if s.Comment != "" {
		parts = append(parts, "// @type "+s.Comment)
	}

	for _, obj := range slices.SortFunc(s.Objects, sortObject) {
		parts = append(parts, obj.Comments(TLObjectTypeEnum)...)
		parts = append(parts, obj.String())
	}

	return strings.Join(parts, "\n")
}

func methodsString(methods []TLObject, group string) (res string) {
	var parts []string

	if group != "" {
		group += "."
	}

	for _, obj := range slices.SortFunc(methods, sortObject) {
		parts = append(parts, obj.Comments(TLObjectTypeMethod)...)
		parts = append(parts, obj.String())
	}

	return strings.Join(parts, "\n")
}
