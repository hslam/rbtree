// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package rbtree

import (
	"testing"
)

func TestRbtree(t *testing.T) {
	for i := 0; i < 100; i++ {
		testRbtree(100, i, true, t)
		testRbtree(100, i, false, t)
		testRbtreeM(100, i+1, true, t)
		testRbtreeM(100, i+1, false, t)
	}
}

func testRbtree(n, j int, r bool, t *testing.T) {
	tree := New()
	if r {
		for i := n - 1; i >= 0; i-- {
			tree.Insert(Int(i))
			testTraversal(tree, t)
		}
	} else {
		for i := 0; i < n; i++ {
			tree.Insert(Int(i))
			testTraversal(tree, t)
		}
	}
	if tree.Length() != n {
		t.Error("")
	}
	tree.Delete(Int(n))
	if tree.Length() != n {
		t.Error("")
	}
	testSearch(tree, j, t)
	tree.Delete(Int(j))
	testTraversal(tree, t)
	testNilNode(tree, j, t)
	if tree.Length() != n-1 {
		t.Error("")
	}
	if r {
		for i := n - 1; i >= 0; i-- {
			tree.Delete(Int(i))
			testTraversal(tree, t)
			testNilNode(tree, i, t)
		}
	} else {
		for i := 0; i < n; i++ {
			tree.Delete(Int(i))
			testTraversal(tree, t)
			testNilNode(tree, i, t)
		}
	}
	if tree.Length() != 0 {
		t.Error(tree.Length())
	}
}

func testRbtreeM(n, j int, r bool, t *testing.T) {
	tree := New()
	if r {
		for i := n; i > 0; i-- {
			tree.Insert(Int(i))
			testTraversal(tree, t)
			tree.Insert(Int(-i))
			testTraversal(tree, t)
		}
	} else {
		for i := 1; i < n+1; i++ {
			tree.Insert(Int(i))
			testTraversal(tree, t)
			tree.Insert(Int(-i))
			testTraversal(tree, t)
		}
	}
	if tree.Length() != n*2 {
		t.Error("")
	}
	testSearch(tree, j, t)
	tree.Delete(Int(j))
	testTraversal(tree, t)
	testNilNode(tree, j, t)
	if tree.Length() != n*2-1 {
		t.Error("")
	}
	j = -j
	testSearch(tree, j, t)
	tree.Delete(Int(j))
	testTraversal(tree, t)
	testNilNode(tree, j, t)
	if tree.Length() != n*2-2 {
		t.Error("")
	}
	if r {
		for i := n; i > 0; i-- {
			tree.Delete(Int(i))
			testTraversal(tree, t)
			testNilNode(tree, i, t)
			tree.Delete(Int(-i))
			testTraversal(tree, t)
			testNilNode(tree, -i, t)
		}
	} else {
		for i := 1; i < n+1; i++ {
			tree.Delete(Int(i))
			testTraversal(tree, t)
			testNilNode(tree, i, t)
			tree.Delete(Int(-i))
			testTraversal(tree, t)
			testNilNode(tree, -i, t)
		}
	}
	if tree.Length() != 0 {
		t.Error(tree.Length())
	}
}

func testTraversal(tree *Tree, t *testing.T) {
	count := 0
	testLength(tree.Root(), &count)
	if tree.Length() != count {
		t.Error(tree.Length(), count)
	}
	traverse(tree.Root(), t)
	testIteratorAscend(tree, t)
	testIteratorDescend(tree, t)
}

func testLength(node *Node, count *int) {
	if node.Item() != nil {
		*count++
	}
	if node != nil {
		testLength(node.Left(), count)
		testLength(node.Right(), count)
	}
}

func traverse(node *Node, t *testing.T) {
	if node.Color() == Red {
		if node.Left() != nil && node.Right() != nil {
			if node.left.Color() != Black || node.right.Color() != Black {
				t.Error("")
			}
		}
	}
	if node.Left() != nil && node.Left().parent != node {
		t.Error("")
	}
	if node.Right() != nil && node.Right().parent != node {
		t.Error("")
	}
	if node != nil {
		traverse(node.Left(), t)
		traverse(node.Right(), t)
	}
}

func testIteratorAscend(tree *Tree, t *testing.T) {
	iter := tree.Root().Min()
	next := iter.Next()
	for iter != nil && next != nil {
		if !iter.Item().Less(next.Item()) {
			t.Error("")
		}
		iter = next
		next = iter.Next()
	}
}

