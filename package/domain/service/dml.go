package service

import (
	"context"
	"github.com/tilau2328/cql/package/domain/provider"
)

type DMLService struct {
}

var _ provider.DML = &DMLService{}

func NewDML() *DMLService {
	return &DMLService{}
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
