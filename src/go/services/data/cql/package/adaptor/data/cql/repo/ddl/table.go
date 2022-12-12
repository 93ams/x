package ddl

import (
	. "context"
	"github.com/scylladb/gocqlx/v2"
	"github.com/tilau2328/cql/src/go/package/adaptor/data/cql/model"
	. "github.com/tilau2328/cql/src/go/package/domain/model"
	. "github.com/tilau2328/cql/src/go/package/domain/provider"
	"github.com/tilau2328/cql/src/go/package/shared/data/cql/ddl"
)

type TableRepo struct{ session gocqlx.Session }

var _ TableProvider = &TableRepo{}

func NewTableRepo(session gocqlx.Session) *TableRepo { return &TableRepo{session: session} }
func (s *TableRepo) Create(ctx Context, table Table) error {
	stmt := ddl.CreateTable(table.KeySpace, table.Name).Cols(model.MapCols(table.Columns)...)
	return SafeExec(ctx, s.session, stmt.String())
}
func (s *TableRepo) Alter(ctx Context, name TableKey, props []Patch) error {
	//return SafeExec(ctx, s.session, `ALTER table `+
	//	name.String()+` (`+strings.Join(props, ",")+`)`)
	return nil
}
func (s *TableRepo) Drop(ctx Context, key TableKey) error {
	return SafeExec(ctx, s.session, ddl.DropTable(key.KeySpace, key.Name).String())
}
func (s *TableRepo) List(ctx Context, filter Table) ([]Table, error) {
	stmt, names := model.Tables.Select()
	var ret []model.Table
	if err := s.session.ContextQuery(ctx, stmt, names).
		BindStruct(model.FromTable(filter)).
		SelectRelease(&ret); err != nil {
		return nil, err
	}
	return model.ToTables(ret), nil
}
func (s *TableRepo) Get(ctx Context, key TableKey) (Table, error) {
	stmt, names := model.Tables.Get()
	var ret model.Table
	if err := s.session.ContextQuery(ctx, stmt, names).
		BindStruct(model.Table{Name: key.Name, KeySpace: key.KeySpace}).
		GetRelease(&ret); err != nil {
		return Table{}, err
	}
	return model.ToTable(ret), nil
}
