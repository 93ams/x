package mods

import (
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
)

var NewCmd = cmd.New(
	cmd.Use("new"),
	cmd.Alias("n"),
	cmd.Run(newFiles),
)

func newFiles(*cobra.Command, []string) {

}
