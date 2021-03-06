package redblacktree

import (
	"fmt"
	"testing"
)

type data int

func (d data) Value() int {
	return int(d)
}

// 前序遍历输出
func printTree(tree *RedBlackNode) {
	fmt.Printf("%v(%t) ", tree.Data(), tree.isRed)
	if tree.lsubtree != nil {
		printTree(tree.lsubtree)
	}
	if tree.rsubtree != nil {
		printTree(tree.rsubtree)
	}
}

func TestSingRotateLeft(t *testing.T) {
	tree := NewReadBlackTree(data(5))
	prt := tree
	prt.rsubtree = newNode(prt, data(7))

	prt = prt.rsubtree
	prt.rsubtree = newNode(prt, data(8))

	tree.singRotateLeft()
	printTree(tree)
}

func TestSingRotateLeft2(t *testing.T) {
	tree := NewReadBlackTree(data(7))
	prt := tree
	prt.lsubtree = newNode(prt, data(5))

	prt = prt.lsubtree
	prt.rsubtree = newNode(prt, data(6))

	prt.singRotateLeft()
	printTree(tree)
}

func TestSingRotateRight(t *testing.T) {
	tree := NewReadBlackTree(data(7))
	prt := tree
	prt.lsubtree = newNode(prt, data(6))

	x := prt.lsubtree
	prt = prt.lsubtree
	prt.lsubtree = newNode(prt, data(3))

	prt = prt.lsubtree
	prt.lsubtree = newNode(prt, data(2))

	x.singRotateRight()
	printTree(tree)
}

func TestInsert(t *testing.T) {
	data := []data{42, 37, 18, 12, 11, 6, 5, 1}
	tree := NewReadBlackTree(data[0])
	for i := 1; i < len(data); i++ {
		tree.Insert(data[i])
	}

	printTree(tree)
}

func TestDelete(t *testing.T) {
	ds := []data{17, 33, 37, 42, 50, 48, 88, 66, 55, 6, 12, 16}
	tree := NewReadBlackTree(ds[0])
	for i := 1; i < len(ds); i++ {
		tree.Insert(ds[i])
	}

	printTree(tree)
	fmt.Println()
	tree.Delete(data(37))
	printTree(tree)
}

func TestTraversal(t *testing.T) {
	ds := []data{17, 33, 37, 42, 50, 48, 88, 66, 55, 6, 12, 16}
	tree := NewReadBlackTree(ds[0])
	for i := 1; i < len(ds); i++ {
		tree.Insert(ds[i])
	}
	tree.Traversal(func(i Interface) bool {
		t.Log(i.Value())
		if i.Value() == 50 {
			return false
		}
		return true
	})
}
