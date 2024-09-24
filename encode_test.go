//revive:disable:add-constant    // It's a test file, we can't avoid magic constants
//revive:disable:function-length // who cares for test files

package tl_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "github.com/xelaj/tl"
)

func BenchmarkEncoder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkFunc(Marshal(&AccountInstallThemeParams{
			Dark:   true,
			Format: ptr("abc"),
			Theme: &InputThemeObj{
				ID:         123,
				AccessHash: 321,
			},
		}))
	}
}

func TestEncode(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name    string
		input   any
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}{{
		name: "Rights",
		input: &Rights{
			DeleteMessages: true,
			BanUsers:       true,
		},
		want:    Hexed("D524B25F18000000"),
		wantErr: assert.NoError,
	}, {
		name: "AccountInstallThemeParams",
		input: &AccountInstallThemeParams{
			Dark:   true,
			Format: ptr("abc"),
			Theme: &InputThemeObj{
				ID:         123,
				AccessHash: 321,
			},
		},
		want:    Hexed("3737E47A0300000003616263E993563C7B000000000000004101000000000000"),
		wantErr: assert.NoError,
	}, {
		name: "AccountUnregisterDeviceParams",
		input: &AccountUnregisterDeviceParams{
			TokenType: 1,
			Token:     "foo",
			OtherUids: []int32{
				1337, 228, 322,
			},
		},
		want:    Hexed("BFC476300100000003666F6F15C4B51C0300000039050000E400000042010000"),
		wantErr: assert.NoError,
	}, {
		name: "respq",
		input: &ResPQ{
			Nonce:        NewInt128(123),
			ServerNonce:  NewInt128(321),
			Pq:           []byte{1, 2, 3},
			Fingerprints: []int64{322, 1337},
		},
		want: Hexed("632416050000000000000000000000000000007B00000000000000000000" +
			"0000000001410301020315C4B51C0200000042010000000000003905000000000000"),
		wantErr: assert.NoError,
	}, {
		name: "respq_raw",
		input: &ResPQRaw{
			Nonce:        [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7b},
			ServerNonce:  [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x41},
			Pq:           []byte{1, 2, 3},
			Fingerprints: []int64{322, 1337},
		},
		want: Hexed("642416050000000000000000000000000000007B00000000000000000000" +
			"0000000001410301020315C4B51C0200000042010000000000003905000000000000"),
		wantErr: assert.NoError,
	}, {
		name: "InitConnectionParams",
		input: &InvokeWithLayerParams{
			Layer: 322,
			Query: &InitConnectionParams{
				APIID:          1337,
				DeviceModel:    "abc",
				SystemVersion:  "def",
				AppVersion:     "123",
				SystemLangCode: "en",
				LangCode:       "en",
				Query:          &SomeNullStruct{},
			},
		},
		want: Hexed(`
			0d0d9bda         // crc
			42010000         // Layer
			a95ecdc1         // crc InitConnectionParams
			        00000000 // flag
			        39050000 // api id
			        03       // length of string
			          616263 // data
			        03       // length of string
			          646566 // data
		            03       // length of string
			          313233 // data
			        02       // length of string
			          656e   // data
			              00 // padding
					00000000 // empty required string
			        02       // length of string
			          656e   // data
			              00 // padding
			        6b18f9c4 // null struct crc`),
		wantErr: assert.NoError,
	}} {
		tt := tt // for parallel tests

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Marshal(tt.input)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
