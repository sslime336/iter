package iter

import "github.com/sslime336/iter/slice"

type Iterator[T any] interface {
	// SliceIter[T]
	// ArrayIter[T]
}

type SliceIter[T any] interface {
	Range(int, int) *slice.Wrapper[T]
	Filter(func(T) bool) *slice.Wrapper[T]
	Map(func(*T)) *slice.Wrapper[T]
	ForEach(func(*T))
	Find(func(T) bool) T
	Count()
	Zip()
	Collect() []T
	Sum()
}

type ArrayIter[T any] interface {
	Range(int, int)
	Filter(func(T) bool)
	Map(func(*T))
	ForEach(func(*T))
	Find(func(T) bool)
	Count()
	Zip()
	Collect()
	Sum()
}

type HashMapIter[K comparable, V any] interface {
	Count()
	Collect()
}
