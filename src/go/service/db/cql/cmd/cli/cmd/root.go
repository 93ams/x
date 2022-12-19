package cmd

import (
	"github.com/gocql/gocql"
	"github.com/samber/lo"
)

var (
	// Flags
	consistency uint16
	ks          string
	hosts       []string
	// RootCmd Command
	RootCmd = New(
		Use("cql"),
		PersistentFlags(
			Uint16P(&consistency, "consistency", "c", "", uint16(gocql.One)),
			StringSlice(&hosts, "hosts", "", "localhost:9042"),
			String(&ks, "ks", "", ""),
		),
		Add(ddl.KeySpaceCmd, ddl.TableCmd, dml.RootCmd),
	)
)

// Execute runs Command
func Execute() error {
	//return GenDocument(RootCmd, "./cmd")
	session, fn := lo.Must2(cql.NewSession(cql.NewCluster(cql.options{
		Consistency: gocql.Consistency(consistency),
		KeySpace:    cql.KeySpace(ks),
		Hosts:       hosts,
	})))
	defer fn()
	return RootCmd.Execute()
}
