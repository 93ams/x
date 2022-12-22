package service

import (
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/coding"
	model2 "github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model/builder"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model/mapper"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model/visitor/filter"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/provider"
	"go/parser"
	"go/token"
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

func (d *Service) Create(w io.Writer, req model.CreateReq, props any) error {
	if req.Pattern {

	}
	return WriteDecls(w, req.Pkg, mapper.NewDecls(props)...)
}
func (d *Service) Search(w io.Writer, req model.SearchReq, ret any) error {
	node, err := coding.ParseFile(token.NewFileSet(), req.File+".go", nil, parser.AllErrors)
	if err != nil {
		return err
	}
	model2.Inspect(node, Filter(func(node model2.Node) bool {
		switch t := node.(type) {
		case *model2.Type:
			return req.Name == "" || t.Name.Name == req.Name
		}
		return false
	}, ret))
	return nil
}
func (d *Service) Modify(w io.Writer, req model.ModifyReq, ret any) error {
	return nil
}
func WriteDecls(w io.Writer, pkg string, d ...builder.DeclBuilder) error {
	return coding.Write(w, builder.File(builder.Ident(pkg)).Decls(d...).Build())
}

func Filter(fn func(model2.Node) bool, ret any) func(model2.Node) bool {
	switch t := ret.(type) {
	case *model.Types:
		*t = model.Types{}
		return filter.OnType(func(m *model2.Type) bool {
			switch s := m.Type.(type) {
			case *model2.Struct:
				t.Struct = append(t.Struct, mapper.MapStruct(m, s))
			case *model2.Interface:
				t.Interfaces = append(t.Interfaces, mapper.MapInterface(m, s))
			}
			return true
		})
	case *model.Struct:
		return filter.OnSpecificType(func(m *model2.Type, s *model2.Struct) bool {
			if fn(m) {
				*t = mapper.MapStruct(m, s)
				return false
			}
			return true
		})
	case *[]model.Struct:
		return filter.OnSpecificType(func(m *model2.Type, s *model2.Struct) bool {
			if fn(m) {
				*t = append(*t, mapper.MapStruct(m, s))
			}
			return true
		})
	case *model.Interface:
		return filter.OnSpecificType(func(m *model2.Type, s *model2.Interface) bool {
			if fn(m) {
				*t = mapper.MapInterface(m, s)
				return false
			}
			return true
		})
	case *[]model.Interface:
		return filter.OnSpecificType(func(m *model2.Type, s *model2.Interface) bool {
			if fn(m) {
				*t = append(*t, mapper.MapInterface(m, s))
			}
			return true
		})
	}
	return nil
}
