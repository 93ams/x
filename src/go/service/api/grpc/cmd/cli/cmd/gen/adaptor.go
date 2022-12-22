package gen

import (
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
)

var AdaptorCmd = cmd.New(
	cmd.Use("adaptor"),
	cmd.Alias("a"),
	cmd.Run(runAdaptor),
)

func runAdaptor(cmd *cobra.Command, _ []string) {
	s := service.FromCtx(cmd.Context())

}
