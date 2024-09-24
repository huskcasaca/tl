package program_test

import (
	"github.com/xelaj/tl/parser/typelang/lexer"
	"testing"

	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/assert"

	. "github.com/xelaj/tl/parser/typelang/program"
)

type TestCase interface {
	Name() string
	Run(t *testing.T)
}

type ParserTestCase[T any] struct {
	name    string
	input   string
	want    T
	wantErr assert.ErrorAssertionFunc
}

func (tt ParserTestCase[T]) Name() string { return tt.name }

func (tt ParserTestCase[T]) Run(t *testing.T) {
	tt.wantErr = noErrAsDefault(tt.wantErr)

	parser := participle.MustBuild[T](
		participle.Lexer(lexer.NewDefinition()),
	)

	got, err := parser.ParseString("test.tl", tt.input)
	if !tt.wantErr(t, err) || err != nil {
		return
	}

	assert.Equal(t, tt.want, *got)
}

func noErrAsDefault(e assert.ErrorAssertionFunc) assert.ErrorAssertionFunc {
	if e == nil {
		return assert.NoError
	}

	return e
}

func stringPtr(s string) *string { return &s }

func TestParseProgram(t *testing.T) {
	for _, tt := range []TestCase{
		ParserTestCase[Program]{
			name: "program",
			input: `
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
