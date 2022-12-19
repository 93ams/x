package service

import (
	"github.com/tilau2328/x/src/go/services/repo/go/package/adaptor/driver/mapper"
	"github.com/tilau2328/x/src/go/services/repo/go/package/wrapper"
	"github.com/tilau2328/x/src/go/services/repo/go/package/wrapper/builder"
	"io"
	"os"
)

type (
	CreateReq struct {
		Pkg, File string
		Props     any
	}
	SearchReq struct {
	}
	TransformReq struct {
	}
	Service struct {
	}
)

func NewService() *Service {
	return &Service{}
}

func (d *Service) Create(req CreateReq) error {
	f, err := os.OpenFile(req.File, os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	return WriteDecls(f, req.Pkg, mapper.NewDecls(req.Props)...)
}
func (d *Service) Search(req SearchReq) {

}
func (d *Service) Transform(req TransformReq) {

}
func WriteDecls(w io.Writer, pkg string, d ...builder.DeclBuilder) error {
	return wrapper.Write(w, builder.File(builder.Ident(pkg)).Decls(d...).Build())
}
