package schema

import "github.com/quenbyako/ext/slices"

type TLObjectsDiff struct {
	Added   []TLDeclaration
	Removed []TLName
}

func (d TLObjectsDiff) Patch(a []TLDeclaration) []TLDeclaration {
	for _, add := range d.Added {
		a = append(a, add)
	}

	for _, rem := range d.Removed {
		for i, obj := range a {
			if cmpDeclName(obj.Name, rem) == 0 {
				a = append(a[:i], a[i+1:]...)
				break
			}
		}
	}

	return slices.SortFunc(a, sortDeclarations)
}

type DiffEnum struct {
	Comment string
	Changes TLObjectsDiff
}

func (a TLEnumDeclaration) Diff(b TLEnumDeclaration) (res DiffEnum) {
	if a.Comment != b.Comment {
		res.Comment = b.Comment
	}

	panic("unimplemented")
}

func (d DiffEnum) Patch(a TLEnumDeclaration) TLEnumDeclaration {
	if d.Comment != "" {
		a.Comment = d.Comment
	}

	a.Declarations = d.Changes.Patch(a.Declarations)

	return a
}
