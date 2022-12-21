package provider

import (
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"io"
)

type GolangProvider interface {
	Create(io.Writer, model.CreateReq) error
	Search(io.Writer, model.SearchReq) error
	Modify(io.Writer, model.ModifyReq) error
}
