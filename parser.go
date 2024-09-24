package tl

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/xelaj/tl/parser/protobuf"
	"github.com/xelaj/tl/parser/typelang"
	"github.com/xelaj/tl/parser/typelang/lexer"
	"io"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/gabriel-vasile/mimetype"
)

func IsTypeLangDefinition(raw []byte, limit uint32) (matched bool) {
	if !utf8.Valid(raw) {
		return false
	}

	str := string(raw)

	// expecting `someEnum#5508ec75 = CoolEnumerate;`
	var line string
	for len(str) > 0 {
		line, str = readLineWithoutComment(str)
		if line != "" {
			break
		}
	}

	l := lexer.New(line, lexer.LexAny())
	var wg sync.WaitGroup
	var err error
	ctx, cancel := context.WithCancel(context.TODO())
	defer func() {
		cancel()
		wg.Done()
		matched = matched && err != nil && !errors.Is(err, context.Canceled)
	}()

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		err = l.Start(ctx)
	}(ctx)

	var t *lexer.Token
	var ok bool
	if t, ok = l.NextToken(); !ok || t == nil ||
		!(t.Type == lexer.TokenLcIdent || t.Type == lexer.TokenLcIdentFull) {
		return false
	}
	if t, ok = l.NextToken(); !ok || t == nil ||
		t.Type != lexer.TokenWS {
		return false
	}
	// TODO: parse full line to make sure that this is typelang spec.

	return true
}

// returns line without comment, and without prefixed or suffixed spaces.
func readLineWithoutComment(str string) (line, left string) {
	i := strings.IndexRune(str, '\n')
	if i >= 0 {
		line, left = str[:i], str[i:]
	} else {
		line, left = str, ""
	}

	commentStart := strings.Index(line, "//")
	return strings.TrimSpace(line[:commentStart]), left
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
		mimetype.Extend(IsTypeLangDefinition, mimeTypeLang, ".tl")
	})

	return mimetype.Detect(b).String()
}

func Parse(filename, predictedMime string, r io.Reader) (*Schema, error) {
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
		return protobuf.Parse(filename, r)

	default:
		return nil, fmt.Errorf("%#v is not supported", predictedMime)
	}
}
