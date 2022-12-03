package cmd

import (
	"github.com/gocql/gocql"
	ddl3 "github.com/tilau2328/cql/cmd/cli/cmd/ddl"
	"github.com/tilau2328/cql/cmd/cli/cmd/dml"
	. "github.com/tilau2328/cql/package/adaptor/data/cql/repo/ddl"
	"github.com/tilau2328/cql/package/domain/provider"
	. "github.com/tilau2328/cql/package/shared/cmd"
	. "github.com/tilau2328/cql/package/shared/cmd/flags"
)

var (
	// Flags
	consistency uint16
	ks          string
	hosts       []string
	// Services
	ddl    provider.DDL
	ksRepo *KeySpaceRepo
	tRepo  *TableRepo
	// RootCmd Command
	RootCmd = New(
		Use("cql"),
		PersistentFlags(
			Uint16P(&consistency, "consistency", "c", "", uint16(gocql.One)),
			StringSlice(&hosts, "hosts", "", "localhost:9042"),
			String(&ks, "ks", "", ""),
		),
		Add(ddl3.KeyspaceCmd, ddl3.TableCmd, dml.RootCmd),
	)
)

// Execute runs Command
func Execute() error {
	return GenDocument(RootCmd, "./cmd")
	//return RootCmd.Execute()
}
