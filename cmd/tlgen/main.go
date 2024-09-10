package main

import (
	"os"

	"github.com/xelaj/tl/cmd/tlgen/root"
	"github.com/xelaj/tl/cmd/tlgen/util"
)

func main() {
	ctx, cancel := util.NewContext(os.Stdin, os.Stdout, os.Stderr)
	defer cancel()

	root.Cmd().ExecuteContext(ctx)
}
