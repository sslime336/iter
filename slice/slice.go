package slice

import (
	"errors"

	"github.com/sslime336/iter"
	"github.com/sslime336/iter/logger"
)

func Iter[T any](sli []T) iter.SliceIter[T] {
	return &wrapper[T]{
		inner:     sli,
		start:     0,
		end:       len(sli),
		fixed:     len(sli),
		funcChain: make([]any, 0, 3),
		collected: sli,
	}
}

type wrapper[T any] struct {
	inner      []T
	fixed      int // inner's primitive size
	start, end int
	moved      bool
	funcChain  []any
	collected  []T
}

// Next increase the start index of the inner slice.
// Return true if any elements exists.
func (w *wrapper[T]) Next() (exists bool) {
	if w.start < w.end {
		if !w.moved {
			w.moved = true
		} else {
			w.start++
		}
		exists = true
	}
	return
}

// Element return the element on the index start,
// ok is false if there are no elements in range(start, end).
func (w *wrapper[T]) Element() (elem T, ok bool) {
	if w.start < w.end {
		elem = w.inner[w.start]
		ok = true
	}
	return
}

// Unwrap return the inner slice of the data
func (w *wrapper[T]) Unwrap() []T {
	return w.inner
}

func (w *wrapper[T]) Range(start, end int) iter.SliceIter[T] {
	defer func() {
		if p := recover(); p != nil {
			logger.Error(p)
		}
	}()
	if start < 0 || end < 0 || start > end || end > w.fixed {
		panic("illegal range")
	}
	w.start, w.end = start, end
	return w
}

// Filter will register a filterFunc, which is lazy.
func (w *wrapper[T]) Filter(filterFunc func(T) bool) iter.SliceIter[T] {
	w.funcChain = append(w.funcChain, filterFunc)
	return w
}

// Map will register a mapFunc, which is lazy.
func (w *wrapper[T]) Map(mapFunc func(*T)) iter.SliceIter[T] {
	w.funcChain = append(w.funcChain, mapFunc)
	return w
}

// ForEach will call handle to every element which is in the range(start <= i < end).
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

// TODO: consider use int64 uint64 float64 complex128 to handle
// of user self-build function to handle
// Should in another go file
func Sum[T operatable](slice []T) int64 {
	return int64(0)
}

type operatable interface {
	// TODO: consider
	~uint8 | ~int8 | ~uint16 | ~int16 | ~uint32 |
		~int32 | ~uint64 | ~int64 | ~uint | ~int |
		~float32 | ~float64 | ~complex64 | ~complex128
}

func exec_funcChain[T any](b *wrapper[T]) {
	for _, chainFunc := range b.funcChain {
		switch t := chainFunc.(type) {
		case /* map functions */ func(*T):
			for i := b.start; i < b.end; i++ {
				t(&b.collected[i])
			}
		case /* fliter functions */ func(T) bool:
			recollect := make([]T, 0, 5)
			collectN := 0
			for i := b.start; i < b.end; i++ {
				if t(b.collected[i]) {
					recollect = append(recollect, b.collected[i])
					collectN++
				}
			}
			copy(b.collected, recollect)
			b.collected = b.collected[:collectN]
			b.start, b.end = 0, collectN
		default:
			logger.Error("unmatched func type")
		}
	}
}
