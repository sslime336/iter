package hashmap

import (
	"github.com/sslime336/iter"
)

func Iter[K comparable, V any](m map[K]V) iter.HashMapIter[K, V] {
	return &wrapper[K, V]{
		inner:  m,
		volume: len(m),
	}
}
