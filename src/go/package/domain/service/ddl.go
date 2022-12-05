package service

import (
	. "context"
	. "github.com/tilau2328/cql/src/go/package/domain/model"
	. "github.com/tilau2328/cql/src/go/package/domain/provider"
	. "github.com/tilau2328/cql/src/go/package/shared/patch"
)

type (
	DDLServiceOptions struct {
		KeySpaceProvider
		TableProvider
	}
	DDLService struct {
		DDLServiceOptions
	}
)

var _ DDL = &DDLService{}

func NewDDL(opts DDLServiceOptions) *DDLService {
	return &DDLService{DDLServiceOptions: opts}
}

func (D DDLService) ListKeySpace(ctx Context, space KeySpace) ([]KeySpace, error) {
	return nil, nil
}

func (D DDLService) GetKeySpace(ctx Context, key KeySpaceKey) (KeySpace, error) {
	return KeySpace{}, nil
}

func (D DDLService) CreateKeySpace(ctx Context, space KeySpace) error {
	return nil
}

func (D DDLService) AlterKeySpace(ctx Context, key KeySpaceKey, patches []Patch) error {
	return nil
}

func (D DDLService) DropKeySpace(ctx Context, key KeySpaceKey) error {
	return nil
}

func (D DDLService) ListTable(ctx Context, table Table) ([]Table, error) {
	return nil, nil
}

func (D DDLService) GetTable(ctx Context, key TableKey) (Table, error) {
	return Table{}, nil
}

func (D DDLService) CreateTable(ctx Context, table Table) error {
	return nil
}

func (D DDLService) AlterTable(ctx Context, key TableKey, patches []Patch) error {
	return nil
}

func (D DDLService) DropTable(ctx Context, key TableKey) error {
	return nil
}
