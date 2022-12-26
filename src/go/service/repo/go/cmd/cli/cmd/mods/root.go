package mods

import (
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/cmd/mods/deps"
)

var RootCmd = cmd.New(
	cmd.Use("mods"), cmd.Alias("m"),
	cmd.Add(FetchCmd, NewCmd, RemCmd, deps.RootCmd),
	cmd.Run(listFiles),
)

func listFiles(*cobra.Command, []string) {

}
