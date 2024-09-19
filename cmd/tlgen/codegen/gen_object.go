package codegen

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"github.com/quenbyako/ext/slices"
	"github.com/xelaj/tl"
	"github.com/xelaj/tl/cmd/tlgen/util"
	"strings"

	"github.com/xelaj/tl/schema"
)

// Some terms (id, api, url etc.) it's important to write in uppercase letters
// or lowercase only (e.g, you can't write Id, or Api, only ID, API, URL, etc.)
var capitalizePatterns = []string{
	"api",
	"id",
	"json",
	"p2p",
	"sha",
	"srp",
	"ttl",
	"url",
}

func getGoName(name string, public bool) string {
	parts := strings.Split(strcase.ToDelimited(name, '|'), "|")
	for i, part := range parts {
		part = strings.ToLower(part)
		if slices.Contains(capitalizePatterns, part) {
			part = strings.ToUpper(part)
		} else if i == 0 && !public {
			part = strings.ToLower(part)
		} else {
			part = strings.Title(part)
		}
		parts[i] = part
	}
	return strings.Join(parts, "")
}

func getFieldName(name schema.TLName) (res string) {
	if name.Namespace != "" {
		res += getGoName(name.Namespace, true)
	}
	return res + getGoName(name.Key, true)
}

func getPredictName(name schema.TLName) (res string) {
	if name.Namespace != "" {
		res += getGoName(name.Namespace, true)
	}
	return /*"TL" + */ res + getGoName(name.Key, true) + "Predict"
}

func getTypeName(name schema.TLName) (res string) {
	if name.Namespace != "" {
		res += getGoName(name.Namespace, true)
	}
	return /*"TL" + */ res + getGoName(name.Key, true)
}

func generateTypeCrcFunctions(typ jen.Code, crc uint32) *jen.Statement {
	hex := fmt.Sprintf("0x%08x", crc)
	return jen.Func().
		Params(jen.Op("*").Add(typ)).
		Id("CRC").
		Params().
		Uint32().
		Block(
			jen.Return(jen.Id(hex)),
		)
}

func generateInterfaceFunctions(typ jen.Code, method string) *jen.Statement {
	return jen.Func().
		Params(jen.Op("*").Add(typ)).
		Id(method).
		Params().
		Block()
}

func generatePredicts(method string, m schema.TLDeclaration) (ret *jen.Statement, objName string) {
	ret = &jen.Statement{}
	if m.Comment != "" {
		ret = ret.Comment(m.Comment).Line()
	}

	predictTypeName := getPredictName(m.Name)

	ret = ret.Type().
		Add(generateGenericTypes(predictTypeName, m.PolyParams)).
		Struct(
			slices.Remap(m.Params, func(p schema.TLParam) jen.Code {
				return generateField(p)
			})...,
		)

	ret = ret.Line()
	ret = ret.Add(generateTypeCrcFunctions(generateGenericNames(predictTypeName, m.PolyParams), m.CRC))
	ret = ret.Line()
	ret = ret.Add(generateInterfaceFunctions(generateGenericNames(predictTypeName, m.PolyParams), method))

	return ret, predictTypeName
}

func generateGenericTypes(name string, polys schema.TLParams) *jen.Statement {
	genericsTypes := make([]jen.Code, len(polys))
	for i, t := range polys {
		genericsTypes[i] = jen.Id(getTypeName(schema.TLName{Key: t.GetName()})).Add(generateFieldType(t.GetType(), true))
	}
	return jen.Id(name).Types(genericsTypes...)
}

func generateGenericNames(name string, polys schema.TLParams) *jen.Statement {
	genericsNames := make([]jen.Code, len(polys))
	for i, t := range polys {
		genericsNames[i] = jen.Id(getTypeName(schema.TLName{Key: t.GetName()}))
	}
	return jen.Id(name).Types(genericsNames...)
}

func generateField(p schema.TLParam) *jen.Statement {
	ret := generateFieldBase(p)
	if comment := p.GetComment(); comment != "" {
		ret = ret.Comment(comment)
	}
	return ret
}

func generateFieldBase(p schema.TLParam) *jen.Statement {
	switch p := p.(type) {
	case schema.TLBitflagParam:
		return generateBitflagField(p)
	case schema.TLRequiredParam:
		return generateRequiredField(p)
	case schema.TLOptionalParam:
		return generateOptionalField(p)
	case schema.TLTriggerParam:
		return generateTriggerField(p)
	default:
		panic("unknown parameter type")
	}
}

func generateBitflagField(p schema.TLBitflagParam) *jen.Statement {
	tag := fmt.Sprintf("%v,%v", p.Name, tl.IsBitflagFlag)
	return jen.Id("_").Struct().Tag(map[string]string{tl.TagName: tag})
}

