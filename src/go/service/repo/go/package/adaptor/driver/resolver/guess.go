package resolver

import "strings"

type GuessResolver map[string]string

func NewGuessResolver() GuessResolver { return GuessResolver{} }
func (r GuessResolver) ResolvePackage(importPath string) (string, error) {
	if n, ok := r[importPath]; ok {
		return n, nil
	} else if !strings.Contains(importPath, "/") {
		return importPath, nil
	}
	return importPath[strings.LastIndex(importPath, "/")+1:], nil
}
