package ddl

import (
	. "github.com/spf13/cobra"
)

var (
	ksFlags     KeySpaceFlags
	KeySpaceCmd = New(
		Use("keyspace"), Alias("k"),
		PersistentFlags(
			StringP(&ksFlags.Name, "name", "n", "", ""),
			BoolP(&ksFlags.Durable, "durable", "d", "", true),
			MapStringStringP(&ksFlags.Replication, "replication", "r", "", NewSimpleReplication(1)),
		),
		Add(
			New(Use("create"), Alias("c"), RunE(func(cmd *Command, _ []string) error {
				return ddlService.CreateKeySpace(cmd.Context(), ToKeySpace(ksFlags))
			})),
			New(Use("alter"), Alias("a"), RunE(func(c *Command, _ []string) error {
				return ddlService.AlterKeySpace(c.Context(), KeySpaceKey(ksFlags.Name), ToKeySpacePatch(ksFlags))
			})),
			New(Use("drop"), Alias("d"), RunE(func(c *Command, _ []string) error {
				return ddlService.DropKeySpace(c.Context(), KeySpaceKey(ksFlags.Name))
			})),
			New(Use("list"), Alias("l"), RunE(func(c *Command, _ []string) error {
				ks, err := ddlService.ListKeySpaces(c.Context(), ToKeySpace(ksFlags))
				if err != nil {
					return err
				}
				PrintKeySpaces(ks)
				return nil
			})),
		),
	)
)
