package iter

type Iterator[T any] interface {
	Iter() SliceIter[T]
}

type Iterator2[K comparable, V any] interface {
	Iter() HashMapIter[K, V]
}

type Iterator3[T any] interface {
	Iter() ChannelIter[T]
}

type SliceIter[T any] interface {
	Next() bool
	Element() (T, bool)
	Range(int, int) SliceIter[T]
	Filter(func(T) bool) SliceIter[T]
	Map(func(*T)) SliceIter[T]
	ForEach(func(*T))
	Find(func(T) bool) (T, bool)
	FindPtr(func(T) bool) (*T, bool)
	Count() int
	Collect() []T
	Unwrap() []T
}

type HashMapIter[K comparable, V any] interface {
	Keys() SliceIter[K]
	Values() SliceIter[V]
	Next() bool
	Element() (K, V, bool)
}

type ChannelIter[T any] interface {
}
