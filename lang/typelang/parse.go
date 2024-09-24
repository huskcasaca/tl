package typelang

import (
	"github.com/xelaj/tl"
	"github.com/xelaj/tl/lang/typelang/lexer"
	"github.com/xelaj/tl/lang/typelang/program"
	"io"
	"io/fs"
	"os"
	"strings"

	"github.com/alecthomas/participle/v2"
)

var parser = participle.MustBuild[program.Program](
	participle.Lexer(lexer.NewDefinition()),
)

func ParseFile(filename string) (*tl.Schema, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Parse(filename, f)
}

func ParseFS(fsys fs.FS, filename string) (*tl.Schema, error) {
	f, err := fsys.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Parse(filename, f)
}

func ParseString(filename string, content string) (*tl.Schema, error) {
	return Parse(filename, strings.NewReader(content))
}

func Parse(filename string, content io.Reader) (*tl.Schema, error) {
	res, err := parser.Parse(filename, content)
	if err != nil {
		return nil, err
	}

	normalized, err := res.Normalize()

	if err != nil {
		return nil, err
	}
	normalized.Layer = 0

	return normalized, nil
}
