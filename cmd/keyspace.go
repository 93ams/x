package cmd

import (
	"github.com/scylladb/gocqlx/v2"
	. "github.com/tilau2328/cql/package/cmd"
	. "github.com/tilau2328/cql/package/cmd/flags"
	. "github.com/tilau2328/cql/package/cql"
)

var (
	KeyspaceFlags Keyspace
	KeyspaceCmd   = New(
		Use("keyspace"),
		Alias("k"),
		PersistentFlags(
			StringP(&KeyspaceFlags.KeyspaceName, "name", "n", "", ""),
			BoolP(&KeyspaceFlags.DurableWrites, "durable", "d", "", true),
			MapStringStringP(&KeyspaceFlags.Replication, "replication", "r", "", map[string]string{
				"class":              string(SimpleReplication),
				"replication_factor": "3",
			}),
		),
		Add(
			New(Use("create"), Alias("c"), RunE(wrapSession(func(s gocqlx.Session) error {
				return CreateKeySpace(s, KeyspaceFlags)
			}))),
			New(Use("alter"), Alias("a"), RunE(wrapSession(func(s gocqlx.Session) error {
				return AlterKeySpace(s, KeyspaceFlags)
			}))),
			New(Use("drop"), Alias("d"), RunE(wrapSession(func(s gocqlx.Session) error {
				return DropKeySpace(s, KeyspaceFlags)
			}))),
			New(Use("list"), Alias("l"), RunE(wrapSession(func(s gocqlx.Session) error {
				return ListKeySpaces(s, KeyspaceFlags)
			}))),
		),
	)
)
