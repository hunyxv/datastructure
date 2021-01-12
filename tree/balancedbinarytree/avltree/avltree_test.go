package avltree

import (
	"fmt"
	"testing"
)

type data int

func (d data) Value() int {
	return int(d)
}

// 前序遍历输出
func printTree(tree *AVLTree) {
	fmt.Printf("%v(%d) ", tree.Data(), tree.Height())
	if tree.lsubtree != nil {
		printTree(tree.lsubtree)
	}
	if tree.rsubtree != nil {
		printTree(tree.rsubtree)
	}
}

func _init(data []data) *AVLTree {
	var avlTree *AVLTree
	avlTree = NewAVLTree(data[0])
	for i := 1; i < len(data); i++ {
		avlTree.Insert(data[i])
	}

	printTree(avlTree)
	fmt.Println()
	return avlTree
}

func TestInsert(t *testing.T) {
	_init([]data{7, 5, 9, 6, 3, 1})
}

func TestSingRotateLeft(t *testing.T) {
	avlTree := _init([]data{7, 5, 9, 6, 3, 1})
	printTree(avlTree)
}

func TestSingRotateRight(t *testing.T) {
	avlTree := _init([]data{5, 1, 8, 6, 10, 11})
	printTree(avlTree)
}

func TestDoubleRotateLR(t *testing.T) {
	avlTree := _init([]data{9, 3, 1, 6, 5, 7, 10})
	printTree(avlTree)
}

func TestDoubleRotateRL(t *testing.T) {
	avlTree := _init([]data{3, 1, 8, 6, 5, 7, 9})
	printTree(avlTree)
}

func TestDelete(t *testing.T) {
	avlTree := _init([]data{1, 2, 3, 4, 5, 6, 7, 8, 9})
	avlTree.Delete(data(2))
	printTree(avlTree)  // 没有打破平衡
	fmt.Println()
	avlTree.Delete(data(1))  // 自平衡
	printTree(avlTree)
}

func TestDepth(t *testing.T) {
	avlTree := _init([]data{1, 2, 3, 4, 5, 6, 7, 8, 9})
	t.Logf("Depth: %d", avlTree.Depth())
}

func TestTraversal(t *testing.T) {
	avlTree := _init([]data{15, 5, 16, 3, 12, 20, 10, 13, 18, 23, 6, 7})
	printTree(avlTree)
	avlTree.Traversal(func(i Interface) bool {
		t.Log(i.Value())
		if i.Value() == 15 {
			return false
		}
		return true
	})
}