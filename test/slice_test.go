package test

import (
	"testing"

	"github.com/sslime336/iter/slice"
)

var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
var numsArr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

func TestSliceAndArray(t *testing.T) {
	slice.Iter(nums)
	slice.Iter(numsArr[:])
}
