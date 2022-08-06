package test

import (
	"fmt"
	"testing"

	"github.com/sslime336/iter"
	"github.com/sslime336/iter/slice"
)

var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var numsArr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

type user struct {
	id   uint64
	name string
	age  uint
	addr string
}

var users = []user{
	{
		id:   422526,
		name: "Alice",
		age:  23,
		addr: "nowhere",
	},
	{
		id:   655623,
		name: "Jack",
		age:  32,
		addr: "nowhere",
	},
	{
		id:   555123,
		name: "Bob",
		age:  17,
		addr: "another",
	},
	{
		id:   751298,
		name: "Foo",
		age:  43,
		addr: "earth",
	},
}

func TestSliceAndArray(t *testing.T) {
	var _ iter.SliceIter[int] = slice.Iter(nums)
	var _ iter.SliceIter[int] = slice.Iter(numsArr[:])
}

func TestCount(t *testing.T) {
	cnt := slice.Iter(nums).
		Filter(func(i int) bool {
			return i%2 == 0
		}).Count()
	fmt.Println(cnt)
}

func TestMap(t *testing.T) {
	slice.Iter(users).Map(func(t *user) {
		t.addr = ""
		t.age = 0
		t.id = 0
		t.name = ""
	}).ForEach(func(t *user) {
		fmt.Println(*t)
	})
}

func TestNextAndElement(t *testing.T) {
	numsIter := slice.Iter(nums)
	for numsIter.Next() {
		fmt.Print(numsIter.Element())
	}

	usersIter := slice.Iter(users)
	for usersIter.Next() {
		if elem, ok := usersIter.Element(); ok {
			fmt.Println(elem)
		}
	}
}

func TestRange(t *testing.T) {
	userIter := slice.Iter(users)
	// out of range
	userIter.Range(3, 5)
	// negative
	userIter.Range(-1, 2)
	// start > end
	userIter.Range(3, 2)
	userIter.Range(-2, -1)
}

func TestRangeAndForEach(t *testing.T) {
	numsIter := slice.Iter(nums)
	numsIter.Range(3, 6).ForEach(func(t *int) {
		*t += 10
	})
	fmt.Println(numsIter.Unwrap())
	fmt.Println(nums)

	usersIter := slice.Iter(users)
	usersIter.Range(2, 4).ForEach(func(t *user) {
		t.id = 0
		t.name = "NONE"
		t.age = 0
		t.addr = ""
	})
	fmt.Println(users)
}

func TestCollect(t *testing.T) {
	numsIter := slice.Iter(nums)
	numbers := numsIter.Collect()
	fmt.Println(numbers)

	nums0 := numsIter.Filter(func(i int) bool {
		return i%2 == 0
	}).Collect()
	fmt.Println(nums0)
}

func TestFuncChain(t *testing.T) {
	usersIter := slice.Iter(users)
	users := usersIter.
		Filter(func(u user) bool {
			return len(u.name) < 4
		}).
		Filter(func(u user) bool {
			return u.age < 30
		}).Collect()
	slice.Iter(users).ForEach(func(t *user) {
		fmt.Printf("%+v\n", *t)
	})
}

func TestFind(t *testing.T) {
	if num, exists := slice.Iter(nums).
		Filter(func(i int) bool {
			return i%2 == 0
		}).Find(func(i int) bool {
		return i == 8
	}); exists {
		fmt.Println(num)
	}
}

func TestSum(t *testing.T) {
	fmt.Println(slice.Sum(nums))
}

var imageNums = []complex64{complex(1, 0), complex(-1, 3), complex(1, -1), complex(0, 3)}

func TestSum2(t *testing.T) {
	fmt.Println(slice.Sum2(imageNums))
}

func TestMutiChainFunc(t *testing.T) {
	usersIter := slice.Iter(users)
	users := usersIter.
		Filter(func(u user) bool {
			return len(u.name) < 4
		}).
		Map(func(t *user) {
			t.name = "###NONAME"
		}).
		Filter(func(u user) bool {
			return u.age < 30
		}).Collect()
	slice.Iter(users).ForEach(func(t *user) {
		fmt.Printf("%+v\n", *t)
	})
}
