package cmd

import (
	"github.com/gocql/gocql"
	. "github.com/spf13/cobra"
	"github.com/tilau2328/cql/cmd/crud"
	. "github.com/tilau2328/cql/internal/adaptor/repo"
	. "github.com/tilau2328/cql/package/cmd"
	. "github.com/tilau2328/cql/package/cmd/flags"
	. "github.com/tilau2328/cql/package/cql"
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
	session, err := NewSession(NewCluster(gocql.Consistency(consistency), ks, hosts...))
	if err != nil {
		return err
	}
	ksRepo = NewRepository(session)
	tRepo = NewTableRepo(session)
	return nil
}
