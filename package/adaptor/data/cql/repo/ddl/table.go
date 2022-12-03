package ddl

import (
	. "context"
	"github.com/scylladb/gocqlx/v2"
	. "github.com/tilau2328/cql/internal/adaptor/cql/model"
	. "github.com/tilau2328/cql/package/adaptor/cql/model"
	. "github.com/tilau2328/cql/package/adaptor/data/cql/model"
	"github.com/tilau2328/cql/package/domain/model"
	"github.com/tilau2328/cql/package/domain/provider"
	. "github.com/tilau2328/cql/package/shared/adaptor/cql/model"
	. "github.com/tilau2328/cql/package/shared/cql"
	. "github.com/tilau2328/cql/package/shared/patch"
	"strings"
)

type TableRepo struct{ session gocqlx.Session }

var _ provider.TableProvider = &TableRepo{}

func NewTableRepo(session gocqlx.Session) *TableRepo { return &TableRepo{session: session} }
func (s *TableRepo) Create(ctx Context, table model.Table) error {
	return SafeExec(ctx, s.session, `CREATE table IF NOT EXISTS `+
		name+` (`+strings.Join(table, ",")+`)`)
}
func (s *TableRepo) Alter(ctx Context, name model.TableKey, props []Patch) error {
	return SafeExec(ctx, s.session, `ALTER table `+
		name.String()+` (`+strings.Join(props, ",")+`)`)
}
func (s *TableRepo) Drop(ctx Context, name model.TableKey) error {
	return SafeExec(ctx, s.session, `DROP table IF EXISTS `+name.String())
}
func (s *TableRepo) List(ctx Context, filter model.Table) (ret []Table, err error) {
	stmt, names := Tables.SelectAll()
	if err = s.session.ContextQuery(ctx, stmt, names).
		BindStruct(filter).SelectRelease(&ret); err != nil {
		return
	}
	return
}
func (s *TableRepo) Get(ctx Context, name model.TableKey) (ret Table, err error) {
	stmt, names := Tables.Get()
	if err = s.session.ContextQuery(ctx, stmt, names).
		BindStruct(Table{Name: name.String()}).
		SelectRelease(&ret); err != nil {
		return
	}
	return
}

func toCqlTable() {

}
func fromCqlTable() {

}
