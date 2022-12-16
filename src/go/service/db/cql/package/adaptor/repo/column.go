package repo

import (
	. "context"
	"github.com/scylladb/gocqlx/v2"
	"github.com/tilau2328/cql/src/go/package/adaptor/data/cql/model"
	. "github.com/tilau2328/cql/src/go/package/domain/model"
	. "github.com/tilau2328/cql/src/go/package/domain/provider"
)

type ColumnRepo struct{ session gocqlx.Session }

var _ ColumnProvider = &ColumnRepo{}

func NewColumnRepo(session gocqlx.Session) *ColumnRepo { return &ColumnRepo{session: session} }
func (s *ColumnRepo) List(ctx Context, filter Column) ([]Column, error) {
	stmt, names := model.Columns.SelectAll()
	var ret []model.Column
	if err := s.session.ContextQuery(ctx, stmt, names).
		BindStruct(filter).SelectRelease(&ret); err != nil {
		return nil, err
	}
	return model.ToColumns(ret), nil
}
func (s *ColumnRepo) Get(ctx Context, key ColumnKey) (Column, error) {
	stmt, names := model.Columns.Get()
	var ret model.Column
	if err := s.session.ContextQuery(ctx, stmt, names).
		BindStruct(model.Column{Name: key.Name, KeySpace: key.KeySpace}).
		GetRelease(&ret); err != nil {
		return Column{}, err
	}
	return model.ToColumn(ret), nil
}
