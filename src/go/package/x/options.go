package x

type (
	Opt[T any]  func(T)
	OptE[T any] func(T) error
)

func IfNotNil[T any](t *T) (ret T) {
	if t != nil {
		ret = *t
	}
	return
}
func Apply[T any](t T, opts []Opt[T]) T {
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func ApplyE[T any](v []T, fn OptE[T]) error {
	for _, i := range v {
		if err := fn(i); err != nil {
			return err
		}
	}
	return nil
}
