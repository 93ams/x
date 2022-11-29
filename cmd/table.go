package cmd

import (
	"github.com/scylladb/gocqlx/v2"
	. "github.com/tilau2328/cql/package/cmd"
	. "github.com/tilau2328/cql/package/cmd/flags"
	. "github.com/tilau2328/cql/package/cql"
)

var (
	TableFields []string
	TableFlags  = Table{}
	TableCmd    = New(
		Use("table"),
		Alias("t"),
		PersistentFlags(
			String(&TableFlags.KeyspaceName, "keyspace_name", "", ""),
			StringP(&TableFlags.TableName, "name", "n", "", ""),
			StringSliceP(&TableFields, "fields", "f", "", ""),
		),
		Add(
			New(Use("create"), Alias("c"), RunE(wrapSession(func(s gocqlx.Session) error {
				return CreateTable(s, TableFlags.TableName, TableFields)
			}))),
			New(Use("alter"), Alias("a"), RunE(wrapSession(func(s gocqlx.Session) error {
				return AlterTable(s, TableFlags.TableName, TableFields)
			}))),
			New(Use("drop"), Alias("d"), RunE(wrapSession(func(s gocqlx.Session) error {
				return DropTable(s, TableFlags.TableName)
			}))),
			New(Use("list"), Alias("l"), RunE(wrapSession(func(s gocqlx.Session) error {
				return ListTables(s, TableFlags.KeyspaceName)
			}))),
		),
	)
)