func testIteratorDescend(tree *Tree, t *testing.T) {
	iter := tree.Root().Max()
	last := iter.Last()
	for iter != nil && last != nil {
		if !last.Item().Less(iter.Item()) {
			t.Error("")
		}
		iter = last
		last = iter.Last()
	}
}

func testSearch(tree *Tree, j int, t *testing.T) {
	if node := tree.SearchNode(Int(j)); node == nil {
		t.Error("")
	} else if int(node.Item().(Int)) != j {
		t.Error("")
	}
	if item := tree.Search(Int(j)); item == nil {
		t.Error("")
	} else if int(item.(Int)) != j {
		t.Error("")
	}
}

func testNilNode(tree *Tree, j int, t *testing.T) {
	if item := tree.Search(Int(j)); item != nil {
		t.Error("")
	}
}

func TestInsertCase1(t *testing.T) {
	tree := New()
	tree.Insert(Int(13))
	testTraversal(tree, t)
}

func TestInsertCase2(t *testing.T) {
	tree := New()
	tree.Insert(Int(13))
	tree.Insert(Int(8))
	testTraversal(tree, t)
	tree.Clear()
	tree.Insert(Int(13))
	tree.Insert(Int(17))
	testTraversal(tree, t)
}

func TestInsertCase3(t *testing.T) {
	tree := New()
	tree.Insert(Int(13))
	tree.Insert(Int(8))
	tree.Insert(Int(17))
	tree.Insert(Int(1))
	testTraversal(tree, t)
	tree.Clear()
	tree.Insert(Int(13))
	tree.Insert(Int(8))
	tree.Insert(Int(17))
	tree.Insert(Int(11))
	testTraversal(tree, t)
	tree.Clear()
	tree.Insert(Int(13))
	tree.Insert(Int(8))
	tree.Insert(Int(17))
	tree.Insert(Int(15))
	testTraversal(tree, t)
	tree.Clear()
	tree.Insert(Int(13))
	tree.Insert(Int(8))
	tree.Insert(Int(17))
	tree.Insert(Int(25))
	testTraversal(tree, t)
}

func TestInsertCase4(t *testing.T) {
	tree := New()
	l := []Int{13, 8, 17, 10, 15, 25}
	for _, v := range l {
		tree.Insert(v)
		testTraversal(tree, t)
	}
	tree.Insert(Int(9))
	testTraversal(tree, t)
}

func TestEmptyTree(t *testing.T) {
	tree := New()
	if tree.Root() != nil {
		t.Error("")
	}
	if tree.Min() != nil {
		t.Error("")
	}
	if tree.Max() != nil {
		t.Error("")
	}
	if tree.Min() != nil {
		t.Error("")
	}
	if tree.Root().Left() != nil {
		t.Error("")
	}
	if tree.Root().Right() != nil {
		t.Error("")
	}
	if tree.Root().Item() != nil {
		t.Error("")
	}
	if tree.Root().Parent() != nil {
		t.Error("")
	}
	if tree.Root().GrandParent() != nil {
		t.Error("")
	}
	if tree.Root().Sibling() != nil {
		t.Error("")
	}
	if tree.Root().Last() != nil {
		t.Error("")
	}
	if tree.Root().Next() != nil {
		t.Error("")
	}
	tree.Insert(Int(0))
	if tree.Root().GrandParent() != nil {
		t.Error("")
	}
	if tree.Root().Sibling() != nil {
		t.Error("")
	}
	tree.Clear()
	tree.Delete(Int(0))
	if tree.Length() != 0 {
		t.Error("")
	}
}

func BenchmarkRBTree(b *testing.B) {
	tree := New()
	for i := 0; i < b.N; i++ {
		tree.Insert(Int(i))
		//tree.Search(Int(i))
		//tree.Delete(Int(i))
	}
}

func TestStringLess(t *testing.T) {
	a := String("a")
	b := String("b")
	if !a.Less(b) {
		t.Error("")
	}
}

func TestReplaceItem(t *testing.T) {
	tree := New()
	n := 1024
	for i := 0; i < n; i++ {
		tree.Insert(Int(i))
		testTraversal(tree, t)
		if tree.Length() != i+1 {
			t.Error("")
		}
	}
	for i := 0; i < n; i++ {
		tree.Insert(Int(i))
		testTraversal(tree, t)
		if tree.Length() != n {
			t.Error("")
		}
	}
	for i := 0; i < n; i++ {
		tree.Delete(Int(i))
		testTraversal(tree, t)
		if tree.Length() != n-i-1 {
			t.Error("")
		}
	}
}
