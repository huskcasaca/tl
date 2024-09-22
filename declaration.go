package tl

import (
	"cmp"
	"fmt"
	hash "hash/crc32"
	"strings"
)

type TLDeclaration struct {
	Comment    string
	Name       TLName
	CRC        uint32
	Params     TLParams
	PolyParams TLParams
	Type       TLType
}

type TLName struct {
	Namespace string
	Key       string
}

func GetTLNameFromString(s string) TLName {
	groups := strings.Split(s, ".")
	var namespace string
	key := groups[0]
	if len(groups) > 1 {
		namespace = groups[0]
		key = groups[1]
	}

	return TLName{Namespace: namespace, Key: key}
}

func (o TLName) String() string {
	if o.Namespace != "" {
		return o.Namespace + "." + o.Key
	}

	return o.Key
}

func (o TLName) IsInterface() bool { return isFirstRuneUpper(o.Key) }

func (o TLName) Cmp(b TLName) int { return cmpDeclName(o, b) }

func cmpDeclName(a, b TLName) int {
	if c := cmp.Compare(a.Namespace, b.Namespace); c != 0 {
		return c
	} else if c := cmp.Compare(a.Key, b.Key); c != 0 {
		return c
	} else {
		return 0
	}
}

func sortDeclarations(a, b TLDeclaration) int { return cmpDeclName(a.Name, b.Name) }

type TLDeclarationType uint8

const (
	TLDeclarationTypeUnknown TLDeclarationType = iota
	TLDeclarationTypeConstructor
	TLDeclarationTypeMethod
)

func (o TLDeclarationType) String() string {
	switch o {
	case TLDeclarationTypeConstructor:
		return "constructor"
	case TLDeclarationTypeMethod:
		return "method"
	default:
		return "<UNKNOWN>"
	}
}

// how CRC is calculated:
// first of all, constructor formats to canonical state, e.g.
// `user#12325 id:int         first_name:string    last_name:string = User`
// ↓↓↓
// `user id:int first_name:string last_name:string = User`
//
// all bit triggers are removing, e.g.
// `messages.clearRecentStickers#8999602d flags:# attached:flags.0?true = Bool`
// ↓↓↓
// `messages.clearRecentStickers flags:# = Bool`
//
// For vectors like `getSmthn items:Vector<int> = Bool` i still don't understand
// how to generate, cause it fails in real mtproto schema.
func (o *TLDeclaration) getCRC() uint32 {
	if o.CRC != 0 {
		return o.CRC
	}

	filtered := make(TLParams, 0, len(o.Params))
	for _, item := range o.Params {
		if _, ok := item.(TLTriggerParam); !ok {
			filtered = append(filtered, item)
		}
	}

	fieldsStr := ""
	if len(filtered) > 0 {
		fieldsStr = " " + filtered.String()
	}

	return hash.ChecksumIEEE([]byte(fmt.Sprintf("%v%v = %v;", o.Name.String(), fieldsStr, o.Type)))
}

func (o *TLDeclaration) String() string {
	fields := ""
	if len(o.Params) > 0 {
		fields = " " + o.Params.String()
	}

	return fmt.Sprintf("%v#%08x%v = %v;", o.Name.String(), o.getCRC(), fields, o.Type)
}

func (o *TLDeclaration) Comments(typ TLDeclarationType) []string {
	var res []string
	if o.Comment != "" {
		res = append(res, "// @"+typ.String()+" "+o.Comment)
	}

	return append(res, o.Params.Comments()...)
}
