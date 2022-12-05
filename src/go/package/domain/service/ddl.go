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

func (D DDLService) ListKeySpaces(ctx Context, space KeySpace) ([]KeySpace, error) {
	//TODO implement me
	panic("implement me")
}

func (D DDLService) GetKeySpace(ctx Context, key KeySpaceKey) (KeySpace, error) {
	//TODO implement me
	panic("implement me")
}

func (D DDLService) CreateKeySpace(ctx Context, space KeySpace) error {
	//TODO implement me
	panic("implement me")
}

func (D DDLService) AlterKeySpace(ctx Context, key KeySpaceKey, patches []Patch) error {
	//TODO implement me
	panic("implement me")
}

func (D DDLService) DropKeySpace(ctx Context, key KeySpaceKey) error {
	//TODO implement me
	panic("implement me")
}

func (D DDLService) ListTables(ctx Context, table Table) ([]Table, error) {
	//TODO implement me
	panic("implement me")
}

func (D DDLService) GetTable(ctx Context, key TableKey) (Table, error) {
	//TODO implement me
	panic("implement me")
}

func (D DDLService) CreateTable(ctx Context, table Table) error {
	//TODO implement me
	panic("implement me")
}

func (D DDLService) AlterTable(ctx Context, key TableKey, patches []Patch) error {
	//TODO implement me
	panic("implement me")
}

func (D DDLService) DropTable(ctx Context, key TableKey) error {
	//TODO implement me
	panic("implement me")
}
