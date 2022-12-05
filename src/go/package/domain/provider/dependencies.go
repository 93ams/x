package provider

import (
	. "context"
	. "github.com/tilau2328/cql/package/domain/model"
	. "github.com/tilau2328/cql/package/shared/patch"
)

//go:generate mockgen -source=dependencies.go -destination=dependencies_mock.go -package=provider
type (
	KeySpaceProvider interface {
		List(Context, KeySpace) ([]KeySpace, error)
		Get(Context, KeySpaceKey) (KeySpace, error)
		Create(Context, KeySpace) error
		Alter(Context, KeySpaceKey, []Patch) error
		Drop(Context, KeySpaceKey) error
	}
	TableProvider interface {
		List(Context, Table) ([]Table, error)
		Get(Context, TableKey) (Table, error)
		Create(Context, Table) error
		Alter(Context, TableKey, []Patch) error
		Drop(Context, TableKey) error
	}
	ColumnProvider interface {
		List(Context, Column) ([]Column, error)
		Get(Context, ColumnKey) (Column, error)
	}
	CrudProvider interface {
		Select(Context) error
		Insert(Context) error
		Update(Context) error
		Delete(Context) error
	}
)
