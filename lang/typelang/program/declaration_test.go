package program_test

import (
	. "github.com/xelaj/tl/lang/typelang/program"
	"testing"
)

func TestParseDeclaration(t *testing.T) {
	for _, tt := range []TestCase{
		ParserTestCase[Declaration]{
			name:  "declaration",
			input: "a#00000000 flags:# underscore_var:underscore_type var:type complexVar:flags.2?Vector<Vector<int>> = F",
			want: Declaration{
				Combinator: "a#00000000",
				Args: []Argument{
					{
						Ident: ArgIdent{
							Ident: "flags",
						},
						Type: Type{
							Ident: TypeIdent{Ident: "#"},
						},
					},
					{
						Ident: ArgIdent{
							Ident: "underscore_var",
						},
						Type: Type{
							Ident: TypeIdent{Ident: "underscore_type"},
						},
					},
					{
						Ident: ArgIdent{
							Ident: "var",
						},
						Type: Type{
							Ident: TypeIdent{Ident: "type"},
						},
					},
					{
						Ident: ArgIdent{
							Ident: "complexVar",
						},
						Flag: &Flag{
							Ident: "flags",
							Index: 2,
						},
						Type: Type{
							Ident: TypeIdent{Ident: "Vector"},
							SubTypes: []Type{
								{
									Ident: TypeIdent{Ident: "Vector"},
									SubTypes: []Type{
										{
											Ident: TypeIdent{Ident: "int"},
										},
									},
								},
							},
						},
					},
				},
				Result: Type{
					Ident: TypeIdent{Ident: "F"},
				},
			},
		},
	} {
		t.Run(tt.Name(), tt.Run)
	}
}
