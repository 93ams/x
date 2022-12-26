package field

import (
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
)

var EditCmd = cmd.New(
	cmd.Use("edit"),
	cmd.Alias("e"),
	cmd.Run(editFunc),
)

func editFunc(*cobra.Command, []string) {

}
