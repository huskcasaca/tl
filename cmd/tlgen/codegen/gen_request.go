package codegen

import (
	"github.com/dave/jennifer/jen"
	"github.com/quenbyako/ext/slices"
	"github.com/xelaj/tl"
	"github.com/xelaj/tl/cmd/tlgen/util"
	"github.com/xelaj/tl/schema"
)

func getRequestName(name schema.TLName) (res string) {
	if name.Namespace != "" {
		res += getGoName(name.Namespace, true)
	}
	return /*"TL" + */ res + getGoName(name.Key, true) + "Request"
}

func generateRequest(m schema.TLDeclaration) (ret *jen.Statement, objName string) {
	ret = &jen.Statement{}
	if m.Comment != "" {
		ret = ret.Comment(m.Comment).Line()
	}

	requestTypeName := getRequestName(m.Name)

	ret = ret.Type().
		Add(generateGenericTypes(requestTypeName, m.PolyParams)).
		Struct(
			slices.Remap(m.Params, func(p schema.TLParam) jen.Code {
				return generateField(p)
			})...,
		)

	ret = ret.Line()
	ret = ret.Add(generateTypeCrcFunctions(generateGenericNames(requestTypeName, m.PolyParams), m.CRC))

	return ret, requestTypeName
}

func generateRequestType(funcReqObj schema.TLDeclaration) *jen.Statement {
	funcName := schema.TLName{Namespace: funcReqObj.Name.Namespace, Key: funcReqObj.Name.Key}
	//funcReqObj.Name = schema.TLName{Namespace: funcReqObj.Name.Namespace, Key: funcReqObj.Name.Key + "Request"}

	requestObjJens, requestTypeName := generateRequest(funcReqObj)

	requestFuncJens := generateRequestFunction(funcName, requestTypeName, funcReqObj.PolyParams, funcReqObj.Type)

	return jen.Add(requestObjJens, jen.Line(), jen.Line(), requestFuncJens, jen.Line())
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
func generateMakeRequesterFunc() *jen.Statement {
	return jen.
		Type().
		Id("Requester").
		Interface(
			jen.Id("MakeRequest").Params(jen.Id("ctx").Qual("context", "Context"), jen.Id("msg").Index().Byte()).Params(jen.Index().Byte(), jen.Error()),
		).
		Line().
		Func().
		Id("request").
		Types(
			jen.Id("IN").Any(),
			jen.Id("OUT").Any(),
		).
		Params(
			jen.Id("ctx").Qual("context", "Context"),
			jen.Id("m").Id("Requester"),
			jen.Id("in").Op("*").Id("IN"),
			jen.Id("out").Op("*").Id("OUT"),
		).
		Params(
			jen.Error(),
		).
		Block(
			jen.If(
				jen.List(jen.Id("msg"), jen.Err()).Op(":=").Qual(util.GetFunctionPathName(tl.Marshal)).Call(jen.Id("in")),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(jen.Qual("fmt", "Errorf").Call(jen.Lit("marshaling: %w"), jen.Err())),
			).Else().If(
				jen.List(jen.Id("respRaw"), jen.Err()).Op(":=").Id("m").Dot("MakeRequest").Call(jen.Id("ctx"), jen.Id("msg")),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(jen.Qual("fmt", "Errorf").Call(jen.Lit("sending: %w"), jen.Err())),
			).Else().If(
				jen.Err().Op(":=").Qual(util.GetFunctionPathName(tl.Unmarshal)).Call(jen.Id("respRaw"), jen.Id("out")),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(jen.Qual("fmt", "Errorf").Call(jen.Lit("got invalid response type: %w"), jen.Err())),
			),
			jen.Return(jen.Nil()),
		).Line()
}

func getRequestFuncName(name schema.TLName) (res string) {
	if name.Namespace != "" {
		res += getGoName(name.Namespace, true)
	}
	return "Make" + res + getGoName(name.Key, true) + "Request"
}

// output:
//
//	func MethodName[Response any](ctx context.Context, m Requester, i MethodNameRequest) (Response, error) {
//		var res Response
//		return res, request(ctx, m, &i, &res)
//	}
func generateRequestFunction(funcName schema.TLName, requestType string, polyParams schema.TLParams, returns schema.TLType) *jen.Statement {
	returnType := generateFieldType(returns, false)

	return jen.Func().
		Add(generateGenericTypes(getRequestFuncName(funcName), polyParams)).
		Params(
			jen.Id("ctx").Qual("context", "Context"),
			jen.Id("m").Id("Requester"),
			jen.Id("i").Add(generateGenericNames(requestType, polyParams)),
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
