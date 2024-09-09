// Copyright (c) 2022-2024 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package codegen

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/quenbyako/ext/maps"
	"github.com/quenbyako/ext/slices"

	"github.com/xelaj/tl/schema"
)

func Generate(s *schema.Schema) (string, error) {
	f := jen.NewFile("main")

	isInterface := func(name schema.TLName) bool {
		_, ok := s.TypeObjMap[name]
		return ok
	}

	for _, name := range slices.SortFunc(maps.Keys(s.TypeObjMap), func(a, b schema.TLName) int { return a.Cmp(b) }) {
		f.Add(generateObjects(name, s.TypeObjMap[name], isInterface))
	}

	for _, name := range slices.SortFunc(maps.Keys(s.EnumObjMap), func(a, b schema.TLName) int { return a.Cmp(b) }) {
		f.Add(generateEnum(name, s.EnumObjMap[name]))
	}

	f.Add(generateRequestBareFunction())

	for _, methodGroup := range slices.Sort(maps.Keys(s.FunctionMap)) {
		methods := s.FunctionMap[methodGroup]
		for _, method := range methods {
			obj := generateRequestType(methodGroup, method, isInterface)
			f.Add(obj, jen.Line())
		}
	}

	return fmt.Sprintf("%#v", f), nil
}
