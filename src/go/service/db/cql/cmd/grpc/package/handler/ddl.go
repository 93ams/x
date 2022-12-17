package handler

import (
	"context"
	"github.com/tilau2328/x/src/go/cmd/grpc/package/model"
	"github.com/tilau2328/x/src/go/package/domain/provider"
)

type (
	DDLOptions struct {
		provider.DDL
	}
	DDLHandler struct {
		model.UnimplementedDDLServer
		DDLOptions
	}
)

var _ model.DDLServer = &DDLHandler{}

func NewDDLHandler(opts DDLOptions) *DDLHandler {
	return &DDLHandler{
		DDLOptions: opts,
	}
}

func (r *DDLHandler) CreateKeySpaces(ctx context.Context, request *model.CreateKeySpacesRequest) (*model.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DDLHandler) AlterKeySpaces(ctx context.Context, request *model.AlterKeySpacesRequest) (*model.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DDLHandler) DropKeySpaces(ctx context.Context, request *model.DropKeySpacesRequest) (*model.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DDLHandler) ListKeySpaces(ctx context.Context, request *model.ListKeySpacesRequest) (*model.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DDLHandler) GetKeySpace(ctx context.Context, empty *model.Empty) (*model.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DDLHandler) CreateTables(ctx context.Context, request *model.CreateTablesRequest) (*model.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DDLHandler) AlterTables(ctx context.Context, request *model.AlterTablesRequest) (*model.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DDLHandler) DropTables(ctx context.Context, request *model.DropTablesRequest) (*model.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DDLHandler) ListTables(ctx context.Context, request *model.ListTablesRequest) (*model.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DDLHandler) GetTable(ctx context.Context, empty *model.Empty) (*model.Empty, error) {
	//TODO implement me
	panic("implement me")
}
