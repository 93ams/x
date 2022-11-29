package cmd

import (
	"github.com/gocql/gocql"
	"github.com/samber/lo"
	"github.com/scylladb/gocqlx/v2"
	"github.com/spf13/cobra"
	. "github.com/tilau2328/cql/package/cmd"
	. "github.com/tilau2328/cql/package/cmd/flags"
	"github.com/tilau2328/cql/package/cql"
)

var (
	// Flags
	consistency uint16
	keyspace    string
	hosts       []string
	// Common
	session gocqlx.Session
	// RootCmd Command
	RootCmd = New(
		Use("cql"),
		PersistentPreRun(boot),
		PersistentFlags(
			Uint16P(&consistency, "consistency", "c", "", uint16(gocql.One)),
			StringP(&keyspace, "keyspace", "k", "", "system_schema"),
			StringSlice(&hosts, "hosts", "", "localhost:9042"),
		),
		Add(KeyspaceCmd, TableCmd),
	)
)

// Execute runs Command
func Execute() error {
	//return GenDocument(RootCmd, "./cmd")
	return RootCmd.Execute()
}

func boot(*cobra.Command, []string) {
	session = lo.Must(cql.NewSession(cql.NewCluster(gocql.Consistency(consistency), keyspace, hosts...)))
}
func wrapSession(fn func(gocqlx.Session) error) func(*cobra.Command, []string) error {
	return func(*cobra.Command, []string) error { return fn(session) }
}
