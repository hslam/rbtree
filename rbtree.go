// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package rbtree implements a red–black tree.
//
// Red–black tree properties: https://en.wikipedia.org/wiki/Red%E2%80%93black_tree
//
// 1. Each node is either red or black.
//
// 2. The root is black. This rule is sometimes omitted. Since the root can always be changed from red to black, but not necessarily vice versa, this rule has little effect on analysis.
//
// 3. All leaves (NIL) are black.
//
// 4. If a node is red, then both its children are black.
//
// 5. Every path from a given node to any of its descendant NIL nodes goes through the same number of black nodes.
//
package rbtree

// Color is node color. Each node is either red or black.
type Color uint8

const (
	// Black represents the black color.
	Black Color = 0
	// Red represents the red color.
	Red Color = 1
)

// Item represents a value in the tree.
type Item interface {
	// Less compares whether the current item is less than the given Item.
	Less(than Item) bool
}

// Int implements the Item interface for int.
type Int int

// Less returns true if int(a) < int(b).
func (a Int) Less(b Item) bool {
	return a < b.(Int)
}

// String implements the Item interface for string.
type String string

// Less returns true if string(a) < string(b).
func (a String) Less(b Item) bool {
	return a < b.(String)
}

// Tree represents a red–black tree.
type Tree struct {
	length int
	root   *Node
}

// New returns a new red–black tree.
func New() *Tree {
	return &Tree{}
}

// Length returns the number of items currently in the red–black tree.
func (t *Tree) Length() int {
	return t.length
}

// Root returns the root node of the red–black tree.
func (t *Tree) Root() *Node {
	return t.root
}

// Max returns the max node of the red–black tree.
func (t *Tree) Max() *Node {
	return t.root.Max()
}

// Min returns the min node of the red–black tree.
func (t *Tree) Min() *Node {
	return t.root.Min()
}

// Search searches the item of the red–black tree.
func (t *Tree) Search(item Item) Item {
	return t.search(item).Item()
}

// SearchNode searches the node of the red–black tree with the item.
func (t *Tree) SearchNode(item Item) *Node {
	return t.search(item)
}

// Insert inserts the item into the red–black tree.
func (t *Tree) Insert(item Item) {
	n := t.root
	var p *Node
	for n != nil {
		p = n
		if item.Less(n.item) {
			n = n.left
		} else if n.item.Less(item) {
			n = n.right
		} else {
			n.item = item
			return
		}
	}
	n = &Node{parent: p, color: Red, item: item}
	t.length++
	if p == nil {
		t.root = n
	} else if n.item.Less(p.item) {
		p.left = n
	} else {
		p.right = n
	}
	t.insertRepairTree(n)
}

func (t *Tree) insertRepairTree(n *Node) {
	if n.parent == nil {
		t.insertCase1(n)
	} else if n.parent.color == Black {
		t.insertCase2(n)
	} else if n.Uncle() != nil && n.Uncle().color == Red {
		t.insertCase3(n)
	} else {
		t.insertCase4(n)
	}
}

func (t *Tree) insertCase1(n *Node) {
	n.color = Black
}

func (t *Tree) insertCase2(n *Node) {
}

func (t *Tree) insertCase3(n *Node) {
	n.Parent().color = Black
	n.Uncle().color = Black
	n.GrandParent().color = Red
	t.insertRepairTree(n.GrandParent())
}

func (t *Tree) insertCase4(n *Node) {
	p := n.Parent()
	g := n.GrandParent()
	if n == p.right && p == g.left {
		t.rotateLeft(p)
		n = n.left
	} else if n == p.left && p == g.right {
		t.rotateRight(p)
		n = n.right
	}
	t.insertCase4Step2(n)
}

func (t *Tree) insertCase4Step2(n *Node) {
	p := n.Parent()
	g := n.GrandParent()
	if n == p.left {
		t.rotateRight(g)
	} else {
		t.rotateLeft(g)
	}
	p.color = Black
	g.color = Red
}

// Clear removes all items from the red–black tree.
func (t *Tree) Clear() {
	t.root = nil
	t.length = 0
}

// Delete deletes the node of the red–black tree with the item.
func (t *Tree) Delete(item Item) {
	n := t.search(item)
	if n == nil {
		return
	}
	if n.left != nil && n.right != nil {
		min := n.right.Min()
		n.item = min.item
		t.deleteOneChild(min)
	} else {
		t.deleteOneChild(n)
	}
	t.length--
}

func (t *Tree) deleteOneChild(n *Node) {
	// Precondition: n has at most one non-leaf child.
	var child *Node
	if n.right == nil {
		child = n.left
	} else {
		child = n.right
	}
	t.replace(n, child)
	if child == nil {
		return
	}
	if n.Color() == Black {
		if child.Color() == Red {
			child.color = Black
		} else {
			t.deleteCase1(child)
		}
	}
	t.free(n)
}

func (t *Tree) replace(n, child *Node) {
	if child != nil {
		child.parent = n.parent
	}
	if n.parent != nil {
		if n == n.parent.left {
			n.parent.left = child
		} else {
			n.parent.right = child
		}
	} else {
		t.root = child
	}
}

func (t *Tree) free(n *Node) {
	// wait to do
}

func (t *Tree) deleteCase1(n *Node) {
	if n.parent != nil {
		t.deleteCase2(n)
	}
}

func (t *Tree) deleteCase2(n *Node) {
	s := n.Sibling()
	if s.Color() == Red {
		n.parent.color = Red
		s.color = Black
		if n == n.parent.left {
			t.rotateLeft(n.parent)
		} else {
			t.rotateRight(n.parent)
		}
	}
	t.deleteCase3(n)

}