func generateRequiredField(p schema.TLRequiredParam) *jen.Statement {
	tag := fmt.Sprintf("%v", p.Name)
	return jen.Id(getFieldName(schema.TLName{Key: p.Name})).Add(generateFieldType(p.Type, false)).Tag(map[string]string{tl.TagName: tag})
}

func generateOptionalField(p schema.TLOptionalParam) *jen.Statement {
	tag := fmt.Sprintf("%v,%v:%v:%v", p.Name, tl.OmitemptyPrefix, p.FlagTrigger, p.BitTrigger)
	return jen.Id(getFieldName(schema.TLName{Key: p.Name})).Add(generateFieldType(p.Type, true)).Tag(map[string]string{tl.TagName: tag})
}

func generateTriggerField(p schema.TLTriggerParam) *jen.Statement {
	tag := fmt.Sprintf("%v,%v:%v:%v,%v", p.Name, tl.OmitemptyPrefix, p.FlagTrigger, p.BitTrigger, tl.ImplicitFlag)
	return jen.Id(getFieldName(schema.TLName{Key: p.Name})).Bool().Tag(map[string]string{tl.TagName: tag})
}

func generateFieldType(t schema.TLType, isOptional bool) *jen.Statement {
	ret := generateFieldTypeCommon(t)
	switch t.Name() {
	case typeBytes, typeDouble, typeInt, typeLong, typeString, typeBool:
		if len(t.Types()) != 0 {
			panic(fmt.Sprintf("incorrect default type: %v", t))
		}
		if isOptional {
			ret = jen.Op("*").Add(ret)
		}
	case typeVectorUc, typeVectorLc:
		if len(t.Types()) != 1 {
			panic(fmt.Sprintf("incorrect vector type: %v", t))
		}
		ret = ret.Add(generateFieldType(t.Types()[0], false))
	default:
		if len(t.Types()) != 0 {
			generics := make([]jen.Code, len(t.Types()))
			for _, t1 := range t.Types() {
				generics = append(generics, generateFieldType(t1, false))
			}
			ret = ret.Index(generics...)
		}
	}
	return ret
}

func generateFieldTypeCommon(typ schema.TLType) *jen.Statement {
	switch typ.Name() {
	case typeBytes:
		return jen.Index().Byte()
	case typeDouble:
		return jen.Float64()
	case typeInt:
		return jen.Int32()
	case typeLong:
		return jen.Int64()
	case typeString:
		return jen.String()
	case typeBool:
		return jen.Bool()
	case typeAny, typeObject:
		return jen.Qual(util.GetTypePathName((*tl.Any)(nil)))
	case typeVectorUc, typeVectorLc:
		return jen.Index()
	case typeInt128:
		return jen.Qual(util.GetTypePathName((*tl.Int128)(nil)))
	case typeInt256:
		return jen.Qual(util.GetTypePathName((*tl.Int256)(nil)))
	default:
		//if !typ.Name().IsInterface() {
		//	panic(fmt.Sprintf("incorrect type name: %v", typ))
		//}
		return jen.Id(getTypeName(typ.Name()))
	}
}

var (
	typeBytes    = schema.TLName{Key: "bytes"}
	typeDouble   = schema.TLName{Key: "double"}
	typeInt      = schema.TLName{Key: "int"}
	typeLong     = schema.TLName{Key: "long"}
	typeString   = schema.TLName{Key: "string"}
	typeBool     = schema.TLName{Key: "Bool"}
	typeAny      = schema.TLName{Key: "Type"}
	typeVectorUc = schema.TLName{Key: "Vector"}
	typeVectorLc = schema.TLName{Key: "vector"}
	typeInt128   = schema.TLName{Key: "int128"}
	typeInt256   = schema.TLName{Key: "int256"}
	typeObject   = schema.TLName{Key: "Object"}
)

func generateObjects(name schema.TLName, objects schema.TLTypeDeclaration) *jen.Statement {
	typeName := getTypeName(name)
	typeMethod := "_" + typeName

	ret := &jen.Statement{}
	if objects.Comment != "" {
		ret = ret.Comment(objects.Comment).Line()
	}
	ret = ret.Type().Id(typeName).Interface(
		generateFieldType(schema.TLAnyType, false),
		jen.Id(typeMethod).Params(),
	)

	checkJens := []jen.Code{}
	objectJens := []*jen.Statement{}

	for _, obj := range objects.Declarations {
		predictJens, predictTypeName := generatePredicts(typeMethod, obj)
		objectJens = append(objectJens, predictJens)
		checkJens = append(checkJens, jen.Id("_").Id(typeName).Op("=").Call(jen.Op("*").Id(predictTypeName)).Call(jen.Nil()))
	}

	ret = ret.Add(jen.Line(), jen.Var().Defs(checkJens...).Line())

	for _, jens := range objectJens {
		ret = ret.Add(jens, jen.Line())
	}

	return ret.Line()
}
