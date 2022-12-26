package interfaces

import (
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/cmd/code/types/interfaces/field"
)

var RootCmd = cmd.New(
	cmd.Use("interface"), cmd.Alias("i"),
	cmd.Add(FetchCmd, NewCmd, field.RootCmd),
	cmd.Run(listInterfaces),
)

func listInterfaces(cmd *cobra.Command, args []string) {

}
