package ddl

import (
	. "context"
	"encoding/json"
	"github.com/scylladb/gocqlx/v2"
	"github.com/tilau2328/cql/package/adaptor/data/cql/model"
	. "github.com/tilau2328/cql/package/domain/model"
	. "github.com/tilau2328/cql/package/domain/provider"
	. "github.com/tilau2328/cql/package/shared/data/cql"
	. "github.com/tilau2328/cql/package/shared/patch"
	"strings"
)

type KeySpaceRepo struct{ session gocqlx.Session }

var _ KeySpaceProvider = &KeySpaceRepo{}

func NewKeySpaceRepo(session gocqlx.Session) *KeySpaceRepo { return &KeySpaceRepo{session: session} }
func (s *KeySpaceRepo) Create(ctx Context, keyspace KeySpace) error {
	replication, err := json.Marshal(keyspace.Replication)
	if err != nil {
		return err
	}
	stmt := "CREATE KEYSPACE IF NOT EXISTS " + string(keyspace.KeySpaceKey) +
		" WITH REPLICATION = " + strings.ReplaceAll(string(replication), "\"", "'")
	if !keyspace.Durable {
		stmt += " AND DURABLE_WRITES = false"
	}
	return SafeExec(ctx, s.session, stmt)
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
func (s *KeySpaceRepo) Drop(ctx Context, name KeySpaceKey) error {
	return SafeExec(ctx, s.session, "DROP KEYSPACE IF EXISTS "+name.String())
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
