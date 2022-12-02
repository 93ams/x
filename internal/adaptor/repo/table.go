package repo

import (
	. "context"
	"github.com/scylladb/gocqlx/v2"
	. "github.com/tilau2328/cql/internal/adaptor/repo/model"
	. "github.com/tilau2328/cql/package/cql"
	. "github.com/tilau2328/cql/package/patch"
	"strings"
)

type (
	TableProvider interface {
		List(Context, Table) ([]Table, error)
		Get(Context, TableKey) (Table, error)
		Create(Context, Table) error
		Alter(Context, Table) error
		Drop(Context, TableKey) error
	}
	TableRepo struct{ session gocqlx.Session }
)

func NewTableRepo(session gocqlx.Session) *TableRepo { return &TableRepo{session: session} }
func (s *TableRepo) Create(ctx Context, table Table) error {
	return SafeExec(ctx, s.session, `CREATE table IF NOT EXISTS `+
		name+` (`+strings.Join(table, ",")+`)`)
}
func (s *TableRepo) Alter(ctx Context, name TableKey, props []Patch) error {
	return SafeExec(ctx, s.session, `ALTER table `+
		name.String()+` (`+strings.Join(props, ",")+`)`)
}
func (s *TableRepo) Drop(ctx Context, name TableKey) error {
	return SafeExec(ctx, s.session, `DROP table IF EXISTS `+name.String())
}
func (s *TableRepo) List(ctx Context, filter Table) (ret []Table, err error) {
	stmt, names := Tables.SelectAll()
	if err = s.session.ContextQuery(ctx, stmt, names).
		BindStruct(filter).SelectRelease(&ret); err != nil {
		return
	}
	return
}
func (s *TableRepo) Get(ctx Context, name TableKey) (ret Table, err error) {
	stmt, names := Tables.Get()
	if err = s.session.ContextQuery(ctx, stmt, names).
		BindStruct(Table{Name: name.String()}).
		SelectRelease(&ret); err != nil {
		return
	}
	return
}
