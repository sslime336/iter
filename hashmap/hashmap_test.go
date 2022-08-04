package hashmap

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	maps := map[int]string{
		1: "0x1",
		2: "0x2",
		3: "0x3",
		4: "0x4",
		5: "0x5",
	}
	fmt.Println(Iter(maps).Keys().Unwrap())
	if found, exists := Iter(maps).Values().Find(func(s string) bool {
		return s == "0x3"
	}); exists {
		fmt.Println(*found)
	}
}

func TestMapSlice(t *testing.T) {
	m := map[string]string{
		"name":    "Alice",
		"age":     "17",
		"gender":  "woman",
		"address": "unknow",
		"phone":   "18111101111",
	}
	fmt.Println(Iter(m).Keys().Collect())
}
