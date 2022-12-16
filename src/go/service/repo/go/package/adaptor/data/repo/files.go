package repo

import (
	"context"
	"github.com/tilau2328/cql/src/go/services/gen/go/package/domain/model"
)

type FilesRepo struct {
}

func (r *FilesRepo) Create(ctx context.Context, action ...model.Mod) error {
	return nil
}
func (r *FilesRepo) Update(ctx context.Context, mod ...model.Mod) error {
	return nil
}
func (r *FilesRepo) Search(ctx context.Context, filter ...model.Filter) error {
	return nil
}
