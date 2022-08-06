package hashmap

import (
	"github.com/sslime336/iter"
)

func Iter[K comparable, V any](m map[K]V) iter.HashMapIter[K, V] {
	volume := len(m)
	keys := collectKeys(m, volume)
	vals := collectVals(m, volume)
	return &wrapper[K, V]{
		inner:  m,
		volume: volume,
		keys:   keys,
		vals:   vals,
		idx:    -1,
	}
}

func collectKeys[K comparable, V any](hp map[K]V, siz int) []K {
	keys := make([]K, 0, siz)
	for key := range hp {
		keys = append(keys, key)
	}
	return keys
}

func collectVals[K comparable, V any](hp map[K]V, siz int) []V {
	vals := make([]V, 0, siz)
	for _, val := range hp {
		vals = append(vals, val)
	}
	return vals
}
