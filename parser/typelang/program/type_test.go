package program_test

import (
	. "github.com/xelaj/tl/parser/typelang/program"
	"testing"
)

func TestParseType(t *testing.T) {
	for _, tt := range []TestCase{
		ParserTestCase[Type]{
			name:  "simple",
			input: "int",
			want: Type{
				Ident: TypeIdent{Ident: "int"},
			},
		},
		ParserTestCase[Type]{
			name:  "simple",
			input: "Type",
			want: Type{
				Ident: TypeIdent{Ident: "Type"},
			},
		},
		ParserTestCase[Type]{
			name:  "boxed",
			input: "Vector<int>",
			want: Type{
				Ident: TypeIdent{Ident: "Vector"},
				SubTypes: []Type{
					{
						Ident: TypeIdent{Ident: "int"},
					},
				},
			},
		},
		ParserTestCase[Type]{
			name:  "boxed",
			input: "Vector<Vector<int>>",
			want: Type{
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
	} {
		t.Run(tt.Name(), tt.Run)
	}
}
