package tl

import (
	"cmp"
	"fmt"
	hash "hash/crc32"
	"strings"
)

type Declaration struct {
	Comment   string
	Name      Name
	CRC       CRC32
	Category  Category
	Params    Params
	OptParams Params
	Type      Type
}

type Name struct {
	Namespace string
	Key       string
}

func ParseNameFromString(s string) Name {
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

func sortDeclarations(a, b Declaration) int { return cmpDeclName(a.Name, b.Name) }

// GenerateCRC32 generates crc32 of this declaration
func (o *Declaration) GenerateCRC32() uint32 {
	if o.CRC != 0 {
		return o.CRC
	}

	filtered := make(Params, 0, len(o.Params))
	for _, item := range o.Params {
		if _, ok := item.(TriggerParam); !ok {
			filtered = append(filtered, item)
		}
	}

	fieldsStr := ""
	if len(filtered) > 0 {
		fieldsStr = " " + filtered.String()
	}

	return hash.ChecksumIEEE([]byte(fmt.Sprintf("%v%v = %v;", o.Name.String(), fieldsStr, o.Type)))
}

func (o *Declaration) String() string {
	fields := ""
	if len(o.Params) > 0 {
		fields = " " + o.Params.String()
	}

	return fmt.Sprintf("%v#%08x%v = %v;", o.Name.String(), o.GenerateCRC32(), fields, o.Type)
}

func (o *Declaration) Comments(typ Category) []string {
	var res []string
	if o.Comment != "" {
		res = append(res, "// @"+typ.String()+" "+o.Comment)
	}

	return append(res, o.Params.Comments()...)
}
