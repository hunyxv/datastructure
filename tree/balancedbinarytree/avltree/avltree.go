package tree

// Interface .
type Interface interface {
	Value() int
}

// AVLTree avl 自平衡二叉树
type AVLTree struct {
	data     Interface
	freq     uint
	height   uint
	lsubtree *AVLTree
	rsubtree *AVLTree
}

// singRotateLeft 向左旋转
func (t *AVLTree) singRotateLeft() {
	
}