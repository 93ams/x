package cql

import (
	"github.com/samber/lo"
	"github.com/scylladb/gocqlx/v2"
	"github.com/tilau2328/cql/internal/models"
	"github.com/tilau2328/cql/package/cmd/pretty"
	"os"
)

type (
	ReplicationStrategy string
	Keyspace            struct {
		KeyspaceName  string
		DurableWrites bool
		Replication   map[string]string
	}
)

const (
	SimpleReplication  ReplicationStrategy = "SimpleStrategy"
	NetworkReplication ReplicationStrategy = "NetworkTopologyStrategy"
)

func ListKeySpaces(session gocqlx.Session, in Keyspace) error {
	var ret []Keyspace
	if err := session.Query(models.Keyspaces.SelectAll()).SelectRelease(&ret); err != nil {
		return err
	}
	pretty.NewTable(
		pretty.Header("#", "KeyspaceName", "DurableWrites", "Replication"),
		pretty.Rows(lo.Map(ret, func(v Keyspace, i int) []any {
			return []any{i, v.KeyspaceName, v.DurableWrites, v.Replication}
		})...),
	).Write(os.Stdout)
	return nil
}
func CreateKeySpace(session gocqlx.Session, in Keyspace) error {
	return session.Query(models.Keyspaces.Insert()).BindStruct(in).Exec()
}
func AlterKeySpace(session gocqlx.Session, in Keyspace) error {
	return session.Query(models.Keyspaces.Update()).BindStruct(in).Exec()
}
func DropKeySpace(session gocqlx.Session, in Keyspace) error {
	return session.Query(models.Keyspaces.Delete()).BindStruct(in).Exec()
}
