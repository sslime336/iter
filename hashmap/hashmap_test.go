package hashmap

import (
	"log"
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
	log.Println(Iter(maps).Keys().Unwrap())
	if found, err := Iter(maps).Values().Find(func(s string) bool {
		return s == "0x3"
	}); err == nil {
		log.Println(*found)
	}
}
