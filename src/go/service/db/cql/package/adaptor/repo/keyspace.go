package repo

import (
	. "context"
	"github.com/scylladb/gocqlx/v2"
	"github.com/tilau2328/cql/src/go/package/adaptor/data/cql/model"
	. "github.com/tilau2328/cql/src/go/package/domain/model"
	. "github.com/tilau2328/cql/src/go/package/domain/provider"
	"github.com/tilau2328/cql/src/go/package/shared/data/cql/ddl"
)

type KeySpaceRepo struct{ session gocqlx.Session }

var _ KeySpaceProvider = &KeySpaceRepo{}

func NewKeySpaceRepo(session gocqlx.Session) *KeySpaceRepo { return &KeySpaceRepo{session: session} }
func (s *KeySpaceRepo) Create(ctx Context, keyspace KeySpace) error {
	stmt := ddl.CreateKeySpace(string(keyspace.KeySpaceKey), ddl.Object(model.FromReplication(keyspace.Replication)))
	if !keyspace.Durable {
		stmt = stmt.WithDurable(false)
	}
	return SafeExec(ctx, s.session, stmt.String())
}
func (s *KeySpaceRepo) Alter(ctx Context, key KeySpaceKey, patches []Patch) error {
	//stmt := "ALTER keyspace ? WITH REPLICATION = " + string(replication) + " WITH DURABLE_WRITES = "
	//if keyspace.Durable {
	//	stmt = "true"
	//} else {
	//	stmt = "false"
	//}
	//return SafeExec(ctx, s.session, stmt, key)
	return nil
}
func (s *KeySpaceRepo) Drop(ctx Context, key KeySpaceKey) error {
	return SafeExec(ctx, s.session, ddl.DropKeySpace(string(key)).String())
}
func (s *KeySpaceRepo) List(ctx Context, filter KeySpace) ([]KeySpace, error) {
	stmt, names := model.Keyspaces.SelectAll()
	var res []model.KeySpace
	if err := s.session.ContextQuery(ctx, stmt, names).SelectRelease(&res); err != nil {
		return nil, err
	}
	return model.ToKeySpaces(res), nil
}
func (s *KeySpaceRepo) Get(ctx Context, key KeySpaceKey) (KeySpace, error) {
	stmt, names := model.Keyspaces.Get()
	var res model.KeySpace
	if err := s.session.ContextQuery(ctx, stmt, names).
		BindStruct(model.KeySpace{Name: key.String()}).
		GetRelease(&res); err != nil {
		return KeySpace{}, err
	}
	return model.ToKeySpace(res), nil
}
