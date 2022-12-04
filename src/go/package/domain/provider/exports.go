package provider

import (
	. "context"
	. "github.com/tilau2328/cql/package/domain/model"
	. "github.com/tilau2328/cql/package/shared/patch"
)

type (
	DDL interface {
		ListKeySpaces(Context, KeySpace) ([]KeySpace, error)
		GetKeySpace(Context, KeySpaceKey) (KeySpace, error)
		CreateKeySpace(Context, KeySpace) error
		AlterKeySpace(Context, KeySpaceKey, []Patch) error
		DropKeySpace(Context, KeySpaceKey) error
		ListTables(Context, Table) ([]Table, error)
		GetTable(Context, TableKey) (Table, error)
		CreateTable(Context, Table) error
		AlterTable(Context, TableKey, []Patch) error
		DropTable(Context, TableKey) error
	}
	DML interface {
		Select(Context) error
		Insert(Context) error
		Update(Context) error
		Delete(Context) error
	}
)
