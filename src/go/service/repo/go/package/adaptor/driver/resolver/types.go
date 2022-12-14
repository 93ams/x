package resolver

import (
	"errors"
	"go/ast"
	"go/types"
)

type TypesResolver map[*ast.Ident]types.Object

func NewTypesResolver(uses map[*ast.Ident]types.Object) TypesResolver { return uses }
func (r TypesResolver) ResolveIdent(file *ast.File, parent ast.Node, parentField string, id *ast.Ident) (string, error) {
	if r == nil {
		return "", errors.New("gotypes.TypesResolver needs Uses in types info")
	} else if se, ok := parent.(*ast.SelectorExpr); ok && parentField == "Sel" {
		if xid, ok := se.X.(*ast.Ident); !ok {
			return "", nil
		} else if obj, ok := r[xid]; !ok {
			return "", nil
		} else if pn, ok := obj.(*types.PkgName); !ok {
			return "", nil
		} else {
			return pn.Imported().Path(), nil
		}
	} else if obj, ok := r[id]; !ok {
		return "", nil
	} else if v, ok := obj.(*types.Var); ok && v.IsField() {
		return "", nil
	} else if pkg := obj.Pkg(); pkg == nil {
		return "", nil
	} else {
		return pkg.Path(), nil
	}
}
