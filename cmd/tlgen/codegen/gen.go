// Copyright (c) 2022-2024 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package codegen

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"github.com/quenbyako/ext/slices"

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

func goify(name string, public bool) (res string) {
	delim := strcase.ToDelimited(name, '|')
	for i, item := range strings.Split(delim, "|") {
		item = strings.ToLower(item)
		if slices.Contains(capitalizePatterns, item) {
			item = strings.ToUpper(item)
		}

		itemRunes := []rune(item)

		if i == 0 && !public {
			// cause `aPI`, `uRL`, etc looks shitty ¯\_(ツ)_/¯
			itemRunes = []rune(strings.ToLower(item))
		} else {
			itemRunes[0] = unicode.ToUpper(itemRunes[0])
		}

		res += string(itemRunes)
	}

	return res
}

func GetFieldName(name schema.TLName) (res string) {
	if name.Namespace != "" {
		res += goify(name.Namespace, true)
	}
	return res + goify(name.Key, true)
}

func GetPredictName(name schema.TLName) (res string) {
	if name.Namespace != "" {
		res += goify(name.Namespace, true)
	}
	return /*"TL" + */ res + goify(name.Key, true) + "Predict"
}

func GetTypeName(name schema.TLName) (res string) {
	if name.Namespace != "" {
		res += goify(name.Namespace, true)
	}
	return /*"TL" + */ res + goify(name.Key, true)
}

func GetEnumName(name schema.TLName) (res string) {
	if name.Namespace != "" {
		res += goify(name.Namespace, true)
	}
	return /*"TL" + */ res + goify(name.Key, true)
}

func createTypeCrcFunc(typ string, crc uint32) *jen.Statement {
	hex := fmt.Sprintf("0x%x", crc)
	return jen.Func().
		Params(jen.Op("*").Id(typ)).
		Id("CRC").
		Params().
		Uint32().
		Block(
			jen.Return(jen.Id(hex)),
		)
}

func createIfaceFunc(typ string, ifaceMethod string) *jen.Statement {
	return jen.Func().
		Params(jen.Op("*").Id(typ)).
		Id(ifaceMethod).
		Params().
		Block()
}

func generateObject(ifaceMethod string, m schema.TLObject) (stmt *jen.Statement, objName string) {
	stmt = &jen.Statement{}
	if m.Comment != "" {
		stmt = stmt.Comment(m.Comment).Line()
	}

	typName := GetPredictName(m.Name)

	typ := stmt.Type().
		Id(typName).
		Struct(
			slices.Remap(m.Params, func(p schema.TLParam) jen.Code {
				return generateField(p)
			})...,
		)

	crc := createTypeCrcFunc(typName, m.CRC)
	res := jen.Add(typ, jen.Line(), crc)
	if ifaceMethod != "" {
		res = res.Add(jen.Line(), createIfaceFunc(typName, ifaceMethod))
	}

	return res, typName
}

func generateField(p schema.TLParam) *jen.Statement {
	var stmt *jen.Statement
	switch p := p.(type) {
	case schema.TLBitflagParam:
		tag := fmt.Sprintf("%v,bitflag", p.Name)
		stmt = jen.Id("_").Struct().Tag(map[string]string{"tl": tag})

	case schema.TLRequiredParam:
		stmt = jen.Id(GetFieldName(schema.TLName{Key: p.Name})).Add(generateFieldType(p.Type, false))

	case schema.TLOptionalParam:
		tag := fmt.Sprintf(",omitempty:%v:%v", p.FlagTrigger, p.BitTrigger)
		stmt = jen.Id(GetFieldName(schema.TLName{Key: p.Name})).Add(generateFieldType(p.Type, true)).Tag(map[string]string{"tl": tag})

	case schema.TLTriggerParam:
		tag := fmt.Sprintf(",omitempty:%v:%v,implicit", p.FlagTrigger, p.BitTrigger)
		stmt = jen.Id(GetFieldName(schema.TLName{Key: p.Name})).Bool().Tag(map[string]string{"tl": tag})

	default:
		panic("unknown parameter type")
	}

	if comment := p.GetComment(); comment != "" {
		stmt = stmt.Comment(comment)
	}

	return stmt
}

func generateFieldType(t schema.TLType, isOptional bool) *jen.Statement {
	switch t := t.(type) {
	case schema.TLTypeCommon:
		if isDefaultType(schema.TLName(t.TLName)) && isOptional {
			return jen.Op("*").Add(generateFieldTypeCommon(t))
		}
		return generateFieldTypeCommon(t)
	case schema.TLTypeVector:
		return jen.Index().Add(generateFieldTypeCommon(t))
	default:
		panic("unknown type")
	}
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
	default:
		if !typ.Name().IsInterface() {
			panic(fmt.Sprintf("incorrect type name: %v", typ))
		}
		return jen.Id(GetTypeName(typ.Name()))
	}
}

var (
	typeBytes  = schema.TLName{Key: "bytes"}
	typeDouble = schema.TLName{Key: "double"}
	typeInt    = schema.TLName{Key: "int"}
	typeLong   = schema.TLName{Key: "long"}
	typeString = schema.TLName{Key: "string"}
	typeBool   = schema.TLName{Key: "Bool"}
)

