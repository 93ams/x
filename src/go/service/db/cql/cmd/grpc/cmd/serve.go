package cmd

import (
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	cql2 "github.com/tilau2328/x/src/go/package/shared/data/cql"
)

var ServeCmd = New(
	Use("serve"), Alias("s"),
	Run(func(cmd *cobra.Command, _ []string) {
		lo.Must0(Init(cql2.Options{}).Serve())
	}),
)
