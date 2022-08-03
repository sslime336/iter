package hashmap

import (
	"github.com/sslime336/iter"
	"github.com/sslime336/iter/slice"
)

func Iter[K comparable, V any](m map[K]V) iter.HashMapIter[K, V] {
	return &Wrapper[K, V]{
		inner:  m,
		volume: len(m),
	}
}

type Wrapper[K comparable, V any] struct {
	inner  map[K]V
	volume int
}

// TODO: whether to do or not
func (w *Wrapper[K, V]) Next() (key K, value V) {
	return
}

func (w *Wrapper[K, V]) Keys() iter.SliceIter[K] {
	keys := make([]K, 0, w.volume)
	for key := range w.inner {
		keys = append(keys, key)
	}
	return slice.Iter(keys)
}

func (w *Wrapper[K, V]) Values() iter.SliceIter[V] {
	values := make([]V, 0, w.volume)
	for _, val := range w.inner {
		values = append(values, val)
	}
	return slice.Iter(values)
}
