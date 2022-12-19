package resolver

import (
	"go/build"
)

type BuildResolver struct {
	FindPackage func(ctxt *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error)
	Context     *build.Context
	Dir         string
	Hints       map[string]string
}

func NewBuildResolver(dir string) *BuildResolver { return &BuildResolver{Dir: dir} }
func (r *BuildResolver) WithContext(context *build.Context) *BuildResolver {
	r.Context = context
	return r
}
func (r *BuildResolver) WithHints(dir string, hints map[string]string) *BuildResolver {
	r.Hints = hints
	return r
}
func (r *BuildResolver) ResolvePackage(importPath string) (string, error) {
	fp := r.FindPackage
	if fp == nil {
		fp = (*build.Context).Import
	}
	bc := r.Context
	if bc == nil {
		bc = &build.Default
	}
	if name, ok := r.Hints[importPath]; ok {
		return name, nil
	} else if p, err := fp(bc, importPath, r.Dir, 0); err != nil {
		return "", err
	} else if p == nil {
		return "", ErrPackageNotFound
	} else {
		return p.Name, nil
	}
}
