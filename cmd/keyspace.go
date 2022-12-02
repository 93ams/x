package cmd

import (
	. "github.com/samber/lo"
	. "github.com/spf13/cobra"
	. "github.com/tilau2328/cql/internal/adaptor/repo/model"
	. "github.com/tilau2328/cql/package/cmd"
	. "github.com/tilau2328/cql/package/cmd/flags"
	"github.com/tilau2328/cql/package/cmd/pretty"
	"os"
)

var (
	ksFlags     KeySpace
	KeyspaceCmd = New(
		Use("keyspace"), Alias("k"),
		PersistentFlags(
			StringP(&ksFlags.Name, "name", "n", "", ""),
			BoolP(&ksFlags.Durable, "durable", "d", "", true),
			MapStringStringP(&ksFlags.Replication, "replication", "r", "", map[string]string{
				"class":              string(SimpleReplicationStrategy),
				"replication_factor": "3",
			}),
		),
		Add(
			New(Use("create"), Alias("c"), RunE(createKS)),
			New(Use("alter"), Alias("a"), RunE(alterKS)),
			New(Use("drop"), Alias("d"), RunE(dropKS)),
			New(Use("list"), Alias("l"), RunE(listKS)),
		),
	)
)

func createKS(c *Command, _ []string) error { return ksRepo.Create(c.Context(), ksFlags) }
func alterKS(c *Command, _ []string) error  { return ksRepo.Alter(c.Context(), ksFlags) }
func dropKS(c *Command, _ []string) error   { return ksRepo.Drop(c.Context(), ksFlags.Name) }
func listKS(c *Command, _ []string) error {
	ks, err := ksRepo.List(c.Context(), ksFlags)
	if err != nil {
		return err
	}
	pretty.NewTable(
		pretty.Header("#", "KeyspaceName", "DurableWrites", "Replication"),
		pretty.Rows(Map(ks, func(v KeySpace, i int) []any {
			return []any{i, v.Name, v.Durable, v.Replication}
		})...),
	).Write(os.Stdout)
	return nil
}
