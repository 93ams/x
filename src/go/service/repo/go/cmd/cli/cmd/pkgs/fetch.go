package pkgs

import (
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
)

var FetchCmd = cmd.New(
	cmd.Use("fetch"),
	cmd.Alias("f"),
	cmd.Run(fetchFiles),
)

func fetchFiles(*cobra.Command, []string) {

}
