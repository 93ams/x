package assert

import model2 "github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model"

func TypeMatch[T any](name string, fn func(*model2.Type, T)) func(*model2.Type, T) bool {
	return func(t *model2.Type, t2 T) bool {
		if t.Name.Name == name {
			fn(t, t2)
			return false
		}
		return true
	}
}
