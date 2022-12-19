package transform

import "github.com/tilau2328/x/src/go/package/cmd"

var RootCmd = cmd.New(
	cmd.Use("transform"),
	cmd.Alias("t"),
	cmd.Flags(),
)
