package resolver

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
	"sync"
)

func NewAstResolver() *AstResolver {
	return &AstResolver{}
}

func WithResolver(resolver RestorerResolver) *AstResolver {
	return &AstResolver{RestorerResolver: resolver}
}

type AstResolver struct {
	RestorerResolver RestorerResolver
	filesM           sync.Mutex
	files            map[*ast.File]map[string]string
}

func (r *AstResolver) ResolveIdent(file *ast.File, parent ast.Node, parentField string, id *ast.Ident) (string, error) {
	if r.RestorerResolver == nil {
		r.RestorerResolver = NewGuessResolver()
	}
	if imports, err := r.imports(file); err != nil {
		return "", err
	} else if se, ok := parent.(*ast.SelectorExpr); !ok || parentField != "Sel" {
		return "", nil
	} else if xid, ok := se.X.(*ast.Ident); !ok {
		return "", nil
	} else if xid.Obj != nil {
		return "", nil
	} else if path, ok := imports[xid.Name]; !ok {
		return "", nil
	} else {
		return path, nil
	}
}

func (r *AstResolver) imports(file *ast.File) (map[string]string, error) {
	r.filesM.Lock()
	defer r.filesM.Unlock()
	if r.files == nil {
		r.files = map[*ast.File]map[string]string{}
	}
	imports, ok := r.files[file]
	if ok {
		return imports, nil
	}
	imports = map[string]string{}
	var done bool
	var outer error
	ast.Inspect(file, func(node ast.Node) bool {
		if done || outer != nil {
			return false
		}
		switch node := node.(type) {
		case *ast.FuncDecl:
			done = true
			return false
		case *ast.GenDecl:
			if node.Tok != token.IMPORT {
				done = true
				return false
			}
			return true
		case *ast.ImportSpec:
			path := mustUnquote(node.Path.Value)
			if path == "C" {
				return false
			}
			var name string
			if node.Name != nil {
				name = node.Name.Name
			}
			switch name {
			case ".":
				outer = fmt.Errorf("goast.AstResolver unsupported dot-import found for %s", path)
				return false
			case "_":
				return false
			case "":
				var err error
				name, err = r.RestorerResolver.ResolvePackage(path)
				if err != nil {
					outer = err
					return false
				}
			}
			if p, ok := imports[name]; ok {
				outer = fmt.Errorf("goast.AstResolver found multiple packages using name %s: %s and %s", name, p, path)
				return false
			}
			imports[name] = path
		}
		return true
	})
	if outer != nil {
		return nil, outer
	}
	r.files[file] = imports
	return imports, nil
}

func mustUnquote(s string) string {
	out, err := strconv.Unquote(s)
	if err != nil {
		panic(err)
	}
	return out
}
