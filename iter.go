package iter

type Iterator[T any] interface {
	// SliceIter[T]
	// HashMapIter[K, V]
}

type SliceIter[T any] interface {
	Range(int, int) SliceIter[T]
	Filter(func(T) bool) SliceIter[T]
	Map(func(*T)) SliceIter[T]
	ForEach(func(*T))
	Find(func(T) bool) (*T, error)
	Count() int
	Collect() []T
	Unwrap() []T
}

// ArrayIter was deleted, as array can use arr[:] to use.
// Like slice.Iter(arr[:]).

// TODO: save or not
type HashMapIter[K comparable, V any] interface {
	Keys() SliceIter[K]
	Values() SliceIter[V]
}
