package funcs

import (
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
)

var RemCmd = cmd.New(
	cmd.Use("rem"),
	cmd.Alias("r"),
	cmd.Run(remFunc),
)

func remFunc(*cobra.Command, []string) {

}
