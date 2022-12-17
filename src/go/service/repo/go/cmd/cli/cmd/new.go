package cmd

import (
	"github.com/spf13/cobra"
	. "github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/services/gen/go/package/adaptor/driver/resolver"
	"github.com/tilau2328/x/src/go/services/gen/go/package/adaptor/driver/restorer"
)

var NewCmd = New(
	Use("new"),
	Alias("n"),
	Flags(),
	Run(runNew),
)

func runNew(*cobra.Command, []string) {
	r := restorer.NewRestorerWithImports("root", resolver.NewGuessResolver())
}
