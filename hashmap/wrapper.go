package hashmap

import (
	"github.com/sslime336/iter"
	"github.com/sslime336/iter/logger"
	"github.com/sslime336/iter/slice"
)

type wrapper[K comparable, V any] struct {
	inner  map[K]V
	volume int
	keys   []K
	vals   []V
	idx    int
	curKey K
	curVal V
}

func (w *wrapper[K, V]) Next() (exists bool) {
	defer func() {
		// when idx overflow
		if p := recover(); p != nil {
			logger.Error(p)
		}
	}()
	w.idx++
	if w.idx < w.volume {
		exists = true
		w.curKey = w.keys[w.idx]
		w.curVal = w.vals[w.idx]
	}
	return
}

func (w *wrapper[K, V]) Element() (key K, value V, ok bool) {
	if w.idx < w.volume {
		key, value = w.curKey, w.curVal
		ok = true
	}
	return
}

func (w *wrapper[K, V]) Unwrap() map[K]V {
	return w.inner
}

func (w *wrapper[K, V]) Keys() iter.SliceIter[K] {
	return slice.Iter(w.keys)
}

func (w *wrapper[K, V]) Values() iter.SliceIter[V] {
	return slice.Iter(w.vals)
}
