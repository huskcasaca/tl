package typelang_test

import (
	"embed"
	"github.com/xelaj/tl"
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
		name     string
		file     string
		expected *tl.Schema
		wantErr  assert.ErrorAssertionFunc
	}{{
		file: "testdata/simplest.tl",
		expected: &tl.Schema{
			Declarations: []Declaration{
				{
					Name:      Name{Key: "someEnum"},
					CRC:       0x5508ec75,
					Category:  CategoryPredict,
					Params:    []Param{},
					OptParams: []Param{},
					Type:      tl.Type{Name: Name{Key: "CoolEnumerate"}},
				},
				{
					Name:      Name{Key: "someFunc"},
					CRC:       0x7da07ec9,
					Category:  CategoryFunction,
					Params:    []Param{},
					OptParams: []Param{},
					Type:      tl.Type{Name: Name{Key: "CoolEnumerate"}},
				},
				{
					Name:      Name{Namespace: "auth", Key: "someFunc"},
					CRC:       0x7da07ec9,
					Category:  CategoryFunction,
					Params:    []Param{},
					OptParams: []Param{},
					Type:      tl.Type{Name: Name{Key: "CoolEnumerate"}},
				},
			},
		},
	}, {
		file: "testdata/many_flags.tl",
		expected: &tl.Schema{
			Declarations: []tl.Declaration{
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
						Type:        tl.Type{Name: Name{Key: "double"}},
						FlagTrigger: "flags2",
						BitTrigger:  9,
					}, RequiredParam{
						Name: "id",
						Type: tl.Type{Name: Name{Key: "long"}},
					}},
					OptParams: []Param{},
					Type:      tl.Type{Name: Name{Key: "ChatFull"}},
				},
			},
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
