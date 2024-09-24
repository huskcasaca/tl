package tl

import (
	"github.com/quenbyako/ext/slices"
)

// LayeredSchema represents a slice of tl schemas with their layers.
type LayeredSchema struct {
	Layers  []uint8          // Layers holds the sequence of layers.
	Schemas map[uint8]Schema // Schemas maps each layer to its corresponding Schema.
}

// Schema represents a single schema layer with its declarations.
type Schema struct {
	Layer      uint8        // Layer indicates the schema layer number.
	Definition []Definition // Definition holds the list of declarations in this schema.
}

// Types returns the list of types in this schema.
func (s *Schema) Types() []Type {
	types := map[string]Type{}
	for _, decl := range s.Definition {
		types[decl.Type.String()] = decl.Type
	}
	var typeSlice []Type
	for _, t := range types {
		typeSlice = append(typeSlice, t)
	}
	return slices.SortFunc(typeSlice, SortType)
}

// Predicts returns the list of predicts in this schema.
func (s *Schema) Predicts() []Definition {
	var predicts []Definition
	for _, d := range s.Definition {
		if d.Category == CategoryPredict {
			predicts = append(predicts, d)
		}
	}
	return slices.SortFunc(predicts, SortDeclarations)
}

// Functions returns the list of functions in this schema.
func (s *Schema) Functions() []Definition {
	var predicts []Definition
	for _, d := range s.Definition {
		if d.Category == CategoryFunction {
			predicts = append(predicts, d)
		}
	}
	return slices.SortFunc(predicts, SortDeclarations)
}

// TypePredicts returns the list of predicts grouped by type.
func (s *Schema) TypePredicts() []TypeDefinitions {
	typePredicts := map[Type][]Definition{}
	for _, d := range s.Definition {
		if d.Category == CategoryPredict {
			typePredicts[d.Type] = append(typePredicts[d.Type], d)
		}
	}
	predicts := []TypeDefinitions{}
	for t, d := range typePredicts {
		predicts = append(predicts, TypeDefinitions{Type: t, Definitions: slices.SortFunc(d, SortDeclarations)})
	}
	return slices.SortFunc(predicts, func(a, b TypeDefinitions) int { return SortType(a.Type, b.Type) })
}

type CRCIndex struct {
	Type  Name
	Index int
}

func (s *Schema) MakeCRCIndex() map[uint32]CRCIndex {
	res := make(map[uint32]CRCIndex, len(s.Definition))
	//for typ, decl := range s.Declarations {
	//	for i, o := range decl.Declarations {
	//		res[o.CRC] = CRCIndex{
	//			Type:  typ,
	//			Index: i,
	//		}
	//	}
	//}

	return res
}

type TypeDefinitions struct {
	Comment     string
	Type        Type
	Definitions []Definition // must be sorted by name
}
