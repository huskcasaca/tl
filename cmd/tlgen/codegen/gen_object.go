package codegen

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"github.com/quenbyako/ext/slices"
	"github.com/xelaj/tl"
	"github.com/xelaj/tl/cmd/tlgen/util"
	"strings"
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

func getFieldName(name tl.Name) (res string) {
	if name.Namespace != "" {
		res += getGoName(name.Namespace, true)
	}
	return res + getGoName(name.Key, true)
}

func getPredictName(name tl.Name) (res string) {
	if name.Namespace != "" {
		res += getGoName(name.Namespace, true)
	}
	return /*"TL" + */ res + getGoName(name.Key, true) + "Predict"
}

func getTypeName(name tl.Name) (res string) {
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

func generatePredicts(method string, m tl.Declaration) (ret *jen.Statement, objName string) {
	ret = &jen.Statement{}
	if m.Comment != "" {
		ret = ret.Comment(m.Comment).Line()
	}

	predictTypeName := getPredictName(m.Name)

	ret = ret.Type().
		Add(generateGenericTypes(predictTypeName, m.OptParams)).
		Struct(
			slices.Remap(m.Params, func(p tl.Param) jen.Code {
				return generateField(p)
			})...,
		)

	ret = ret.Line()
	ret = ret.Add(generateTypeCrcFunctions(generateGenericNames(predictTypeName, m.OptParams), m.CRC))
	ret = ret.Line()
	ret = ret.Add(generateInterfaceFunctions(generateGenericNames(predictTypeName, m.OptParams), method))

	return ret, predictTypeName
}

func generateGenericTypes(name string, polys tl.Params) *jen.Statement {
	genericsTypes := make([]jen.Code, len(polys))
	for i, t := range polys {
		genericsTypes[i] = jen.Id(getTypeName(tl.Name{Key: t.GetName()})).Add(generateFieldType(t.GetType(), true))
	}
	return jen.Id(name).Types(genericsTypes...)
}

func generateGenericNames(name string, polys tl.Params) *jen.Statement {
	genericsNames := make([]jen.Code, len(polys))
	for i, t := range polys {
		genericsNames[i] = jen.Id(getTypeName(tl.Name{Key: t.GetName()}))
	}
	return jen.Id(name).Types(genericsNames...)
}

func generateField(p tl.Param) *jen.Statement {
	ret := generateFieldBase(p)
	if comment := p.GetComment(); comment != "" {
		ret = ret.Comment(comment)
	}
	return ret
}

func generateFieldBase(p tl.Param) *jen.Statement {
	switch p := p.(type) {
	case tl.BitflagParam:
		return generateBitflagField(p)
	case tl.RequiredParam:
		return generateRequiredField(p)
	case tl.OptionalParam:
		return generateOptionalField(p)
	case tl.TriggerParam:
		return generateTriggerField(p)
	default:
		panic("unknown parameter type")
	}
}

func generateBitflagField(p tl.BitflagParam) *jen.Statement {
	tag := fmt.Sprintf("%v,%v", p.Name, tl.IsBitflagFlag)
	return jen.Id("_").Struct().Tag(map[string]string{tl.TagName: tag})
}

func generateRequiredField(p tl.RequiredParam) *jen.Statement {
	tag := fmt.Sprintf("%v", p.Name)
	return jen.Id(getFieldName(tl.Name{Key: p.Name})).Add(generateFieldType(p.Type, false)).Tag(map[string]string{tl.TagName: tag})
}

func generateOptionalField(p tl.OptionalParam) *jen.Statement {
	tag := fmt.Sprintf("%v,%v:%v:%v", p.Name, tl.OmitemptyPrefix, p.FlagTrigger, p.BitTrigger)
	return jen.Id(getFieldName(tl.Name{Key: p.Name})).Add(generateFieldType(p.Type, true)).Tag(map[string]string{tl.TagName: tag})
}

func generateTriggerField(p tl.TriggerParam) *jen.Statement {
	tag := fmt.Sprintf("%v,%v:%v:%v,%v", p.Name, tl.OmitemptyPrefix, p.FlagTrigger, p.BitTrigger, tl.ImplicitFlag)
	return jen.Id(getFieldName(tl.Name{Key: p.Name})).Bool().Tag(map[string]string{tl.TagName: tag})
}

func generateFieldType(t tl.Type, isOptional bool) *jen.Statement {
	ret := generateFieldTypeCommon(t)
	switch t.Name {
	case typeBytes, typeDouble, typeInt, typeLong, typeString, typeBool:
		if t.SubType != nil {
			panic(fmt.Sprintf("incorrect default type: %v", t))
		}
		if isOptional {
			ret = jen.Op("*").Add(ret)
		}
	case typeVectorUc, typeVectorLc:
		if t.SubType == nil {
			panic(fmt.Sprintf("incorrect vector type: %v", t))
		}
		ret = ret.Add(generateFieldType(*t.SubType, false))
	default:
		if t.SubType != nil {
			ret = ret.Index(generateFieldType(*t.SubType, false))
		}
	}
	return ret
}

func generateFieldTypeCommon(typ tl.Type) *jen.Statement {
	switch typ.Name {
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
		return jen.Id(getTypeName(typ.Name))
	}
}

var (
	typeBytes    = tl.Name{Key: "bytes"}
	typeDouble   = tl.Name{Key: "double"}
	typeInt      = tl.Name{Key: "int"}
	typeLong     = tl.Name{Key: "long"}
	typeString   = tl.Name{Key: "string"}
	typeBool     = tl.Name{Key: "Bool"}
	typeAny      = tl.Name{Key: "Type"}
	typeVectorUc = tl.Name{Key: "Vector"}
	typeVectorLc = tl.Name{Key: "vector"}
	typeInt128   = tl.Name{Key: "int128"}
	typeInt256   = tl.Name{Key: "int256"}
	typeObject   = tl.Name{Key: "Object"}
)

func generateObjects(td tl.TypeDeclaration) *jen.Statement {
	typeName := getTypeName(td.Type.Name)
	typeMethod := "_" + typeName

	ret := &jen.Statement{}
	if td.Comment != "" {
		ret = ret.Comment(td.Comment).Line()
	}
	ret = ret.Type().Id(typeName).Interface(
		generateFieldType(tl.AnyType, false),
		jen.Id(typeMethod).Params(),
	)

	checkJens := []jen.Code{}
	objectJens := []*jen.Statement{}

	for _, obj := range td.Declarations {
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
