package repo

import (
	. "context"
	"encoding/json"
	"github.com/scylladb/gocqlx/v2"
	. "github.com/tilau2328/cql/internal/adaptor/repo/model"
	. "github.com/tilau2328/cql/package/cql"
)

type (
	KeySpaceProvider interface {
		List(Context, KeySpace) ([]KeySpace, error)
		Get(Context, KeySpaceKey) (KeySpace, error)
		Create(Context, KeySpace) error
		Alter(Context, KeySpace) error
		Drop(Context, string) error
	}
	KeySpaceRepo struct{ session gocqlx.Session }
)

func NewRepository(session gocqlx.Session) *KeySpaceRepo { return &KeySpaceRepo{session: session} }
func (s *KeySpaceRepo) Create(ctx Context, keyspace KeySpace) error {
	replication, err := json.Marshal(keyspace.Replication)
	if err != nil {
		return err
	}
	stmt := "CREATE keyspace ? WITH REPLICATION = " + string(replication)
	if !keyspace.Durable {
		stmt += " WITH DURABLE_WRITES = false"
	}
	return SafeExec(ctx, s.session, stmt, keyspace.Name)
}
func (s *KeySpaceRepo) Alter(ctx Context, keyspace KeySpace) error {
	replication, err := json.Marshal(keyspace.Replication)
	if err != nil {
		return err
	}
	stmt := "ALTER keyspace ? WITH REPLICATION = " + string(replication) + " WITH DURABLE_WRITES = "
	if keyspace.Durable {
		stmt = "true"
	} else {
		stmt = "false"
	}
	return SafeExec(ctx, s.session, stmt, keyspace.Name)
}
func (s *KeySpaceRepo) Drop(ctx Context, name KeySpaceKey) error {
	return SafeExec(ctx, s.session, "DROP table IF EXISTS "+name.String())
}
func (s *KeySpaceRepo) List(ctx Context, filter KeySpace) (ret []KeySpace, err error) {
	stmt, names := Keyspaces.SelectAll()
	if err = s.session.ContextQuery(ctx, stmt, names).
		BindStruct(filter).SelectRelease(&ret); err != nil {
		return
	}
	return
}
func (s *KeySpaceRepo) Get(ctx Context, name KeySpaceKey) (ret KeySpace, err error) {
	stmt, names := Keyspaces.Get()
	if err = s.session.ContextQuery(ctx, stmt, names).
		BindStruct(KeySpace{Name: name.String()}).
		SelectRelease(&ret); err != nil {
		return
	}
	return
}
