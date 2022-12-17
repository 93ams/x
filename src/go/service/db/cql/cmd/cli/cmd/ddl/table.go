package ddl

import (
	. "github.com/spf13/cobra"
)

var (
	tableFlags = TableFlags{}
	TableCmd   = New(
		Use("table"), Alias("t"),
		PersistentFlags(
			StringP(&tableFlags.Keyspace, "keyspace_name", "k", "", ""),
			StringSliceP(&tableFlags.Fields, "fields", "f", "", ""),
			StringP(&tableFlags.Name, "name", "n", "", ""),
		),
		Add(
			New(Use("create"), Alias("c"), RunE(func(c *Command, _ []string) error {
				return ddlService.CreateTable(c.Context(), ToTable(tableFlags))
			})),
			New(Use("alter"), Alias("a"), RunE(func(c *Command, _ []string) error {
				return ddlService.AlterTable(c.Context(), ToTableKey(tableFlags), ToTablePatch(tableFlags))
			})),
			New(Use("drop"), Alias("d"), RunE(func(c *Command, _ []string) error {
				return ddlService.DropTable(c.Context(), ToTableKey(tableFlags))
			})),
			New(Use("list"), Alias("l"), RunE(func(c *Command, _ []string) error {
				t, err := ddlService.ListTables(c.Context(), ToTable(tableFlags))
				if err != nil {
					return err
				}
				PrintTables(t)
				return nil
			})),
		),
	)
)
