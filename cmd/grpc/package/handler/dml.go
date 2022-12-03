package handler

import (
	"github.com/tilau2328/cql/package/domain/provider"
	"grpc/package/model"
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
