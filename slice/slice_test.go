package slice

import (
	"log"
	"testing"
)

func TestForEach(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	Iter(nums).ForEach(func(t *int) {
		*t += 100
	})
	log.Println("ints:", nums)

	// with named struct
	type user struct {
		name string
		age  int
	}
	users := []user{
		{
			name: "Alice",
			age:  17,
		},
		{
			name: "Jack",
			age:  19,
		},
	}
	Iter(users).ForEach(func(t *user) {
		t.name = "NONE"
		t.age = -1
	})
	log.Println("users:", users)

	// with anonymous struct
	usersNoName := []struct {
		name string
		age  int
	}{
		{
			name: "Alice",
			age:  17,
		},
		{
			name: "Jack",
			age:  19,
		},
	}
	Iter(usersNoName).ForEach(func(t *struct {
		name string
		age  int
	}) {
		t.name = "NONE"
		t.age = -1
	})
	log.Println("anaymous users:", usersNoName)
}

func TestZip(t *testing.T) {
	keys := []int{1, 2, 3, 4, 5}
	vals := []string{"Apple", "Banana", "Orange"}
	vals2 := []string{"Apple", "Banana", "Orange", "Peach", "QwQ"}
	res1, e := Zip(keys, vals)
	log.Println("error:", e)
	res2, _ := Zip(keys, vals2)
	log.Printf("zip: %+v", res1)
	log.Printf("zip2: %+v", res2)
}

func TestFind(t *testing.T) {

}

func TestCollect(t *testing.T) {

}