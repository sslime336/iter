package slice

import (
	"log"
	"sync"
)

func Iter[T any](sli []T) *Wrapper[T] {
	// w := pool.Get().(*Wrapper[T])
	// w.inner = sli
	// w.start, w.end = 0, len(sli)
	// w.funcChain = make([]any, 3)
	// w.collected = make([]T, 0, 5)
	return &Wrapper[T]{
		inner:     sli,
		start:     0,
		end:       len(sli),
		funcChain: make([]any, 3),
		collected: make([]T, 0, 5),
	}
}

/*
	sync.Pool does not fit!
 */

type SlicePool[T any] struct {
	sync.Pool
}

type typMapper[T comparable] struct {
	list map[T]SlicePool[T]
}

type Wrapper[T any] struct {
	inner      []T
	start, end int
	funcChain  []any
	collected  []T
}

func (w *Wrapper[T]) Range(start, end int) *Wrapper[T] {
	defer func() {
		if p := recover(); p != nil {
			log.Println(p)
		}
	}()
	if start < 0 || end < 0 || start > end {
		panic("illegal range")
	}
	w.start, w.end = start, end
	return w
}

func (w *Wrapper[T]) Filter(filterFunc func(T) bool) *Wrapper[T] {
	w.funcChain = append(w.funcChain, filterFunc)
	return w
}

func (w *Wrapper[T]) Map(mapFunc func(*T)) *Wrapper[T] {
	w.funcChain = append(w.funcChain, mapFunc)
	return w
}

func (w *Wrapper[T]) ForEach(handle func(*T)) {
	for i := w.start; i < w.end; i++ {
		handle(&w.inner[i])
	}
}

func (w *Wrapper[T]) Find(qualified func(T) bool) *T {
	// defer func() {
	// 	pool.Put(w)
	// }()
	for i := w.start; i < w.end; i++ {
		if qualified(w.inner[i]) {
			q := w.inner[i]
			return &q
		}
	}
	return nil
}

func (w *Wrapper[T]) Count() int {
	// defer func() {
	// 	pool.Put(w)
	// }()
	exec_funcChain(w)
	return len(w.collected)
}

// TODO: how to achieve
func (w *Wrapper[T]) Zip() {
}

// Collect will return the current slice, which has been copied from
// the Wrapper's inner(that has been dealed with chainFuncs).
func (w *Wrapper[T]) Collect() []T {
	// defer func() {
	// 	pool.Put(w)
	// }()
	exec_funcChain(w)
	collected := make([]T, len(w.collected))
	copy(collected, w.collected)
	return collected
}

// TODO: whether save Sum or not
func (w *Wrapper[T]) Sum() {

}

func exec_funcChain[T any](b *Wrapper[T]) {
	for _, chainFunc := range b.funcChain {
		switch t := chainFunc.(type) {
		case /* map functions */ func(*T):
			for i := b.start; i < b.end; i++ {
				t(&b.inner[i])
			}
		case /* fliter functions */ func(T) bool:
			for i := b.start; i < b.end; i++ {
				if t(b.inner[i]) {
					b.collected = append(b.collected, b.inner[i])
				}
			}
		default:
			panic("unmatched func type")
		}
	}
}
