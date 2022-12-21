package create

import (
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/cmd/create/pattern"
)

var RootCmd = cmd.New(
	cmd.Use("create"),
	cmd.Add(MethodCmd, StructCmd, InterfaceCmd, pattern.RootCmd),
)
