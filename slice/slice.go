package slice

import (
	"errors"

	"github.com/sslime336/iter"
	"github.com/sslime336/iter/logger"
)

func Iter[T any](sli []T) iter.SliceIter[T] {
	return &wrapper[T]{
		inner:      sli,
		start:      0,
		end:        len(sli),
		funcChain:  make([]any, 3),
		emptyChain: true,
		collected:  make([]T, 0, 5),
	}
}

type wrapper[T any] struct {
	inner      []T
	start, end int
	funcChain  []any
	emptyChain bool
	collected  []T
}

// Next is useless now?
func (w *wrapper[T]) Next() (*T, bool) {
	if w.start < w.end {
		w.start++
		return &w.inner[w.start], true
	}
	return nil, false
}

func (w *wrapper[T]) Unwrap() []T {
	return w.inner
}

func (w *wrapper[T]) Range(start, end int) iter.SliceIter[T] {
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

func (w *wrapper[T]) Filter(filterFunc func(T) bool) iter.SliceIter[T] {
	// As the field emptyChain is simple, doing this is litte faster,
	// though logically unsuitable.
	w.emptyChain = false
	w.funcChain = append(w.funcChain, filterFunc)
	return w
}

func (w *wrapper[T]) Map(mapFunc func(*T)) iter.SliceIter[T] {
	w.emptyChain = false
	w.funcChain = append(w.funcChain, mapFunc)
	return w
}

func (w *wrapper[T]) ForEach(handle func(*T)) {
	exec_funcChain(w)
	for i := w.start; i < w.end; i++ {
		handle(&w.inner[i])
	}
}

// FindPtr will return the pointer of the found value.
// Return nil and false if not found.
func (w *wrapper[T]) FindPtr(qualified func(T) bool) (*T, bool) {
	exec_funcChain(w)
	for i := w.start; i < w.end; i++ {
		if qualified(w.inner[i]) {
			return &w.inner[i], true
		}
	}
	return nil, false
}

// Find will return the copy of the T, if found, the param `found`
// will be true
func (w *wrapper[T]) Find(qualified func(T) bool) (res T, found bool) {
	exec_funcChain(w)
	for i := w.start; i < w.end; i++ {
		if qualified(w.inner[i]) {
			res, found = w.inner[i], true
			return
		}
	}
	return
}

func (w *wrapper[T]) Count() int {
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
// the wrapper's inner(that has been dealed with chainFuncs).
func (w *wrapper[T]) Collect() []T {
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
	~uint8 | ~int8 | ~uint16 | ~int16 | ~uint32 |
		~int32 | ~uint64 | ~int64 | ~uint | ~int |
		~float32 | ~float64 | ~complex64 | ~complex128
}

func exec_funcChain[T any](b *wrapper[T]) {
	if b.emptyChain {
		b.collected = append(b.collected, b.inner...)
		return
	}
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
