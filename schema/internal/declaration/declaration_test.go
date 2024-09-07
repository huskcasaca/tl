// Copyright (c) 2022-2024 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package declaration_test

import (
	"testing"

	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/assert"

	. "github.com/xelaj/tl/schema/internal/declaration"
	"github.com/xelaj/tl/schema/internal/lexer"
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
				Term: ArgType{
					Simple: TypeIdent{Ident: "long"},
				},
			},
		},
		TcaseParseIt[Declaration]{
			name: "declaration",
			in:   "a#00000000 flags:# b_x:c_x d:e pipka:flags.2?Vector<popka> = F",
			want: Declaration{
				Combinator: "a#00000000",
				Args: []Argument{
					{
						Ident: ArgIdent{
							Ident: "flags",
						},
						Term: ArgType{
							Simple: TypeIdent{Ident: "#"},
						},
					},
					{
						Ident: ArgIdent{
							Ident: "b_x",
						},
						Term: ArgType{
							Simple: TypeIdent{Ident: "c_x"},
						},
					},
					{
						Ident: ArgIdent{
							Ident: "d",
						},
						Term: ArgType{
							Simple: TypeIdent{Ident: "e"},
						},
					},
					{
						Ident: ArgIdent{
							Ident: "pipka",
						},
						Flag: &Flag{
							Ident: "flags",
							Index: 2,
						},
						Term: ArgType{
							Simple: TypeIdent{Ident: "Vector"},
							Extension: &Extension{
								Inner: []TypeIdent{{Ident: "popka"}},
							},
						},
					},
				},
				Result: RetType{
					Simple: TypeIdent{Ident: "F"},
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
`,
			want: Program{
				Constraints: []ProgramEntry{
					{Newline: true},
					{Declaration: &Declaration{
						Combinator: "a#123",
						Result: RetType{
							Simple: TypeIdent{Ident: "B"},
						},
					}},
					{Declaration: &Declaration{
						Combinator: "c#456",
						Result: RetType{
							Simple: TypeIdent{Ident: "D"},
						},
					}},
				},
				Methods: []ProgramEntry{
					{Newline: true},
					{Comment: stringPtr("// another comment")},
					{Newline: true},
					{Comment: stringPtr("// comment")},
					{Declaration: &Declaration{
						Combinator: "e#789",
						Result: RetType{
							Simple: TypeIdent{Ident: "F"},
						},
					}},
					{Declaration: &Declaration{
						Combinator: "g#012",
						Result: RetType{
							Simple: TypeIdent{Ident: "H"},
						},
					}},
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
