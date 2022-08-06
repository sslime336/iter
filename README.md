# iter

English | [简体中文](README_cn.md)

### About

Experimental implementation of iterators using Go for **functional programming**.

### Warn

Since all of these are experimental, this module should not be used in productive environment.

### Examples

#### There are 2 ways to get the iterator.
1. Generate from existed slice, array or map.
2. Implement the Iterator interface. You can easily implement the interface by using the function Iter(any) from relevant package.


[1] Generate from existed slice, array or map.

```go
package main

import (
	"github.com/sslime336/iter"
	"github.com/sslime336/iter/hashmap"
	"github.com/sslime336/iter/slice"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	intarray := [...]int{1, 2, 3, 4, 5}
	grades := map[string]string{
		"Alice": "A",
		"Bob":   "B+",
		"Jack":  "B-",
		"Foo":   "C",
	}
	// from slice
	var _ iter.SliceIter[int] = slice.Iter(nums)
	// from array
	var _ iter.SliceIter[int] = slice.Iter(intarray[:])
	// from map
	var _ iter.HashMapIter[string, string] = hashmap.Iter(grades)
}

```


[2] Implement the Iterator / Iterator2 interface.

```go
package main

import (
	"github.com/sslime336/iter"
	"github.com/sslime336/iter/hashmap"
	"github.com/sslime336/iter/slice"
)

type numbers struct {
	inner []int
}

func (s *numbers) Iter() iter.SliceIter[int] {
	return slice.Iter(s.inner)
}

type grades struct {
	inner map[string]string
}

func (s *grades) Iter() iter.HashMapIter[string, string] {
	return hashmap.Iter(s.inner)
}

func main() {
	numberIter := numbers{inner: []int{1, 2, 3, 4, 5}}
	var _ iter.SliceIter[int] = numberIter.Iter()

	gradesIter := grades{inner: map[string]string{
		"Alice": "A",
		"Bob":   "B+",
		"Foo":   "C",
	}}
	var _ iter.HashMapIter[string, string] = gradesIter.Iter()
}

```

#### You can also see some examples over [here](./examples/demo.go).

### Logger

A WeakLogger interface is provided for inner use, which means you can customize the log output using your favorite logger.

```go
package main

import (
	"github.com/sirupsen/logrus"
	"github.com/sslime336/iter/logger"
)

func main() {
	logger.SetLogger(logrus.StandardLogger())
}

```

### Misc

Keep updating(maybe)...
