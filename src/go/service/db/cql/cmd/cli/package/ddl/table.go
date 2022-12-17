package ddl

import (
	. "github.com/samber/lo"
	"os"
)

func PrintKeySpaces(ks []KeySpace) {
	NewTable(
		Header("#", "KeyspaceName", "DurableWrites", "Replication"),
		Rows(Map(ks, func(v KeySpace, i int) []any {
			return []any{i, v.KeySpaceKey, v.Durable, v.Replication}
		})...),
	).Write(os.Stdout)
}

func PrintTables(t []Table) {
	NewTable(
		Header("#", "Id", "KeyspaceName", "TableName", "comment"),
		Rows(Map(t, func(v Table, i int) []any {
			return []any{i, v.Id, v.KeySpace, v.Name, v.Comment}
		})...),
	).Write(os.Stdout)
}
