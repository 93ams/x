package filter

import (
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model"
)

func OnInterface(fn func(*model.Type, *model.Interface) bool) func(model.Node) bool {
	return OnType(func(m *model.Type) bool {
		if s, ok := m.Type.(*model.Interface); ok {
			return fn(m, s)
		}
		return true
	})
}
func OnSpecificType[T any](fn func(*model.Type, T) bool) func(model.Node) bool {
	return OnType(func(m *model.Type) bool {
		if s, ok := m.Type.(T); ok {
			return fn(m, s)
		}
		return true
	})
}
func OnType(fn func(*model.Type) bool) func(model.Node) bool {
	return func(node model.Node) bool {
		if i, ok := node.(*model.Type); ok {
			return fn(i)
		}
		return true
	}
}
