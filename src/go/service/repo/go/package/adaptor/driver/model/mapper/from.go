package mapper

import (
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model"
	. "github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
)

func MapType(p *model.Type) any {
	return nil
}
func MapIdent(p *model.Ident) Ident {
	return Ident{
		Name: p.Name,
		Path: p.Path,
	}
}
func MapNames(i []*model.Ident) []string {
	return lo.Map(i, func(item *model.Ident, index int) string {
		return item.Name
	})
}
func MapField(item *model.Field) Field {
	ret := Field{Names: MapNames(item.Names)}
	switch t := item.Type.(type) {
	case *model.Ident:
		ret.Type = MapIdent(t)
	case *model.Selector:
		ret.Type = Ident{
			Name: t.X.(*model.Ident).Name,
			Path: t.Sel.Name,
		}
	}
	return ret
}
func MapFields(l *model.FieldList) []Field {
	return lo.Map(l.List, func(item *model.Field, index int) Field {
		return MapField(item)
	})
}
func MapStruct(t *model.Type, p *model.Struct) Struct {
	return Struct{
		Ident:  MapIdent(t.Name),
		Fields: MapFields(p.Fields),
	}
}
func MapInterface(t *model.Type, p *model.Interface) Interface {
	return Interface{
		Ident:   MapIdent(t.Name),
		Methods: MapMethods(p.Methods),
	}
}
func MapMethods(methods *model.FieldList) []FuncType {
	return lo.Map(methods.List, func(item *model.Field, index int) FuncType {
		return FuncType{
			Name: item.Names[0].Name,
			In:   MapFields(item.Type.(*model.FuncType).Params),
			Out:  MapFields(item.Type.(*model.FuncType).Results),
		}
	})
}
func MapExpr(expr model.Expr) {

}
func MapDecl() {

}
