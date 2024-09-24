package typelang_test

import (
	"embed"
	. "github.com/xelaj/tl/parser/typelang"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/xelaj/tl"
)

//go:embed testdata
var testdata embed.FS

func TestParseFile(t *testing.T) {
	for _, tt := range []struct {
		name    string
		input   string
		want    *Schema
		wantErr assert.ErrorAssertionFunc
	}{{
		input: "testdata/simplest.tl",
		want: &Schema{
			Definition: []Definition{
				{
					Name:      Name{Key: "someEnum"},
					CRC:       0x5508ec75,
					Category:  CategoryPredict,
					Params:    []Param{},
					OptParams: []Param{},
					Type:      Type{Name: Name{Key: "CoolEnumerate"}},
				},
				{
					Name:      Name{Key: "someFunc"},
					CRC:       0x7da07ec9,
					Category:  CategoryFunction,
					Params:    []Param{},
					OptParams: []Param{},
					Type:      Type{Name: Name{Key: "CoolEnumerate"}},
				},
				{
					Name:      Name{Namespace: "auth", Key: "someFunc"},
					CRC:       0x7da07ec9,
					Category:  CategoryFunction,
					Params:    []Param{},
					OptParams: []Param{},
					Type:      Type{Name: Name{Key: "CoolEnumerate"}},
				},
			},
		},
	}, {
		input: "testdata/flags.tl",
		want: &Schema{
			Definition: []Definition{
				{
					Name:     Name{Key: "a"},
					CRC:      0xf2355507,
					Category: CategoryPredict,
					Params: []Param{BitflagParam{
						Name: "flags",
					}, TriggerParam{
						Name:        "opt_prop",
						FlagTrigger: "flags",
						BitTrigger:  3,
					}, BitflagParam{
						Name: "flags2",
					}, OptionalParam{
						Name:        "opt2_prop",
						Type:        Type{Name: Name{Key: "double"}},
						FlagTrigger: "flags2",
						BitTrigger:  9,
					}, RequiredParam{
						Name: "id",
						Type: Type{Name: Name{Key: "long"}},
					}},
					OptParams: []Param{},
					Type:      Type{Name: Name{Key: "ChatFull"}},
				},
			},
		},
	}, {
		input: "testdata/annotations.tl",
		want: &Schema{
			Definition: []Definition{
				{
					Name:       Name{Key: "predict"},
					CRC:        0x12345678,
					Category:   CategoryPredict,
					Params:     []Param{},
					OptParams:  []Param{},
					Type:       Type{Name: Name{Key: "SomeType"}},
					Comment:    "this is a comment for predict",
					RetComment: "this is a comment for return type",
				},
				{
					Name:     Name{Key: "function"},
					CRC:      305419896,
					Category: CategoryFunction,
					Params: []Param{
						RequiredParam{
							Comment: "this is a comment for param1",
							Name:    "param1",
							Type: Type{
								Name: Name{Key: "string"},
							},
						},
						RequiredParam{
							Comment: "this is a comment for param2",
							Name:    "param2",
							Type: Type{
								Name: Name{Key: "bool"},
							},
						},
					},
					OptParams: []Param{},
					Type: Type{
						Name: Name{Key: "SomeType"},
					},
					Comment:    "this is a comment for function",
					RetComment: "this is a comment for return type",
				},
				{
					Name:      Name{Key: "functionNoFields"},
					CRC:       0x12345678,
					Category:  CategoryFunction,
					Params:    []Param{},
					OptParams: []Param{},
					Type: Type{
						Name: Name{
							Key: "SomeType",
						},
					},
				},
				{
					Name:     Name{Key: "functionWithoutComments"},
					CRC:      0x12345678,
					Category: CategoryFunction,
					Params: []Param{
						BitflagParam{
							Name: "flags",
						},
						TriggerParam{
							Name:        "a",
							FlagTrigger: "flags",
							BitTrigger:  1,
						},
					},
					OptParams: []Param{},
					Type: Type{
						Name: Name{Key: "SomeType"},
					},
				},
			},
		},
	}} {
		tt := tt // for parallel tests
		tt.wantErr = noErrAsDefault(tt.wantErr)

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data, err := testdata.ReadFile(tt.input)
			require.NoError(t, err)

			got, err := ParseString(tt.input, string(data))
			if !tt.wantErr(t, err) || err != nil {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func noErrAsDefault(e assert.ErrorAssertionFunc) assert.ErrorAssertionFunc {
	if e == nil {
		return assert.NoError
	}

	return e
}
