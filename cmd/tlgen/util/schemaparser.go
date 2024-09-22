package util

import (
	"bufio"
	"fmt"
	"github.com/xelaj/tl"
	typelang "github.com/xelaj/tl/parser/tl"
	"io"
	"sync"

	"github.com/gabriel-vasile/mimetype"

	"github.com/xelaj/tl/parser/proto"
)

func DetectAndParseSchema(filename, predictedMime string, r io.Reader) (*tl.Schema, error) {
	if predictedMime == "" {
		buf := bufio.NewReader(r)
		r = buf

		b, err := buf.Peek(mimeBufferMax)
		if err != nil {
			return nil, fmt.Errorf("predicting data format: %w", err)
		}
		predictedMime = predictMime(b)
	}

	switch predictedMime {
	case mimeTypeLang:
		return typelang.Parse(filename, r)

	case mimeProtobuf:
		return proto.Parse(filename, r)

	default:
		return nil, fmt.Errorf("%#v is not supported", predictedMime)
	}
}

// for both of these file formats, there is no official mime type, so using
// most common mime variants.
const (
	mimeTypeLang    = "application/x-typelang"
	mimeTypeLangBin = "application/x-tlb"
	mimeProtobuf    = "application/x-protofile"

	mimeBufferMax = 512
)

var onceExtend sync.Once

func predictMime(b []byte) string {
	onceExtend.Do(func() {
		mimetype.SetLimit(mimeBufferMax)
		mimetype.Extend(tl.IsTypeLangDefinition, mimeTypeLang, ".tl")
	})

	return mimetype.Detect(b).String()
}
