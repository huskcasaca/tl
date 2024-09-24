package tl

import (
	"cmp"
	"fmt"
	hash "hash/crc32"
	"strings"
)

type Definition struct {
	Name       Name
	CRC        CRC32
	Category   Category
	Params     Params
	OptParams  Params
	Type       Type
	Comment    string
	RetComment string
}

type Name struct {
	Namespace string
	Key       string
}

func NameFromString(s string) Name {
	groups := strings.Split(s, ".")
	var namespace string
	key := groups[0]
	if len(groups) > 1 {
		namespace = groups[0]
		key = groups[1]
	}

	return Name{Namespace: namespace, Key: key}
}

func (o Name) String() string {
	if o.Namespace != "" {
		return o.Namespace + "." + o.Key
	}

	return o.Key
}

func (o Name) Cmp(b Name) int { return cmpDeclName(o, b) }

func cmpDeclName(a, b Name) int {
	if c := cmp.Compare(a.Namespace, b.Namespace); c != 0 {
		return c
	} else if c := cmp.Compare(a.Key, b.Key); c != 0 {
		return c
	} else {
		return 0
	}
}

func SortDeclarations(a, b Definition) int { return cmpDeclName(a.Name, b.Name) }

// GenerateCRC32 generates crc32 of this declaration
func (d *Definition) GenerateCRC32() uint32 {
	if d.CRC != 0 {
		return d.CRC
	}

	filtered := make(Params, 0, len(d.Params))
	for _, item := range d.Params {
		if _, ok := item.(TriggerParam); !ok {
			filtered = append(filtered, item)
		}
	}

	fieldsStr := ""
	if len(filtered) > 0 {
		fieldsStr = " " + filtered.String()
	}

	return hash.ChecksumIEEE([]byte(fmt.Sprintf("%v%v = %v;", d.Name.String(), fieldsStr, d.Type)))
}
