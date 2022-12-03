package handler

import (
	"context"
	"github.com/tilau2328/cql/cmd/grpc/package/model"
	"github.com/tilau2328/cql/package/domain/provider"
)

type (
	DDLOptions struct {
		provider.DDL
	}
	DDL struct {
		model.UnimplementedDDLServer
		DDLOptions
	}
)

var _ model.DDLServer = &DDL{}

func NewDDL(opts DDLOptions) *DDL {
	return &DDL{
		DDLOptions: opts,
	}
}

func (D DDL) CreateKeySpace(ctx context.Context, empty *model.Empty) (*model.Empty, error) {
	return nil, nil
}

func (D DDL) AlterKeySpace(ctx context.Context, empty *model.Empty) (*model.Empty, error) {
	return nil, nil
}

func (D DDL) DropKeySpace(ctx context.Context, empty *model.Empty) (*model.Empty, error) {
	return nil, nil
}

func (D DDL) ListKeySpaces(ctx context.Context, empty *model.Empty) (*model.Empty, error) {
	return nil, nil
}

func (D DDL) GetKeySpace(ctx context.Context, empty *model.Empty) (*model.Empty, error) {
	return nil, nil
}

func (D DDL) CreateTable(ctx context.Context, empty *model.Empty) (*model.Empty, error) {
	return nil, nil
}

func (D DDL) AlterTable(ctx context.Context, empty *model.Empty) (*model.Empty, error) {
	return nil, nil
}

func (D DDL) DropTable(ctx context.Context, empty *model.Empty) (*model.Empty, error) {
	return nil, nil
}

func (D DDL) ListTables(ctx context.Context, empty *model.Empty) (*model.Empty, error) {
	return nil, nil
}

func (D DDL) GetTable(ctx context.Context, empty *model.Empty) (*model.Empty, error) {
	return nil, nil
}
