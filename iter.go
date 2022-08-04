package iter

type Iterator[T any] interface {
	// SliceIter[T]
	// HashMapIter[K, V]
}

type SliceIter[T any] interface {
	Next() (*T, bool)
	Range(int, int) SliceIter[T]
	Filter(func(T) bool) SliceIter[T]
	Map(func(*T)) SliceIter[T]
	ForEach(func(*T))
	Find(func(T) bool) (*T, bool)
	Count() int
	Collect() []T
	Unwrap() []T
}

// TODO: save or not
type HashMapIter[K comparable, V any] interface {
	Keys() SliceIter[K]
	Values() SliceIter[V]
}
