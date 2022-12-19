package cmd

import (
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/services/repo/go/cmd/cli/cmd/create"
	"github.com/tilau2328/x/src/go/services/repo/go/cmd/cli/cmd/search"
	"github.com/tilau2328/x/src/go/services/repo/go/cmd/cli/cmd/transform"
)

var RootCmd = cmd.New(
	cmd.Use("xgo"),
	cmd.Add(search.RootCmd, create.RootCmd, transform.RootCmd),
)

func Execute() error { return RootCmd.Execute() }
