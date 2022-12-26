package field

import (
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
)

var RootCmd = cmd.New(
	cmd.Use("field"),
	cmd.Add(EditCmd, NewCmd, RemCmd),
	cmd.Run(listFields),
)

func listFields(*cobra.Command, []string) {

}
