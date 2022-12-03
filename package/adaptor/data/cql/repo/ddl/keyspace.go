package ddl

import (
	. "context"
	"encoding/json"
	"github.com/scylladb/gocqlx/v2"
	"github.com/tilau2328/cql/package/adaptor/data/cql/model"
	. "github.com/tilau2328/cql/package/domain/model"
	"github.com/tilau2328/cql/package/domain/provider"
	. "github.com/tilau2328/cql/package/shared/data/cql"
	. "github.com/tilau2328/cql/package/shared/patch"
)

type KeySpaceRepo struct{ session gocqlx.Session }

var _ provider.KeySpaceProvider = &KeySpaceRepo{}

func NewKeySpaceRepo(session gocqlx.Session) *KeySpaceRepo { return &KeySpaceRepo{session: session} }
func (s *KeySpaceRepo) Create(ctx Context, keyspace KeySpace) error {
	replication, err := json.Marshal(keyspace.Replication)
	if err != nil {
		return err
	}
	stmt := "CREATE keyspace ? WITH REPLICATION = " + string(replication)
	if !keyspace.Durable {
		stmt += " WITH DURABLE_WRITES = false"
	}
	return SafeExec(ctx, s.session, stmt, keyspace.KeySpaceKey)
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
	return SafeExec(ctx, s.session, "DROP table IF EXISTS "+name.String())
}
func (s *KeySpaceRepo) List(ctx Context, filter KeySpace) (ret []KeySpace, err error) {
	stmt, names := model.Keyspaces.SelectAll()
	if err = s.session.ContextQuery(ctx, stmt, names).
		BindStruct(filter).SelectRelease(&ret); err != nil {
		return
	}
	return
}
func (s *KeySpaceRepo) Get(ctx Context, name KeySpaceKey) (ret KeySpace, err error) {
	stmt, names := model.Keyspaces.Get()
	if err = s.session.ContextQuery(ctx, stmt, names).
		BindStruct(model.KeySpace{Name: string(name)}).
		SelectRelease(&ret); err != nil {
		return
	}
	return
}

func toCqlKeySpace(in KeySpace) (ret model.KeySpace, err error) {
	return
}
func fromCqlKeySpace(in model.KeySpace) (ret KeySpace, err error) {
	return
}
