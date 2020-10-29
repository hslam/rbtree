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
			traverse(tree.Root(), t)
		}
	} else {
		for i := 0; i < n; i++ {
			tree.Insert(Int(i))
			traverse(tree.Root(), t)
		}
	}
	if tree.Length() != n {
		t.Error("")
	}
	testSearch(tree, j, t)
	tree.Delete(Int(j))
	traverse(tree.Root(), t)
	testNilNode(tree, j, t)
	if tree.Length() != n-1 {
		t.Error("")
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
	if node != nil {
		traverse(node.Left(), t)
		traverse(node.Right(), t)
	}
}

func testSearch(tree *Tree, j int, t *testing.T) {
	if node := tree.Search(Int(j)); node == nil {
		t.Error("")
	} else if int(node.Item().(Int)) != j {
		t.Error("")
	}
}
func testNilNode(tree *Tree, j int, t *testing.T) {
	if node := tree.Search(Int(j)); node != nil {
		t.Error("")
	}
}
