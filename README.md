# rbtree
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hslam/rbtree)](https://pkg.go.dev/github.com/hslam/rbtree)
[![Build Status](https://travis-ci.org/hslam/rbtree.svg?branch=master)](https://travis-ci.org/hslam/rbtree)
[![codecov](https://codecov.io/gh/hslam/rbtree/branch/master/graph/badge.svg)](https://codecov.io/gh/hslam/rbtree)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/rbtree)](https://goreportcard.com/report/github.com/hslam/rbtree)
[![LICENSE](https://img.shields.io/github/license/hslam/rbtree.svg?style=flat-square)](https://github.com/hslam/rbtree/blob/master/LICENSE)

Package rbtree implements a red–black tree.

**[Properties](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree "properties")**
* Each node is either red or black.
* The root is black. This rule is sometimes omitted. Since the root can always be changed from red to black, but not necessarily vice versa, this rule has little effect on analysis.
* All leaves (NIL) are black.
* If a node is red, then both its children are black.
* Every path from a given node to any of its descendant NIL nodes goes through the same number of black nodes.

## Get started

### Install
```
go get github.com/hslam/rbtree
```
### Import
```
import "github.com/hslam/rbtree"
```
### Usage
#### Example
```go
package main

import (
	"fmt"
	"github.com/hslam/rbtree"
)

func main() {
	tree := rbtree.New()
	str := item("Hello World")
	tree.Insert(str)
	fmt.Println(tree.Search(str))
	tree.Delete(str)
}

type item string

func (n item) Less(b rbtree.Item) bool {
	return n < b.(item)
}
```

#### Output
```
Hello World
```

#### Iterator Example
```go
package main

import (
	"fmt"
	"github.com/hslam/rbtree"
)

func main() {
	tree := rbtree.New()
	l := []Int{13, 8, 17, 1, 11, 15, 25, 6, 22, 27}
	for _, v := range l {
		tree.Insert(v)
	}
	iter := tree.Min()
	for iter != nil {
		fmt.Printf("%d\t", iter.Item())
		iter = iter.Next()
	}
}

type Int int

func (a Int) Less(b rbtree.Item) bool {
	return a < b.(Int)
}
```
#### Red–Black Tree
<img src="https://raw.githubusercontent.com/hslam/rbtree/master/rbtree.png" alt="rbtree" align=center>

#### Output
```
1	6	8	11	13	15	17	22	25	27
```


### License
This package is licensed under a MIT license (Copyright (c) 2020 Meng Huang)

### Author
rbtree was written by Meng Huang.


