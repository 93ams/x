package provider

import (
	. "context"
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
