package cli

import (
	"github.com/gocql/gocql"
	. "github.com/spf13/cobra"
	"github.com/tilau2328/cql/cmd/cli/crud"
	. "github.com/tilau2328/cql/package/adaptor/data/cql/repo/ddl"
	. "github.com/tilau2328/cql/package/shared/cmd"
	. "github.com/tilau2328/cql/package/shared/cmd/flags"
	. "github.com/tilau2328/cql/package/shared/cql"
)

var (
	// Flags
	consistency uint16
	ks          string
	hosts       []string
	// Services
	ksRepo *KeySpaceRepo
	tRepo  *TableRepo
	// RootCmd Command
	RootCmd = New(
		Use("cql"),
		PersistentPreRunE(boot),
		PersistentFlags(
			Uint16P(&consistency, "consistency", "c", "", uint16(gocql.One)),
			StringSlice(&hosts, "hosts", "", "localhost:9042"),
			String(&ks, "ks", "", ""),
		),
		Add(KeyspaceCmd, TableCmd, crud.RootCmd),
	)
)

// Execute runs Command
func Execute() error {
	return GenDocument(RootCmd, "./cmd")
	//return RootCmd.Execute()
}

func boot(*Command, []string) error {
	session, err, _ := NewSession(NewCluster(gocql.Consistency(consistency), ks, hosts...))
	if err != nil {
		return err
	}
	ksRepo = NewKeySpaceRepo(session)
	tRepo = NewTableRepo(session)
	return nil
}
