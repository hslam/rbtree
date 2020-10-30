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

#### Iterator Example
```
package main

import (
	"fmt"
	"github.com/hslam/rbtree"
)

func main() {
	t := rbtree.New()
	l := []Int{13, 8, 17, 1, 11, 15, 25, 6, 22, 27}
	for _, v := range l {
		t.Insert(v)
	}
	iter := t.Root().Min()
	for iter != nil {
		fmt.Printf("%d\t", iter.Item())
		iter = iter.Next()
	}
}

type Int int

func (a Int) Less(than rbtree.Item) bool {
	b, _ := than.(Int)
	return a < b
}
```
#### Tree
<img src="https://raw.githubusercontent.com/hslam/rbtree/master/tree.png" alt="tree" align=center>

#### Output
```
1	6	8	11	13	15	17	22	25	27
```


### License
This package is licensed under a MIT license (Copyright (c) 2020 Meng Huang)

### Author
rbtree was written by Meng Huang.


