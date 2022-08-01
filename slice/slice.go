package slice

import (
	"errors"
	"log"

	"github.com/sslime336/iter"
)

func Iter[T any](sli []T) iter.SliceIter[T] {
	return &Wrapper[T]{
		inner:     sli,
		start:     0,
		end:       len(sli),
		funcChain: make([]any, 3),
		collected: make([]T, 0, 5),
	}
}

type Wrapper[T any] struct {
	inner      []T
	start, end int
	funcChain  []any
	collected  []T
}

func (w *Wrapper[T]) Unwrap() []T {
	return w.inner
}

func (w *Wrapper[T]) Range(start, end int) iter.SliceIter[T] {
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

func (w *Wrapper[T]) Filter(filterFunc func(T) bool) iter.SliceIter[T] {
	w.funcChain = append(w.funcChain, filterFunc)
	return w
}

func (w *Wrapper[T]) Map(mapFunc func(*T)) iter.SliceIter[T] {
	w.funcChain = append(w.funcChain, mapFunc)
	return w
}

func (w *Wrapper[T]) ForEach(handle func(*T)) {
	for i := w.start; i < w.end; i++ {
		handle(&w.inner[i])
	}
}

func (w *Wrapper[T]) Find(qualified func(T) bool) (*T, error) {
	// defer func() {
	// 	pool.Put(w)
	// }()
	for i := w.start; i < w.end; i++ {
		if qualified(w.inner[i]) {
			q := w.inner[i]
			return &q, nil
		}
	}
	return nil, errors.New("not found")
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
