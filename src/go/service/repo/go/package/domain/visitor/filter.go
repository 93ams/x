package visitor

import (
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
)

type Seeker struct {
}

func On[T any](fn func(T) bool) func(model.Node) bool {
	return func(node model.Node) bool {
		if i, ok := node.(T); ok {
			return fn(i)
		}
		return true
	}
}
func OnType[T any](fn func(*model.Type, T) bool) func(model.Node) bool {
	return On(func(m *model.Type) bool {
		if s, ok := m.Type.(T); ok {
			return fn(m, s)
		}
		return true
	})
}
