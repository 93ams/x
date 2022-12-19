package cmd

import (
	"github.com/spf13/cobra"
	. "github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/services/repo/go/package/wrapper/coding/resolver"
	"github.com/tilau2328/x/src/go/services/repo/go/package/wrapper/coding/restorer"
)

var NewCmd = New(
	Use("new"),
	Alias("n"),
	Flags(),
	Run(runNew),
)

func runNew(*cobra.Command, []string) {
	r := restorer.NewRestorerWithImports("root", resolver.NewGuessResolver())
	r.Print()
}
