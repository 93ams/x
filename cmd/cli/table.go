package cli

import (
	. "github.com/samber/lo"
	. "github.com/spf13/cobra"
	. "github.com/tilau2328/cql/package/adaptor/data/cql/model"
	. "github.com/tilau2328/cql/package/cql"
	. "github.com/tilau2328/cql/package/shared/cmd"
	. "github.com/tilau2328/cql/package/shared/cmd/flags"
	"github.com/tilau2328/cql/package/shared/cmd/pretty"
	"os"
)

var (
	tFields  []string
	tFlags   = Table{}
	TableCmd = New(
		Use("table"), Alias("t"),
		PersistentFlags(
			StringP(&tFlags.Keyspace, "keyspace_name", "k", "", ""),
			StringSliceP(&tFields, "fields", "f", "", ""),
			StringP(&tFlags.Name, "name", "n", "", ""),
		),
		Add(
			New(Use("create"), Alias("c"), RunE(createT)),
			New(Use("alter"), Alias("a"), RunE(alterT)),
			New(Use("drop"), Alias("d"), RunE(dropT)),
			New(Use("list"), Alias("l"), RunE(listT)),
		),
	)
)

func createT(c *Command, _ []string) error { return tRepo.Create(c.Context(), tFlags.Name, tFields) }
func alterT(c *Command, _ []string) error  { return tRepo.Alter(c.Context(), tFlags.Name, tFields) }
func dropT(c *Command, _ []string) error   { return tRepo.Drop(c.Context(), tFlags.Name) }
func listT(c *Command, _ []string) error {
	ret, err := tRepo.List(c.Context(), tFlags)
	if err != nil {
		return err
	}
	pretty.NewTable(
		pretty.Header("#", "Id", "KeyspaceName", "TableName", "Comment"),
		pretty.Rows(Map(ret, func(v Table, i int) []any {
			return []any{i, v.Id, v.Keyspace, v.Name, v.Comment}
		})...),
	).Write(os.Stdout)
	return nil
}
