package service

import (
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/coding"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model/builder"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model/mapper"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/provider"
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
	return coding.Write(w, builder.File(builder.Ident(pkg)).Decls(d...).Build())
}
