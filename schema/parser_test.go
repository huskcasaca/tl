package schema_test

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/xelaj/tl/schema"
)

//go:embed internal/testdata
var testdata embed.FS

func TestParseFile(t *testing.T) {
	for _, tt := range []struct {
		name     string
		file     string
		expected *TLSchema
		wantErr  assert.ErrorAssertionFunc
	}{{
		file: "internal/testdata/simplest.tl",
		expected: &TLSchema{
			ObjSeq:     []TLName{{Key: "CoolEnumerate"}},
			TypeObjMap: map[TLName]TypeTLObjects{},
			EnumObjMap: map[TLName]EnumTLObjects{
				{Key: "CoolEnumerate"}: {
					Objects: []TLObject{{
						Name:       TLName{Key: "someEnum"},
						CRC:        0x5508ec75,
						Params:     []TLParam{},
						PolyParams: []TLParam{},
						Type:       TLTypeCommon{TLName: TLName{Key: "CoolEnumerate"}},
					}},
				},
			},
			FunctionSeq: []string{"", "auth"},
			FunctionMap: map[string][]TLObject{
				"": {{
					Name:       TLName{Key: "someFunc"},
					CRC:        0x7da07ec9,
					Params:     []TLParam{},
					PolyParams: []TLParam{},
					Type:       TLTypeCommon{TLName: TLName{Key: "CoolEnumerate"}},
				}},
				"auth": {{
					Name:       TLName{Namespace: "auth", Key: "someFunc"},
					CRC:        0x7da07ec9,
					Params:     []TLParam{},
					PolyParams: []TLParam{},
					Type:       TLTypeCommon{TLName: TLName{Key: "CoolEnumerate"}},
				}},
			},
		},
	}, {
		file: "internal/testdata/many_flags.tl",
		expected: &TLSchema{
			ObjSeq: []TLName{{Key: "ChatFull"}},
			TypeObjMap: map[TLName]TypeTLObjects{
				{Key: "ChatFull"}: {
					Objects: []TLObject{{
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
							Type:        TLTypeCommon{TLName: TLName{Key: "double"}},
							FlagTrigger: "flags2",
							BitTrigger:  9,
						}, TLRequiredParam{
							Name: "id",
							Type: TLTypeCommon{TLName: TLName{Key: "long"}},
						}},
						PolyParams: []TLParam{},
						Type:       TLTypeCommon{TLName: TLName{Key: "ChatFull"}},
					}},
				},
			},
			EnumObjMap:  map[TLName]EnumTLObjects{},
			FunctionSeq: []string{},
			FunctionMap: map[string][]TLObject{},
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
		file: "internal/testdata/simplest.tl",
	}, {
		name: "many_flags",
		file: "internal/testdata/many_flags.tl",
	}, {
		name: "with_comments",
		file: "internal/testdata/with_comments.tl",
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
