package proto

import (
	"github.com/xelaj/tl"
	"io"
	"io/fs"
	"os"
	"strings"
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
		return nil, err //nolint:wrapcheck // it's important to keep error
	}

	normalized, err := normalizeProto(res)
	if err != nil {
		return nil, err
	}

	return normalized, nil
}

func normalizeProto(p *Proto) (*tl.Schema, error) {
	panic("Unimplemented")
}
