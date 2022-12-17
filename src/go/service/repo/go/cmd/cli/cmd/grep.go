package cmd

import (
	"github.com/spf13/cobra"
)

var GrepCmd = New(
	Use("grep"),
	Alias("g"),
	Flags(),
	Run(runGrep),
)

func runGrep(*cobra.Command, []string) {

}
