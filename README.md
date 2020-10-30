# rbtree
Package rbtree implements a red–black tree.

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
```
package main

import (
	"fmt"
	"github.com/hslam/rbtree"
)

func main() {
	t := rbtree.New()
	str := item("Hello World")
	t.Insert(str)
	fmt.Println(t.Search(str))
	t.Delete(str)
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

### License
This package is licensed under a MIT license (Copyright (c) 2020 Meng Huang)

### Author
rbtree was written by Meng Huang.


