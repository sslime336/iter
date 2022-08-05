# iter

[English](README.md) | 简体中文

### 是什么

使用 Go 实现的实验性质的迭代器，以支持函数式编程

### 警告

所有接口都是试验性的，且不应该运用在生产环境中

### 例子

这里提供了两种获取 Iterator 对象的方式
1. 从已经存在的切片、数组或 map 中生成
2. 给结构体实现相应的迭代器接口。使用已经提供的函数，可以简化迭代器接口的实现

[1] 从已经存在的切片、数组或 map 中生成

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


[2] 实现 Iterator / Iterator2 接口

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

### Logger

模块提供了一个 WeakLogger 接口，以在迭代过程中使用你喜欢的日志框架格式化日志信息

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

### 杂项

持续更新ing...(大概)
