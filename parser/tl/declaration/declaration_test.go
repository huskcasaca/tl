package declaration_test

import (
	. "github.com/xelaj/tl/parser/tl/declaration"
	"github.com/xelaj/tl/parser/tl/lexer"
	"testing"

	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/assert"
)

type TestCase interface {
	Name() string
	Run(t *testing.T)
}

type TcaseParseIt[T any] struct {
	name    string
	in      string
	want    T
	wantErr assert.ErrorAssertionFunc
}

func (tt TcaseParseIt[T]) Name() string { return tt.name }

func (tt TcaseParseIt[T]) Run(t *testing.T) {
	tt.wantErr = noErrAsDefault(tt.wantErr)

	parser := participle.MustBuild[T](
		participle.Lexer(lexer.NewDefinition()),
	)

	got, err := parser.ParseString("test.tl", tt.in)
	if !tt.wantErr(t, err) || err != nil {
		return
	}

	assert.Equal(t, tt.want, *got)
}

func TestParseIt(t *testing.T) {
	for _, tt := range []TestCase{
		TcaseParseIt[Argument]{
			name: "args",
			in:   "user_id:long",
			want: Argument{
				Ident: ArgIdent{
					Ident: "user_id",
				},
				Type: Type{
					Ident: TypeIdent{Ident: "long"},
				},
			},
		},
		TcaseParseIt[Declaration]{
			name: "declaration",
			in:   "a#00000000 flags:# underscore_var:underscore_type var:type complexVar:flags.2?Vector<Vector<int>> = F",
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
		TcaseParseIt[Program]{
			name: "program",
			in: `
a#123 = B;
c#456 = D;
---functions---

// another comment

// comment
e#789 = F;
g#012 = H;

---types---

i#345 = J;
k#678 = L;

`,
			want: Program{
				Entries: []ProgramEntry{
					{Newline: true},
					{Declaration: &Declaration{
						Combinator: "a#123",
						Result: Type{
							Ident: TypeIdent{Ident: "B"},
						},
					}},
					{Declaration: &Declaration{
						Combinator: "c#456",
						Result: Type{
							Ident: TypeIdent{Ident: "D"},
						},
					}},

					{FuncDecl: true},
					{Newline: true},
					{Comment: stringPtr("// another comment")},
					{Newline: true},
					{Comment: stringPtr("// comment")},
					{Declaration: &Declaration{
						Combinator: "e#789",
						Result: Type{
							Ident: TypeIdent{Ident: "F"},
						},
					}},
					{Declaration: &Declaration{
						Combinator: "g#012",
						Result: Type{
							Ident: TypeIdent{Ident: "H"},
						},
					}},
					{Newline: true},

					{TypeDecl: true},
					{Newline: true},
					{Declaration: &Declaration{
						Combinator: "i#345",
						Result: Type{
							Ident: TypeIdent{Ident: "J"},
						},
					}},
					{Declaration: &Declaration{
						Combinator: "k#678",
						Result: Type{
							Ident: TypeIdent{Ident: "L"},
						},
					}},
					{Newline: true},
				},
			},
		},
	} {
		t.Run(tt.Name(), tt.Run)
	}
}

func noErrAsDefault(e assert.ErrorAssertionFunc) assert.ErrorAssertionFunc {
	if e == nil {
		return assert.NoError
	}

	return e
}

func stringPtr(s string) *string { return &s }
