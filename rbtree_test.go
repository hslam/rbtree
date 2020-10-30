package rbtree

import (
	"testing"
)

func TestRbtree(t *testing.T) {
	for i := 0; i < 100; i++ {
		testRbtree(100, i, true, t)
		testRbtree(100, i, false, t)
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
	testSearch(tree, j, t)
	tree.Delete(Int(j))
	testTraversal(tree, t)
	testNilNode(tree, j, t)
	if tree.Length() != n-1 {
		t.Error("")
	}
}

func testTraversal(tree *Tree, t *testing.T) {
	traverse(tree.Root(), t)
	testIteratorAscend(tree, t)
	testIteratorDescend(tree, t)
}

func traverse(node *Node, t *testing.T) {
	if node.Color() == Red {
		if node.Left() != nil && node.Right() != nil {
			if node.left.Color() != Black || node.right.Color() != Black {
				t.Error("")
			}
		}
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
	if node := tree.Search(Int(j)); node != nil {
		t.Error("")
	}
}

func TestEmptyTree(t *testing.T) {
	tree := New()
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
