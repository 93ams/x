package deps

import (
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
)

var RootCmd = cmd.New(
	cmd.Use("deps"),
	cmd.Alias("d"),
	cmd.Run(listDeps),
)

func listDeps(*cobra.Command, []string) {

}
