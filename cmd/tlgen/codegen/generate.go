package codegen

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/quenbyako/ext/maps"
	"github.com/quenbyako/ext/slices"

	"github.com/xelaj/tl/schema"
)

func Generate(s *schema.TLSchema) (string, error) {
	f := jen.NewFile("main")

	for _, name := range slices.SortFunc(maps.Keys(s.TypeDeclMap), func(a, b schema.TLName) int { return a.Cmp(b) }) {
		f.Add(generateObjects(name, s.TypeDeclMap[name]))
	}

	for _, name := range slices.SortFunc(maps.Keys(s.EnumDeclMap), func(a, b schema.TLName) int { return a.Cmp(b) }) {
		f.Add(generateEnums(name, s.EnumDeclMap[name]))
	}

	f.Add(generateRequestFunc())

	for _, namespace := range slices.Sort(maps.Keys(s.FuncDeclMap)) {
		methods := s.FuncDeclMap[namespace]
		for _, method := range methods {
			obj := generateRequestType(namespace, method)
			f.Add(obj, jen.Line())
		}
	}

	return fmt.Sprintf("%#v", f), nil
}
