package service

import (
	"context"
	"github.com/tilau2328/cql/src/go/package/domain/provider"
)

type (
	DMLServiceOptions struct {
	}
	DMLService struct {
		DMLServiceOptions
	}
)

var _ provider.DML = &DMLService{}

func NewDML(opts DMLServiceOptions) *DMLService {
	return &DMLService{DMLServiceOptions: opts}
}

func (D DMLService) Select(ctx context.Context) error {
	return nil
}

func (D DMLService) Insert(ctx context.Context) error {
	return nil
}

func (D DMLService) Update(ctx context.Context) error {
	return nil
}

func (D DMLService) Delete(ctx context.Context) error {
	return nil
}
