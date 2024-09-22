package tl_test

import (
	"embed"
	"github.com/xelaj/tl"
	. "github.com/xelaj/tl/parser/tl"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/xelaj/tl"
)

//go:embed testdata
var testdata embed.FS

func TestParseFile(t *testing.T) {
	for _, tt := range []struct {
		name     string
		file     string
		expected *tl.TLSchema
		wantErr  assert.ErrorAssertionFunc
	}{{
		file: "internal/testdata/simplest.tl",
		expected: &tl.TLSchema{
			TypeSeq: []TLName{{Key: "CoolEnumerate"}},
			TypeDeclMap: map[TLName]tl.TLTypeDeclaration{
				{Key: "CoolEnumerate"}: {
					Declarations: []TLDeclaration{{
						Name:       TLName{Key: "someEnum"},
						CRC:        0x5508ec75,
						Params:     []TLParam{},
						PolyParams: []TLParam{},
						Type:       tl.TLType{TLName: TLName{Key: "CoolEnumerate"}},
					}},
				},
			},
			FuncSeq: []TLName{{Key: "someFunc"}, {Namespace: "auth", Key: "someFunc"}},
			FuncDeclMap: map[TLName]TLDeclaration{
				TLName{Key: "someFunc"}: {
					Name:       TLName{Key: "someFunc"},
					CRC:        0x7da07ec9,
					Params:     []TLParam{},
					PolyParams: []TLParam{},
					Type:       tl.TLType{TLName: TLName{Key: "CoolEnumerate"}},
				},
				TLName{Namespace: "auth", Key: "someFunc"}: {
					Name:       TLName{Namespace: "auth", Key: "someFunc"},
					CRC:        0x7da07ec9,
					Params:     []TLParam{},
					PolyParams: []TLParam{},
					Type:       tl.TLType{TLName: TLName{Key: "CoolEnumerate"}},
				},
			},
		},
	}, {
		file: "internal/testdata/many_flags.tl",
		expected: &tl.TLSchema{
			TypeSeq: []TLName{{Key: "ChatFull"}},
			TypeDeclMap: map[TLName]tl.TLTypeDeclaration{
				{Key: "ChatFull"}: {
					Declarations: []TLDeclaration{{
						Name: TLName{Key: "a"},
						CRC:  0xf2355507,
						Params: []TLParam{TLBitflagParam{
							Name: "flags",
						}, TLTriggerParam{
							Name:        "opt_prop",
							FlagTrigger: "flags",
							BitTrigger:  3,
						}, TLBitflagParam{
							Name: "flags2",
						}, TLOptionalParam{
							Name:        "opt2_prop",
							Type:        tl.TLType{TLName: TLName{Key: "double"}},
							FlagTrigger: "flags2",
							BitTrigger:  9,
						}, TLRequiredParam{
							Name: "id",
							Type: tl.TLType{TLName: TLName{Key: "long"}},
						}},
						PolyParams: []TLParam{},
						Type:       tl.TLType{TLName: TLName{Key: "ChatFull"}},
					}},
				},
			},
			FuncSeq:     []TLName{},
			FuncDeclMap: map[TLName]TLDeclaration{},
		},
	}} {
		tt := tt // for parallel tests
		tt.wantErr = noErrAsDefault(tt.wantErr)

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data, err := testdata.ReadFile(tt.file)
			require.NoError(t, err)

			got, err := ParseString(tt.file, string(data))
			if !tt.wantErr(t, err) || err != nil {
				return
			}

			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestEquality(t *testing.T) {
	for _, tt := range []struct {
		name string
		file string
	}{{
		name: "simplest",
		file: "testdata/simplest.tl",
	}, {
		name: "many_flags",
		file: "testdata/many_flags.tl",
	}, {
		name: "with_comments",
		file: "testdata/with_comments.tl",
	}} {
		tt := tt // for parallel tests

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data, err := testdata.ReadFile(tt.file)
			require.NoError(t, err)

			got, err := ParseString(tt.file, string(data))
			require.NoError(t, err)

			assert.Equal(t, string(data), got.String())
		})
	}
}

func noErrAsDefault(e assert.ErrorAssertionFunc) assert.ErrorAssertionFunc {
	if e == nil {
		return assert.NoError
	}

	return e
}
