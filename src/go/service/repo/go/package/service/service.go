package service

import (
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/mapper"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/provider"
	"github.com/tilau2328/x/src/go/service/repo/go/package/wrapper"
	"github.com/tilau2328/x/src/go/service/repo/go/package/wrapper/builder"
	"io"
)

type (
	Service struct {
	}
)

var _ provider.GolangProvider = &Service{}

func NewService() *Service {
	return &Service{}
}

func (d *Service) Create(w io.Writer, req model.CreateReq) error {
	if req.Pattern {

	}
	return WriteDecls(w, req.Pkg, mapper.NewDecls(req.Props)...)
}
func (d *Service) Search(w io.Writer, req model.SearchReq) error {
	return nil
}
func (d *Service) Modify(w io.Writer, req model.ModifyReq) error {
	return nil
}
func WriteDecls(w io.Writer, pkg string, d ...builder.DeclBuilder) error {
	return wrapper.Write(w, builder.File(builder.Ident(pkg)).Decls(d...).Build())
}
