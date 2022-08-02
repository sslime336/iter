package slice

import (
	"errors"

	"github.com/sslime336/iter"
	"github.com/sslime336/iter/logger"
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
			logger.Error(p)
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
	for i := w.start; i < w.end; i++ {
		if qualified(w.inner[i]) {
			q := w.inner[i]
			return &q, nil
		}
	}
	return nil, errors.New("not found")
}

func (w *Wrapper[T]) Count() int {
	exec_funcChain(w)
	return len(w.collected)
}

// ZipPtr combines two slices, keys and vals.
// The type of vals will be convert into the pointer of its type as
// it's easy to achieve, but is not appropriate.
func ZipPtr[K comparable, V any](keys []K, vals []V) (map[K]*V, error) {
	klen, vlen := len(keys), len(vals)
	if vlen > klen {
		return nil, errors.New("values is more than keys")
	} else if vlen < klen {
		logger.Warn("the number of keys and values is not matched, keys are more")
	}
	hp := make(map[K]*V, klen)
	containV := append([]V(nil), vals...)
	for i := 0; i < klen; i++ {
		if i < vlen {
			hp[keys[i]] = &containV[i]
		} else {
			hp[keys[i]] = nil
		}
	}
	return hp, nil
}

// Zip combines two slices, keys and vals.
// Diffs from ZipPtr, it will do what it looks like.
func Zip[K comparable, V any](keys []K, vals []V) (map[K]V, error) {
	klen, vlen := len(keys), len(vals)
	if vlen != klen {
		return nil, errors.New("the number of A and B are not equal")
	}
	hp := make(map[K]V, klen)
	containV := append([]V(nil), vals...)
	for i := 0; i < klen; i++ {
		hp[keys[i]] = containV[i]
	}
	return hp, nil
}

// Collect will return the current slice, which has been copied from
// the Wrapper's inner(that has been dealed with chainFuncs).
func (w *Wrapper[T]) Collect() []T {
	exec_funcChain(w)
	collected := make([]T, len(w.collected))
	copy(collected, w.collected)
	return collected
}

// TODO: whether save Sum or not
func Sum[T operatable](slice []T) int64 {
	return int64(0)
}

type operatable interface {
	// TODO: finish this
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
			logger.Error("unmatched func type")
		}
	}
}
