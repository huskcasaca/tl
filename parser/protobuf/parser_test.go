package protobuf_test

import (
	"embed"
	"github.com/xelaj/tl"
	. "github.com/xelaj/tl/parser/protobuf"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testdata
var testdata embed.FS

func TestParseFile(t *testing.T) {
	for _, tt := range []struct {
		name    string
		file    string
		want    *tl.Schema
		wantErr assert.ErrorAssertionFunc
	}{
		// {
		// 	file:     "testdata/simplest.tl",
		// 	expected: &schema.Schema{},
		// },
	} {
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
