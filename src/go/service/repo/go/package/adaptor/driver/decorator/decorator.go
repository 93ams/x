package decorator

import (
	"fmt"
	resolver2 "github.com/tilau2328/cql/src/go/services/gen/go/package/adaptor/driver/resolver"
	model2 "github.com/tilau2328/cql/src/go/services/gen/go/package/domain/model"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/packages"
	"os"
	"strings"
)

type (
	Decorator struct {
		model2.Map
		Filenames        map[*model2.File]string
		Fset             *token.FileSet
		Resolver         resolver2.DecoratorResolver
		Path             string
		ResolveLocalPath bool
	}
	fileDecorator struct {
		*Decorator
		file          *ast.File
		cursor        int
		fragments     []fragment
		startIndents  map[ast.Node]int
		endIndents    map[ast.Node]int
		before, after map[ast.Node]model2.SpaceType
		decorations   map[ast.Node]map[string][]string
	}
)

func NewDecorator(fset *token.FileSet) *Decorator {
	if fset == nil {
		fset = token.NewFileSet()
	}
	return &Decorator{
		Map:       model2.NewMap(),
		Filenames: map[*model2.File]string{},
		Fset:      fset,
	}
}
func NewDecoratorWithImports(fset *token.FileSet, path string, resolver resolver2.DecoratorResolver) *Decorator {
	dec := NewDecorator(fset)
	dec.Path = path
	dec.Resolver = resolver
	return dec
}
func NewDecoratorFromPackage(pkg *packages.Package) *Decorator {
	return NewDecoratorWithImports(pkg.Fset, pkg.PkgPath, resolver2.NewTypesResolver(pkg.TypesInfo.Uses))
}
func (d *Decorator) Parse(src interface{}) (*model2.File, error) {
	return d.ParseFile("", src, parser.ParseComments)
}
func (d *Decorator) ParseFile(filename string, src interface{}, mode parser.Mode) (*model2.File, error) {
	if f, perr := parser.ParseFile(d.Fset, filename, src, mode|parser.ParseComments); perr != nil && f == nil {
		return nil, perr
	} else if file, err := d.DecorateFile(f); err != nil {
		return nil, err
	} else {
		return file, perr
	}
}
func (d *Decorator) ParseDir(dir string, filter func(os.FileInfo) bool, mode parser.Mode) (map[string]*model2.Package, error) {
	pkgs, err := parser.ParseDir(d.Fset, dir, filter, mode|parser.ParseComments)
	if err != nil {
		return nil, err
	}
	out := map[string]*model2.Package{}
	for k, v := range pkgs {
		pkg, err := d.DecorateNode(v)
		if err != nil {
			return nil, err
		}
		out[k] = pkg.(*model2.Package)
	}
	return out, nil
}
func (d *Decorator) DecorateFile(f *ast.File) (*model2.File, error) {
	file, err := d.DecorateNode(f)
	if err != nil {
		return nil, err
	}
	return file.(*model2.File), nil
}
func (d *Decorator) DecorateNode(n ast.Node) (model2.Node, error) {
	if d.Resolver == nil && d.Path != "" {
		panic("Decorator Path should be empty when Resolver is nil")
	} else if d.Resolver != nil && d.Path == "" {
		panic("Decorator Path should be set when Resolver is set")
	}
	fd := d.newFileDecorator()
	if f, ok := n.(*ast.File); ok {
		fd.file = f
	}
	fd.fragment(n)
	fd.link()
	out, err := fd.decorateNode(nil, "", "", "", n)
	if err != nil {
		return nil, err
	}
	switch n := n.(type) {
	case *ast.Package:
		for k, v := range n.Files {
			d.Filenames[d.Dst.Nodes[v].(*model2.File)] = k
		}
	case *ast.File:
		d.Filenames[out.(*model2.File)] = d.Fset.File(n.Pos()).Name()
	}

	return out, nil
}
func (pd *Decorator) newFileDecorator() *fileDecorator {
	return &fileDecorator{
		Decorator:    pd,
		startIndents: map[ast.Node]int{},
		endIndents:   map[ast.Node]int{},
		before:       map[ast.Node]model2.SpaceType{},
		after:        map[ast.Node]model2.SpaceType{},
		decorations:  map[ast.Node]map[string][]string{},
	}
}
func (f *fileDecorator) decorateSelectorExpr(parent ast.Node, parentName, parentField, parentFieldType string, n *ast.SelectorExpr) (model2.Node, error) {
	if f.Resolver == nil {
		return nil, nil
	}
	path, err := f.resolvePath(true, n, "Selector", "Sel", "Ident", n.Sel)
	if err != nil {
		return nil, err
	}

	if path == "" {
		return nil, nil
	}
	out := &model2.Ident{}
	f.Dst.Nodes[n] = out
	f.Dst.Nodes[n.X] = out
	f.Dst.Nodes[n.Sel] = out
	f.Ast.Nodes[out] = n
	out.Name = n.Sel.Name
	ob, err := f.decorateObject(n.Sel.Obj)
	if err != nil {
		return nil, err
	}
	out.Obj = ob
	out.Path = path
	out.Decs.Before = f.before[n]
	out.Decs.After = f.after[n]
	var nStart, xBefore, xStart, xEnd, xAfter, nX, sBefore, sStart, sEnd, sAfter, nEnd interface{}
	xBefore = f.before[n.X]
	xAfter = f.after[n.X]
	sBefore = f.before[n.Sel]
	sAfter = f.after[n.Sel]
	if decs, ok := f.decorations[n]; ok {
		nStart = decs["Start"]
		nX = decs["X"]
		nEnd = decs["End"]
	}
	if decs, ok := f.decorations[n.X]; ok {
		xStart = decs["Start"]
		xEnd = decs["End"]
	}
	if decs, ok := f.decorations[n.Sel]; ok {
		sStart = decs["Start"]
		sEnd = decs["End"]
	}
	if iStart := mergeDecorations(nStart, xBefore, xStart); len(iStart) > 0 {
		out.Decs.Start.Append(iStart...)
	}
	if iX := mergeDecorations(xEnd, xAfter, nX, sBefore, sStart); len(iX) > 0 {
		out.Decs.X.Append(iX...)
	}
	if iEnd := mergeDecorations(sEnd, sAfter, nEnd); len(iEnd) > 0 {
		out.Decs.End.Append(iEnd...)
	}
	return out, nil
}
func (f *fileDecorator) resolvePath(force bool, parent ast.Node, parentName, parentField, parentFieldType string, id *ast.Ident) (string, error) {
	if f.Resolver == nil {
		panic("resolvePath needs a Resolver")
	} else if !force {
		if model2.Avoid[parentName+"."+parentField] {
			return "", nil
		} else if parentFieldType != "Expr" {
			panic(fmt.Sprintf("decorateIdent: unsupported parentName %s, parentField %s, parentFieldType %s", parentName, parentField, parentFieldType))
		}
	}
	if path, err := f.Resolver.ResolveIdent(f.file, parent, parentField, id); err != nil {
		return "", err
	} else if path = stripVendor(path); !f.ResolveLocalPath && path == stripVendor(f.Path) {
		return "", nil
	} else {
		return path, nil
	}
}
func stripVendor(path string) string {
	findVendor := func(path string) (index int, ok bool) {
		switch {
		case strings.Contains(path, "/vendor/"):
			return strings.LastIndex(path, "/vendor/") + 1, true
		case strings.HasPrefix(path, "vendor/"):
			return 0, true
		}
		return 0, false
	}
	i, ok := findVendor(path)
	if !ok {
		return path
	}
	return path[i+len("vendor/"):]
}
func (f *fileDecorator) decorateObject(o *ast.Object) (*model2.Object, error) {
	if o == nil {
		return nil, nil
	} else if do, ok := f.Dst.Objects[o]; ok {
		return do, nil
	}
	out := &model2.Object{}
	f.Dst.Objects[o] = out
	f.Ast.Objects[out] = o
	out.Kind = model2.ObjKind(o.Kind)
	out.Name = o.Name
	switch decl := o.Decl.(type) {
	case *ast.Scope:
		s, err := f.decorateScope(decl)
		if err != nil {
			return nil, err
		}
		out.Decl = s
	case ast.Node:
		n, err := f.decorateNode(nil, "", "", "", decl)
		if err != nil {
			return nil, err
		}
		out.Decl = n
	case nil:
	default:
		panic(fmt.Sprintf("o.Decl is %T", o.Data))
	}
	switch data := o.Data.(type) {
	case int:
		out.Data = data
	case *ast.Scope:
		s, err := f.decorateScope(data)
		if err != nil {
			return nil, err
		}
		out.Data = s
	case ast.Node:
		n, err := f.decorateNode(nil, "", "", "", data)
		if err != nil {
			return nil, err
		}
		out.Data = n
	case nil:
	default:
		panic(fmt.Sprintf("o.Data is %T", o.Data))
	}
	return out, nil
}
func (f *fileDecorator) decorateScope(s *ast.Scope) (*model2.Scope, error) {
	if s == nil {
		return nil, nil
	} else if ds, ok := f.Dst.Scopes[s]; ok {
		return ds, nil
	}
	out := &model2.Scope{}
	f.Dst.Scopes[s] = out
	f.Ast.Scopes[out] = s
	outer, err := f.decorateScope(s.Outer)
	if err != nil {
		return nil, err
	}
	out.Outer = outer
	out.Objects = map[string]*model2.Object{}
	for k, v := range s.Objects {
		ob, err := f.decorateObject(v)
		if err != nil {
			return nil, err
		}
		out.Objects[k] = ob
	}
	return out, nil
}
func mergeDecorations(decorationsOrLineSpace ...interface{}) []string {
	var endsWithNewLine bool
	var out []string
	for _, v := range decorationsOrLineSpace {
		switch v := v.(type) {
		case nil:
		case []string:
			if len(v) == 0 {
				continue
			}
			out = append(out, v...)
			endsWithNewLine = v[len(v)-1] == "\n" || strings.HasPrefix(v[len(v)-1], "//")
		case model2.SpaceType:
			switch v {
			case model2.NewLine:
				if endsWithNewLine {
					// nothing to do
				} else {
					out = append(out, "\n")
				}
				endsWithNewLine = true
			case model2.EmptyLine:
				if endsWithNewLine {
					out = append(out, "\n")
				} else {
					out = append(out, "\n", "\n")
				}
				endsWithNewLine = true
			}
		default:
			panic(fmt.Sprintf("%T", v))
		}
	}
	return out
}
