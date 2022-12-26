package structs

import (
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/cmd/code/types/structs/fields"
)

var RootCmd = cmd.New(
	cmd.Use("struct"), cmd.Alias("s"),
	cmd.Add(FetchCmd, NewCmd, fields.RootCmd),
	cmd.Run(listStructs),
)

func listStructs(cmd *cobra.Command, args []string) {

}
