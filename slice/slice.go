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

// Sum return the sum of the given slice.
func Sum[T operatable](slice []T) (sum float64) {
	for i := 0; i < len(slice); i++ {
		sum += float64(slice[i])
	}
	return
}

// Sum2 return the sum of the given slice, which elements are
// imaginary number.
func Sum2[T operatable2](slice []T) (sum complex128) {
	for i := 0; i < len(slice); i++ {
		sum += complex128(slice[i])
	}
	return
}

type operatable interface {
	~uint | ~int |
		~uint8 | ~int8 |
		~uint16 | ~int16 |
		~uint32 | ~int32 |
		~uint64 | ~int64 |
		~float32 | ~float64
}

type operatable2 interface {
	~complex64 | ~complex128
}
