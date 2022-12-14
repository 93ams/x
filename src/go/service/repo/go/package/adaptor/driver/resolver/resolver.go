package resolver

import (
	"errors"
	"go/ast"
)

type RestorerResolver interface {
	ResolvePackage(path string) (string, error)
}
type DecoratorResolver interface {
	ResolveIdent(*ast.File, ast.Node, string, *ast.Ident) (string, error)
}

var ErrPackageNotFound = errors.New("package not found")
