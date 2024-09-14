package schema

import "github.com/quenbyako/ext/slices"

type TLDeclarationDiff struct {
	Added   []TLDeclaration
	Removed []TLName
}

func (d TLDeclarationDiff) Patch(a []TLDeclaration) []TLDeclaration {
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
	Changes TLDeclarationDiff
}
