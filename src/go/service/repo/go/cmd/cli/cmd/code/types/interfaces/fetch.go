package interfaces

import (
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
)

var FetchCmd = cmd.New(
	cmd.Use("fetch"),
	cmd.Alias("f"),
	cmd.Run(fetchFunc),
)

func fetchFunc(*cobra.Command, []string) {

}
