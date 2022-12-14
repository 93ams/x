package resolver

type SimpleResolver map[string]string

func NewSimpleResolver(m map[string]string) SimpleResolver { return m }
func (r SimpleResolver) ResolvePackage(importPath string) (string, error) {
	if n, ok := r[importPath]; ok {
		return n, nil
	}
	return "", ErrPackageNotFound
}
