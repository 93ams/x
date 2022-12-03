package ddl

import (
	. "context"
	"github.com/scylladb/gocqlx/v2"
	"github.com/tilau2328/cql/package/adaptor/data/cql/model"
	model2 "github.com/tilau2328/cql/package/domain/model"
	provider2 "github.com/tilau2328/cql/package/domain/provider"
	. "github.com/tilau2328/cql/src/go/package/shared/data/cql"
	. "github.com/tilau2328/cql/src/go/package/shared/patch"
)

type TableRepo struct{ session gocqlx.Session }

var _ provider2.TableProvider = &TableRepo{}

func NewTableRepo(session gocqlx.Session) *TableRepo { return &TableRepo{session: session} }
func (s *TableRepo) Create(ctx Context, table model2.Table) error {
	//return SafeExec(ctx, s.session, `CREATE table IF NOT EXISTS `+
	//	table.Name+` (`+strings.Join(table, ",")+`)`)
	return nil
}
func (s *TableRepo) Alter(ctx Context, name model2.TableKey, props []Patch) error {
	//return SafeExec(ctx, s.session, `ALTER table `+
	//	name.String()+` (`+strings.Join(props, ",")+`)`)
	return nil
}
func (s *TableRepo) Drop(ctx Context, name model2.TableKey) error {
	return SafeExec(ctx, s.session, `DROP table IF EXISTS `+name.String())
}
func (s *TableRepo) List(ctx Context, filter model2.Table) (ret []model2.Table, err error) {
	stmt, names := model.Tables.SelectAll()
	if err = s.session.ContextQuery(ctx, stmt, names).
		BindStruct(filter).SelectRelease(&ret); err != nil {
		return
	}
	return
}
func (s *TableRepo) Get(ctx Context, name model2.TableKey) (ret model2.Table, err error) {
	stmt, names := model.Tables.Get()
	if err = s.session.ContextQuery(ctx, stmt, names).
		BindStruct(model.Table{Name: name.String()}).
		SelectRelease(&ret); err != nil {
		return
	}
	return
}

func toCqlTable(in model2.Table) (ret model.Table, err error) {
	return
}
func fromCqlTable(in model.Table) (ret model2.Table, err error) {
	return
}
