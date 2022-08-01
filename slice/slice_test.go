package slice

import (
	"fmt"
	"sync"
	"testing"
)

var _pool = sync.Pool{
	New: func() any {
		return new(Wrapper[any])
	},
}

func TestWrapper(t *testing.T) {
	s, ok := _pool.Get().(*Wrapper[bool])
	if !ok {
		fmt.Println("not ok")
	} else {
		fmt.Println("ok")
	}
	_pool.Put(s)

	s2, ok2 := _pool.Get().(*Wrapper[int])
	if !ok2 {
		fmt.Println("not ok2")
	} else {
		fmt.Println("ok2")
	}
	_pool.Put(s2)
}

func TestCopySlice(t *testing.T) {
	src := []int{1, 2, 3, 4}
	// var s1 []int
	s1 := make([]int, len(src))
	// var s2 = new([]int)
	copy(s1, src)
	fmt.Printf("s1: %v\n", s1)
	// copy(s2, src)
}