func (t *Tree) deleteCase3(n *Node) {
	s := n.Sibling()
	if s != nil {
		if n.parent.Color() == Black && s.Color() == Black && s.left.Color() == Black && s.right.Color() == Black {
			s.color = Red
			t.deleteCase1(n.parent)
		} else {
			t.deleteCase4(n)
		}
	}
}

func (t *Tree) deleteCase4(n *Node) {
	s := n.Sibling()
	if n.parent.Color() == Red && s.Color() == Black && s.left.Color() == Black && s.right.Color() == Black {
		s.color = Red
		n.parent.color = Black
	} else {
		t.deleteCase5(n)
	}
}

func (t *Tree) deleteCase5(n *Node) {
	s := n.Sibling()
	// This if statement is trivial, due to case 2 (even though case 2 changed
	// the sibling to a sibling's child, the sibling's child can't be red, since
	// no red parent can have a red child).
	if s.Color() == Black {
		// The following statements just force the red to be on the left of the
		// left of the parent, or right of the right, so case six will rotate
		// correctly.
		if n == n.parent.left && s.right.Color() == Black && s.left.Color() == Red {
			// This last test is trivial too due to cases 2-4.
			s.color = Red
			s.left.color = Black
			t.rotateRight(s)
		} else if n == n.parent.right && s.left.Color() == Black && s.right.Color() == Red {
			// This last test is trivial too due to cases 2-4.
			s.color = Red
			s.right.color = Black
			t.rotateLeft(s)
		}
	}
	t.deleteCase6(n)
}

func (t *Tree) deleteCase6(n *Node) {
	s := n.Sibling()
	s.color = n.parent.color
	n.parent.color = Black
	if n == n.parent.left {
		s.right.color = Black
		t.rotateLeft(n.parent)
	} else {
		s.left.color = Black
		t.rotateRight(n.parent)
	}
}

func (t *Tree) rotateLeft(n *Node) {
	if n.parent == nil {
		t.root = n.rotateLeft()
	} else {
		n.rotateLeft()
	}
}

func (t *Tree) rotateRight(n *Node) {
	if n.parent == nil {
		t.root = n.rotateRight()
	} else {
		n.rotateRight()
	}
}

func (t *Tree) search(item Item) *Node {
	n := t.root
	for n != nil {
		if item.Less(n.item) {
			n = n.left
		} else if n.item.Less(item) {
			n = n.right
		} else {
			return n
		}
	}
	return nil
}

// Node represents a node in the red–black tree.
type Node struct {
	color  Color
	left   *Node
	right  *Node
	parent *Node
	item   Item
}

// Color returns the color of this node.
func (n *Node) Color() Color {
	if n == nil {
		return Black
	}
	return n.color
}

// Left returns the left child node.
func (n *Node) Left() *Node {
	if n == nil {
		return nil
	}
	return n.left
}

// Right returns the right child node.
func (n *Node) Right() *Node {
	if n == nil {
		return nil
	}
	return n.right
}

// Parent returns the parent node.
func (n *Node) Parent() *Node {
	if n == nil {
		return nil
	}
	return n.parent
}

// GrandParent returns the grandparent node.
func (n *Node) GrandParent() *Node {
	return n.Parent().Parent()
}

// Sibling returns the sibling node.
func (n *Node) Sibling() *Node {
	if n == nil {
		return nil
	}
	p := n.parent
	if p == nil {
		return nil
	}
	if n == p.left {
		return p.right
	}
	return p.left
}

// Uncle returns the uncle node.
func (n *Node) Uncle() *Node {
	return n.Parent().Sibling()
}

// Item returns the item of this node.
func (n *Node) Item() Item {
	if n == nil {
		return nil
	}
	return n.item
}

// Max returns the max node of this node's subtree.
func (n *Node) Max() *Node {
	if n == nil {
		return nil
	}
	for n.right != nil {
		return n.right.Max()
	}
	return n
}

// Min returns the min node of this node's subtree.
func (n *Node) Min() *Node {
	if n == nil {
		return nil
	}
	for n.left != nil {
		return n.left.Min()
	}
	return n
}

// Last returns the last node less than this node.
func (n *Node) Last() *Node {
	if n == nil {
		return nil
	}
	if n.left != nil {
		return n.left.Max()
	}
	left := n
	p := left.parent
	for p != nil && left == p.left {
		left = p
		p = left.parent
	}
	return p
}

// Next returns the next node more than this node.
func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}
	if n.right != nil {
		return n.right.Min()
	}
	right := n
	p := right.parent
	for p != nil && right == p.right {
		right = p
		p = right.parent
	}
	return p
}

func (n *Node) rotateLeft() *Node {
	newParent := n.right
	n.right = newParent.left
	if newParent.left != nil {
		newParent.left.parent = n
	}
	p := n.parent
	if p != nil {
		if n == p.left {
			p.left = newParent
		} else {
			p.right = newParent
		}
	}
	newParent.parent = p
	n.parent = newParent
	newParent.left = n
	return newParent
}

func (n *Node) rotateRight() *Node {
	newParent := n.left
	n.left = newParent.right
	if newParent.right != nil {
		newParent.right.parent = n
	}
	p := n.parent
	if p != nil {
		if n == p.left {
			p.left = newParent
		} else {
			p.right = newParent
		}
	}
	newParent.parent = p
	n.parent = newParent
	newParent.right = n
	return newParent
}
