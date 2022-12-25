package create

import (
	"github.com/tilau2328/x/src/go/package/cmd"
)

var RootCmd = cmd.New(
	cmd.Use("create"),
	cmd.Add(
		MapperCmd,
	),
)
