package resolver

import (
	"fmt"
	"golang.org/x/tools/go/packages"
)

type PackagesResolver struct {
	Dir    string
	Config packages.Config
	Hints  map[string]string
}

var LoadMode = packages.NeedName |
	packages.NeedFiles |
	packages.NeedCompiledGoFiles |
	packages.NeedImports |
	packages.NeedTypes |
	packages.NeedTypesSizes |
	packages.NeedSyntax |
	packages.NeedTypesInfo

func NewPackagesResolver(dir string) *PackagesResolver {
	return &PackagesResolver{Dir: dir}
}
func (r *PackagesResolver) WithConfig(config packages.Config) *PackagesResolver {
	r.Config = config
	return r
}
func (r *PackagesResolver) WithHints(hints map[string]string) *PackagesResolver {
	r.Hints = hints
	return r
}
func (r *PackagesResolver) ResolvePackage(path string) (string, error) {
	if name, ok := r.Hints[path]; ok {
		return name, nil
	} else if r.Dir != "" {
		r.Config.Dir = r.Dir
	}
	r.Config.Mode = LoadMode
	r.Config.Tests = false
	if pkgs, err := packages.Load(&r.Config, "pattern="+path); err != nil {
		return "", err
	} else if len(pkgs) > 1 {
		return "", fmt.Errorf("%d packages found for %s, %s", len(pkgs), path, r.Config.Dir)
	} else if len(pkgs) == 0 {
		return "", ErrPackageNotFound
	} else if p := pkgs[0]; len(p.Errors) > 0 {
		return "", p.Errors[0]
	} else {
		return p.Name, nil
	}
}
