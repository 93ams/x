package x

import "github.com/samber/lo"

type (
	Builder[T any]  struct{ T T }
	IBuilder[T any] interface{ Build() T }
)

func MapBuilder[T any, B IBuilder[T]](builders []B) []T {
	return lo.Map(builders, func(item B, _ int) T { return item.Build() })
}
