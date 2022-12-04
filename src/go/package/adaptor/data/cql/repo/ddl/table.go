package ddl

import (
	. "context"
	"github.com/scylladb/gocqlx/v2"
	"github.com/tilau2328/cql/package/adaptor/data/cql/model"
	. "github.com/tilau2328/cql/package/domain/model"
	. "github.com/tilau2328/cql/package/domain/provider"
	. "github.com/tilau2328/cql/package/shared/data/cql"
	. "github.com/tilau2328/cql/package/shared/patch"
)

type TableRepo struct{ session gocqlx.Session }

var _ TableProvider = &TableRepo{}

func NewTableRepo(session gocqlx.Session) *TableRepo { return &TableRepo{session: session} }
func (s *TableRepo) Create(ctx Context, table Table) error {
	//return SafeExec(ctx, s.session, `CREATE table IF NOT EXISTS `+
	//	table.Name+` (`+strings.Join(table, ",")+`)`)
	return nil
}
func (s *TableRepo) Alter(ctx Context, name TableKey, props []Patch) error {
	//return SafeExec(ctx, s.session, `ALTER table `+
	//	name.String()+` (`+strings.Join(props, ",")+`)`)
	return nil
}
func (s *TableRepo) Drop(ctx Context, name TableKey) error {
	return SafeExec(ctx, s.session, `DROP table IF EXISTS `+name.String())
}
func (s *TableRepo) List(ctx Context, filter Table) (ret []Table, err error) {
	stmt, names := model.Tables.SelectAll()
	if err = s.session.ContextQuery(ctx, stmt, names).
		BindStruct(filter).SelectRelease(&ret); err != nil {
		return
	}
	return
}
func (s *TableRepo) Get(ctx Context, key TableKey) (ret Table, err error) {
	stmt, names := model.Tables.Get()
	if err = s.session.ContextQuery(ctx, stmt, names).
		BindStruct(model.Table{Name: key.Name, Keyspace: key.KeySpace}).
		GetRelease(&ret); err != nil {
		return
	}
	return
}
