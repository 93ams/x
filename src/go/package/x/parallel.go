package x

import (
	"sync"
)

type WaitGroup struct {
	errLock sync.Mutex
	Errors  []error
	sync.WaitGroup
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{Errors: make([]error, 0)}
}
func (g *WaitGroup) Run(fn func() error) {
	g.Add(1)
	go func() {
		defer g.Done()
		if err := fn(); err != nil {
			g.errLock.Lock()
			g.Errors = append(g.Errors, err)
			g.errLock.Unlock()
		}
	}()
}
func (g *WaitGroup) Wait() error {
	g.WaitGroup.Wait()
	return MultiErr(g.Errors, "")
}

func ConcurrentTry(fns ...func() error) error {
	wg := NewWaitGroup()
	for _, fn := range fns {
		wg.Run(fn)
	}
	return wg.Wait()
}
func ParallelTry[T any](fn func(T) error, args []T) error {
	wg := NewWaitGroup()
	for _, v := range args {
		val := v
		wg.Run(func() error { return fn(val) })
	}
	return wg.Wait()
}
func ParallelMapTry[K comparable, V any](fn func(K, V) error, args map[K]V) error {
	wg := NewWaitGroup()
	for k, v := range args {
		key, val := k, v
		wg.Run(func() error { return fn(key, val) })
	}
	return wg.Wait()
}
