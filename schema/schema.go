package schema

import (
	"fmt"
	"strings"

	"github.com/quenbyako/ext/slices"
)

type TLSchema struct {
	TypeSeq     []TLName
	TypeDeclMap map[TLName]TLTypeDeclaration // key is type name
	EnumDeclMap map[TLName]TLEnumDeclaration // key is enum name

	FuncSeq     []string
	FuncDeclMap map[string][]TLDeclaration // methods must be sorted by name
}

func (s *TLSchema) String() string {
	var parts []string
	for _, typ := range s.TypeSeq {
		if decl, ok := s.TypeDeclMap[typ]; ok {
			parts = append(parts, decl.String())
		} else if enum, ok := s.EnumDeclMap[typ]; ok {
			parts = append(parts, enum.String())
		} else {
			panic(fmt.Sprintf("missed type %#v", typ))
		}
	}

	if len(s.FuncSeq) == 0 {
		return strings.Join(parts, "\n\n") + "\n"
	}

	parts = append(parts, "---functions---")

	for _, group := range s.FuncSeq {
		decl, ok := s.FuncDeclMap[group]
		if !ok {
			panic(fmt.Sprintf("missed group %#v", group))
		}

		parts = append(parts, methodsString(decl, group))
	}

	return strings.Join(parts, "\n\n") + "\n"
}

type CRCIndex struct {
	Type  TLName
	Index int
}

func (s *TLSchema) MakeCRCIndex() map[uint32]CRCIndex {
	res := make(map[uint32]CRCIndex, len(s.TypeDeclMap))
	for typ, decl := range s.TypeDeclMap {
		for i, o := range decl.Declarations {
			res[o.CRC] = CRCIndex{
				Type:  typ,
				Index: i,
			}
		}
	}

	return res
}

type TLTypeDeclaration struct {
	Comment      string
	Declarations []TLDeclaration // must be sorted by name
}

func (s TLTypeDeclaration) String() string {
	var parts []string
	if s.Comment != "" {
		parts = append(parts, "// @type "+s.Comment)
	}

	for _, decl := range slices.SortFunc(s.Declarations, sortDeclarations) {
		parts = append(parts, decl.Comments(TLDeclarationTypeConstructor)...)
		parts = append(parts, decl.String())
	}

	return strings.Join(parts, "\n")
}

type TLEnumDeclaration struct {
	Comment      string
	Declarations []TLDeclaration // must be sorted by name
}

func (s TLEnumDeclaration) String() (res string) {
	var parts []string
	if s.Comment != "" {
		parts = append(parts, "// @type "+s.Comment)
	}

	for _, decl := range slices.SortFunc(s.Declarations, sortDeclarations) {
		parts = append(parts, decl.Comments(TLDeclarationTypeEnum)...)
		parts = append(parts, decl.String())
	}

	return strings.Join(parts, "\n")
}

func methodsString(methods []TLDeclaration, group string) (res string) {
	var parts []string

	if group != "" {
		group += "."
	}

	for _, decl := range slices.SortFunc(methods, sortDeclarations) {
		parts = append(parts, decl.Comments(TLDeclarationTypeMethod)...)
		parts = append(parts, decl.String())
	}

	return strings.Join(parts, "\n")
}
