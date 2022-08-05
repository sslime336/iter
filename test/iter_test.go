package test

import (
	"fmt"
	"testing"

	"github.com/sslime336/iter"
	"github.com/sslime336/iter/slice"
)

type ints struct {
	inner []int
}

func (i *ints) Iter() iter.SliceIter[int] {
	return slice.Iter(i.inner)
}

func TestIter(t *testing.T) {
	intvar := ints{inner: []int{1, 2, 3, 4, 5}}
	fmt.Println(intvar.Iter().Count())
}
