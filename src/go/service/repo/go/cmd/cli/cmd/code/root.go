package code

import (
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/cmd/code/deps"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/cmd/code/funcs"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/cmd/code/types"
)

var RootCmd = cmd.New(
	cmd.Use("code"), cmd.Alias("c"),
	cmd.Add(deps.RootCmd, funcs.RootCmd, types.RootCmd),
)
