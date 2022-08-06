package channel

type wrapper[T any] struct {
	inner chan T
}

func (w *wrapper[T]) Next() (exists bool) {
	return
}

func (w *wrapper[T]) _()  {
	
}
