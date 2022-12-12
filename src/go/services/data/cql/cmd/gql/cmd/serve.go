package cmd

import (
	"github.com/gocql/gocql"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	cql2 "github.com/tilau2328/cql/src/go/package/shared/data/cql"
)

var ServeCmd = New(
	Use("serve"), Alias("s"),
	Run(func(cmd *cobra.Command, _ []string) {
		server, finish := lo.Must2(Init(cql2.Options{Consistency: gocql.One}, "/graphql"))
		defer finish()
		server.WithPlayground("/").Serve()
	}),
)
