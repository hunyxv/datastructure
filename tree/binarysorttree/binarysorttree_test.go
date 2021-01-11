package binarysorttree

import (
	"fmt"
	"testing"
)

type data int

func (d data) Value() int {
	return int(d)
}

var binaryTree *BSTNode

// 前序遍历输出
func printTree(tree *BSTNode) {
	fmt.Printf("%v ", tree.Data())
	if tree.lsubtree != nil {
		printTree(tree.lsubtree)
	}
	if tree.rsubtree != nil {
		printTree(tree.rsubtree)
	}
}

func TestInsert(t *testing.T) {
	d := []data{15, 5, 16, 3, 12, 20, 10, 13, 18, 23, 6, 7}
	binaryTree = NewBSTree(d[0])

	for i := 1; i < len(d); i++ {
		binaryTree.Insert(d[i])
	}
	printTree(binaryTree)
}

func TestRepeatedInsert(t *testing.T) {
	d := []data{ 10, 13, 23, 6, 7}

	for i := 0; i < len(d); i++ {
		binaryTree.Insert(d[i])
	}
	printTree(binaryTree)
}

func TestDepth(t *testing.T) {
	t.Logf("Depth: %d", binaryTree.Depth())
}

func TestFind(t *testing.T) {
	t.Logf("Find: %+v", binaryTree.Find(data(18)))
}

func TestDeleteLeafNode(t *testing.T) {
	t.Logf("Delete: %t", binaryTree.Delete(data(3)))
	printTree(binaryTree)
}

func TestDeleteSingleBranch(t *testing.T) {
	t.Logf("Delete: %t", binaryTree.Delete(data(10)))
	printTree(binaryTree)
}

func TestDeleteNode(t *testing.T) {
	t.Logf("Delete: %t", binaryTree.Delete(data(20)))
	printTree(binaryTree)
}