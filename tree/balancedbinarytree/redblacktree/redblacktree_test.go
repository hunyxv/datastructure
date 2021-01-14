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

func TestSingRotateLeft(t *testing.T){
	tree := NewReadBlackTree(data(5))
	prt := tree
	prt.rsubtree = newNode(prt, data(7))

	prt = prt.rsubtree
	prt.rsubtree = newNode(prt, data(8))

	tree.singRotateLeft()
	printTree(tree)
}

func TestSingRotateLeft2(t *testing.T){
	tree := NewReadBlackTree(data(7))
	prt := tree
	prt.lsubtree = newNode(prt, data(5))

	prt = prt.lsubtree
	prt.rsubtree = newNode(prt, data(6))

	prt.singRotateLeft()
	printTree(tree)
}

func TestSingRotateRight(t *testing.T){
	tree := NewReadBlackTree(data(7))
	prt := tree
	prt.lsubtree = newNode(prt, data(6))
	
	x:=prt.lsubtree
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
	for i:=1;i<len(data); i++ {
		tree.Insert(data[i])
		printTree(tree)
		fmt.Println()
	}

	printTree(tree)
}
