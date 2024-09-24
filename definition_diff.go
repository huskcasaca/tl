package tl

import "github.com/quenbyako/ext/slices"

type DefinitionDiff struct {
	Added   []Definition
	Removed []Name
}

func (d DefinitionDiff) Patch(a []Definition) []Definition {
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

	return slices.SortFunc(a, SortDeclarations)
}

type DiffEnum struct {
	Comment string
	Changes DefinitionDiff
}
