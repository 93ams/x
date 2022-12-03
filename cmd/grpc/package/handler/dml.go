package handler

import (
	"github.com/tilau2328/cql/cmd/grpc/package/model"
	"github.com/tilau2328/cql/package/domain/provider"
)

type (
	DMLOptions struct {
		provider.DML
	}
	DML struct {
		model.UnimplementedDMLServer
		DMLOptions
	}
)

var _ model.DMLServer = &DML{}

func NewDML(opts DMLOptions) *DML {
	return &DML{
		DMLOptions: opts,
	}
}
