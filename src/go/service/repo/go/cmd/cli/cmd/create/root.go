package create

import "github.com/tilau2328/x/src/go/package/cmd"

var RootCmd = cmd.New(
	cmd.Use("new"),
	cmd.Add(MethodCmd, StructCmd, InterfaceCmd),
)