func isDefaultType(typeName schema.TLName) bool {
	switch typeName {
	case typeBytes, typeDouble, typeInt, typeLong, typeString, typeBool:
		return true
	default:
		return false
	}
}

func generateObjects(name schema.TLName, objects schema.TypeTLObjects) *jen.Statement {
	ifaceName := GetTypeName(name)
	ifaceMethod := "_" + ifaceName

	iface := &jen.Statement{}
	if objects.Comment != "" {
		iface = iface.Comment(objects.Comment).Line()
	}
	iface = iface.Type().Id(ifaceName).Interface(
		jen.Qual("github.com/xelaj/tl", "Object"),
		jen.Id(ifaceMethod).Params(),
	)

	checks := []jen.Code{}
	implementations := []*jen.Statement{}

	for _, v := range objects.Objects {
		impl, typeName := generateObject(ifaceMethod, v)
		implementations = append(implementations, impl)
		checks = append(checks, jen.Id("_").Id(ifaceName).Op("=").Call(jen.Op("*").Id(typeName)).Call(jen.Nil()))
	}

	res := jen.Add(iface, jen.Line(), jen.Var().Defs(checks...).Line())

	for _, impl := range implementations {
		res = res.Add(impl, jen.Line())
	}

	return res.Line()
}

func generateEnum(name schema.TLName, objects schema.EnumTLObjects) *jen.Statement {
	defName := GetEnumName(name)
	def := &jen.Statement{}
	if objects.Comment != "" {
		def = def.Comment(objects.Comment).Line()
	}
	def = jen.Type().Id(defName).Uint32().Line()

	constants := make([]jen.Code, len(objects.Objects))
	for i, obj := range objects.Objects {
		hex := fmt.Sprintf("0x%x", obj.CRC)

		c := jen.Id(GetPredictName(obj.Name)).Id(defName).Op("=").Id(hex)
		if obj.Comment != "" {
			c = c.Comment(obj.Comment)
		}
		constants[i] = c
	}

	return jen.Add(def, jen.Line(), jen.Const().Defs(constants...).Line(), jen.Line())
}

func generateRequestType(group string, obj schema.TLObject) *jen.Statement {
	funcName := schema.TLName{Namespace: group, Key: obj.Name.Key}
	obj.Name = schema.TLName{Namespace: group, Key: obj.Name.Key + "Request"}

	reqObj, reqName := generateObject("", obj)
	methodFunc := genFunction(funcName, reqName, obj.PolyParams, obj.Type)

	return jen.Add(reqObj, jen.Line(), jen.Line(), methodFunc, jen.Line())
}

// output:
//
//	func request[IN, OUT any](ctx context.Context, m Requester, in *IN, out *OUT) error {
//		if msg, err := tl.Marshal(in); err != nil {
//			return fmt.Errorf("marshaling: %w", err)
//		} else if respRaw, err := m.MakeRequest(ctx, msg); err != nil {
//			return fmt.Errorf("sending: %w", err)
//		} else if err := Unmarshal(respRaw, out); err != nil {
//			return fmt.Errorf("got invalid response type: %w", err)
//		}
//
//		return nil
//	}
func generateRequestBareFunction() *jen.Statement {
	requester := jen.Type().Id("Requester").Interface(
		jen.Id("MakeRequest").Params(jen.Id("ctx").Qual("context", "Context"), jen.Id("msg").Index().Byte()).Params(jen.Index().Byte(), jen.Error()),
	)

	request := jen.Id(`func request[IN, OUT any](ctx context.Context, m Requester, in *IN, out *OUT) error {
	if msg, err := tl.Marshal(in); err != nil {
		return fmt.Errorf("marshaling: %w", err)
	} else if respRaw, err := m.MakeRequest(ctx, msg); err != nil {
		return fmt.Errorf("sending: %w", err)
	} else if err := Unmarshal(respRaw, out); err != nil {
		return fmt.Errorf("got invalid response type: %w", err)
	}

	return nil
}`)

	return requester.Line().Add(request).Line()
}

// output:
//
//	func MethodName[Response any](ctx context.Context, m Requester, i MethodNameRequest) (Response, error) {
//		var res Response
//		return res, request(ctx, m, &i, &res)
//	}
func genFunction(funcName schema.TLName, requestType string, polyTypes schema.TLParams, returns schema.TLType) *jen.Statement {
	returnType := generateFieldType(returns, false)

	types := make([]jen.Code, len(polyTypes))
	for i, t := range polyTypes {
		types[i] = jen.Id(GetTypeName(schema.TLName{Key: t.GetName()})).Any()
	}

	return jen.Func().
		Id(GetFieldName(funcName)).
		Types(types...).
		Params(
			jen.Id("ctx").Qual("context", "Context"),
			jen.Id("m").Id("Requester"),
			jen.Id("i").Id(requestType),
		).
		Params(
			returnType,
			jen.Error(),
		).
		Block(
			jen.Var().Id("res").Add(returnType),
			jen.Return(
				jen.Id("res"),
				jen.Id("request").Call(
					jen.Id("ctx"),
					jen.Id("m"),
					jen.Op("&").Id("i"),
					jen.Op("&").Id("res"),
				),
			),
		)
}
