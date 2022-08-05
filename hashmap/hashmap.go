package hashmap

import (
	"github.com/sslime336/iter"
	"github.com/sslime336/iter/slice"
)

func Iter[K comparable, V any](m map[K]V) iter.HashMapIter[K, V] {
	return &wrapper[K, V]{
		inner:  m,
		volume: len(m),
	}
}

type wrapper[K comparable, V any] struct {
	inner  map[K]V
	volume int
	curKey K
	curVal V
}

// TODO: as this do not guarantee the order of the inner elements,
// keeping this method or not is under consideration.
func (w *wrapper[K, V]) Next() iter.HashMapIter[K, V] {
	return nil
}

func (w *wrapper[K, V]) Unwrap() (key K, val V) {
	key = w.curKey
	val = w.curVal
	return
}

func (w *wrapper[K, V]) Keys() iter.SliceIter[K] {
	keys := make([]K, 0, w.volume)
	for key := range w.inner {
		keys = append(keys, key)
	}
	return slice.Iter(keys)
}

func (w *wrapper[K, V]) Values() iter.SliceIter[V] {
	values := make([]V, 0, w.volume)
	for _, val := range w.inner {
		values = append(values, val)
	}
	return slice.Iter(values)
}
