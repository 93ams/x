package search

import (
	"github.com/tilau2328/x/src/go/package/cmd"
)

var RootCmd = cmd.New(
	cmd.Use("search"),
	cmd.Alias("s"),
	cmd.Flags(),
)
