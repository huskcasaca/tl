package program_test

import (
	. "github.com/xelaj/tl/lang/typelang/program"
	"testing"
)

func TestParseArgument(t *testing.T) {
	for _, tt := range []TestCase{
		ParserTestCase[Argument]{
			name:  "args",
			input: "user_id:long",
			want: Argument{
				Ident: ArgIdent{
					Ident: "user_id",
				},
				Type: Type{
					Ident: TypeIdent{Ident: "long"},
				},
			},
		},
	} {
		t.Run(tt.Name(), tt.Run)
	}
}
