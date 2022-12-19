package wrapper

import (
	"bytes"
	"errors"
	"github.com/tilau2328/x/src/go/services/repo/go/package/wrapper/coding/decorator"
	"github.com/tilau2328/x/src/go/services/repo/go/package/wrapper/coding/resolver"
	"github.com/tilau2328/x/src/go/services/repo/go/package/wrapper/coding/restorer"
	"github.com/tilau2328/x/src/go/services/repo/go/package/wrapper/model"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/tools/go/packages"
)

func Parse(src interface{}) (*model.File, error) {
	return decorator.NewDecorator(token.NewFileSet()).Parse(src)
}
func ParseFile(fset *token.FileSet, filename string, src interface{}, mode parser.Mode) (*model.File, error) {
	return decorator.NewDecorator(fset).ParseFile(filename, src, mode)
}
func ParseDir(fset *token.FileSet, dir string, filter func(os.FileInfo) bool, mode parser.Mode) (map[string]*model.Package, error) {
	return decorator.NewDecorator(fset).ParseDir(dir, filter, mode)
}
func Decorate(fset *token.FileSet, n ast.Node) (model.Node, error) {
	return decorator.NewDecorator(fset).DecorateNode(n)
}
func DecorateFile(fset *token.FileSet, f *ast.File) (*model.File, error) {
	return decorator.NewDecorator(fset).DecorateFile(f)
}
func Load(cfg *packages.Config, patterns ...string) ([]*Package, error) {
	if cfg == nil {
		cfg = &packages.Config{Mode: resolver.LoadMode}
	}
	if cfg.Mode&packages.NeedSyntax == 0 {
		return nil, errors.New("config mode should include NeedSyntax")
	}
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		return nil, err
	}
	dpkgs := map[*packages.Package]*Package{}
	var convert func(p *packages.Package) (*Package, error)
	convert = func(pkg *packages.Package) (*Package, error) {
		if dp, ok := dpkgs[pkg]; ok {
			return dp, nil
		}
		p := &Package{
			Package: pkg,
			Imports: map[string]*Package{},
		}
		dpkgs[pkg] = p
		if len(pkg.Syntax) > 0 {
			goFiles := make(map[string]bool, len(pkg.GoFiles))
			for _, fpath := range pkg.GoFiles {
				goFiles[fpath] = true
			}
			p.Decorator = decorator.NewDecoratorFromPackage(pkg)
			for _, f := range pkg.Syntax {
				fpath := pkg.Fset.File(f.Pos()).Name()
				if !goFiles[fpath] {
					continue
				}
				file, err := p.Decorator.DecorateFile(f)
				if err != nil {
					return nil, err
				}
				p.Syntax = append(p.Syntax, file)
			}
			dir, _ := filepath.Split(pkg.Fset.File(pkg.Syntax[0].Pos()).Name())
			p.Dir = dir
			for path, imp := range pkg.Imports {
				dimp, err := convert(imp)
				if err != nil {
					return nil, err
				}
				p.Imports[path] = dimp
			}
		}
		return p, nil
	}
	var out []*Package
	for _, pkg := range pkgs {
		p, err := convert(pkg)
		if err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, nil
}

type Package struct {
	*packages.Package
	Dir       string
	Decorator *decorator.Decorator
	Imports   map[string]*Package
	Syntax    []*model.File
}

func (p *Package) Save() error {
	return p.save(resolver.NewPackagesResolver(p.Dir), ioutil.WriteFile)
}
func (p *Package) SaveWithResolver(resolver resolver.RestorerResolver) error {
	return p.save(resolver, ioutil.WriteFile)
}

func (p *Package) save(resolver resolver.RestorerResolver, writeFile func(filename string, data []byte, perm os.FileMode) error) error {
	r := restorer.NewRestorerWithImports(p.PkgPath, resolver)
	for _, file := range p.Syntax {
		buf := &bytes.Buffer{}
		if err := r.Fprint(buf, file); err != nil {
			return err
		}
		if err := writeFile(p.Decorator.Filenames[file], buf.Bytes(), 0666); err != nil {
			return err
		}
	}
	return nil
}
