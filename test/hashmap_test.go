package test

import (
	"fmt"
	"testing"

	"github.com/sslime336/iter/hashmap"
)

var grades = map[string]string{
	"Alice": "A",
	"Bob":   "B+",
	"Jack":  "B-",
	"Foo":   "C",
}

func TestNext(t *testing.T) {
	hpIter := hashmap.Iter(grades)
	for hpIter.Next() {
		fmt.Println(hpIter.Element())
	}
}
