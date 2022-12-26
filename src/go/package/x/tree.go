package x

type (
	Visitor[T comparable]     func(T) Visitor[T]
	Tree[T comparable]        map[T]Tree[T]
	PathVisitor[T comparable] []T
)

func (t Tree[T]) Add(path []T) Tree[T] {
	if len(path) == 0 {
		return t
	}
	curr := path[0]
	if n, ok := t[curr]; ok {
		t[curr] = n.Add(path[1:])
	} else {
		t[curr] = Tree[T]{}.Add(path[1:])
	}
	return t
}
func (t Tree[T]) Walk(v Visitor[T]) {
	for k, node := range t {
		node.Walk(v(k))
	}
}
func (w PathVisitor[T]) Visitor(fn func([]T)) Visitor[T] {
	return func(t T) Visitor[T] {
		p := append(w, t)
		fn(p)
		return p.Visitor(fn)
	}
}
