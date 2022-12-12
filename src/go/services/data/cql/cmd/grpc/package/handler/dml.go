package handler

import (
	"github.com/tilau2328/cql/src/go/cmd/grpc/package/model"
	provider2 "github.com/tilau2328/cql/src/go/package/domain/provider"
)

type (
	DMLOptions struct {
		provider2.DML
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
