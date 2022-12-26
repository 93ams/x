package provider

import (
	"context"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
)

type GolangProvider interface {
	Create(context.Context, ...model.CreateReq) error
	Modify(context.Context, ...model.ModifyReq) error
	Delete(context.Context, ...model.DeleteReq) error
	Read(context.Context, ...model.ReadReq) (map[string]*model.File, error)
	Search(context.Context, ...model.SearchReq) (map[string][]model.Node, error)
}
