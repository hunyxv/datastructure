package avltree

import "github.com/hunyxv/datastructure/stack"

// Interface .
type Interface interface {
	Value() int
}

// AVLTree avl 自平衡二叉树
type AVLTree struct {
	data     Interface
	freq     int
	isRoot   bool
	height   int
	lsubtree *AVLTree
	rsubtree *AVLTree
}

func NewAVLTree(d Interface) *AVLTree {
	return &AVLTree{
		data:   d,
		freq:   1,
		isRoot: true,
		height: 1,
	}
}

func newNode(d Interface) *AVLTree {
	return &AVLTree{data: d, freq: 1}
}

// Height 当前节点的高度
func (t *AVLTree) Height() int {
	if t != nil {
		return t.height
	}
	return -1
}

// Data 当前节点的保存的数据
func (t *AVLTree) Data() Interface {
	return t.data
}

// Insert 插入
func (t *AVLTree) Insert(d Interface) {
	defer func() {
		t.height = max(t.lsubtree.Height(), t.rsubtree.Height()) + 1
	}()
	if d.Value() > t.data.Value() {
		if t.rsubtree == nil {
			t.rsubtree = newNode(d)
			return
		}
		t.rsubtree.Insert(d)
		if t.rsubtree.Height()-t.lsubtree.Height() == 2 {
			if t.rsubtree.data.Value() < d.Value() {
				t.singRotateRight()
			} else {
				t.doubleRotateRL()
			}
		}
	} else if d.Value() < t.data.Value() {
		if t.lsubtree == nil {
			t.lsubtree = newNode(d)
			return
		}
		t.lsubtree.Insert(d)
		if t.lsubtree.Height()-t.rsubtree.Height() == 2 {
			if t.lsubtree.data.Value() > d.Value() {
				t.singRotateLeft()
			} else {
				t.doubleRotateLR()
			}
		}
	} else {
		t.freq++
	}
}

// Find 查找
func (t *AVLTree) Find(d Interface) *AVLTree {
	if d.Value() > t.data.Value() && t.rsubtree != nil {
		return t.rsubtree.Find(d)
	} else if d.Value() < t.data.Value() && t.lsubtree != nil {
		return t.lsubtree.Find(d)
	} else if d.Value() == t.data.Value() {
		return t
	}
	return nil
}

func (t *AVLTree) getLeftSubTreeMax(parent *AVLTree) (*AVLTree, *AVLTree) {
	node := t.rsubtree
	if node == nil {
		return parent, t
	}

	return node.getLeftSubTreeMax(t)
}

func (t *AVLTree) getRightSubTreeMin(parent *AVLTree) (*AVLTree, *AVLTree) {
	node := t.lsubtree
	if node == nil {
		return parent, t
	}

	return node.getRightSubTreeMin(t)
}

// Delete 删除
func (t *AVLTree) Delete(d Interface) bool {
	if t == nil {
		return false
	}

	return t.delete(t, d)
}

func (t *AVLTree) delete(parent *AVLTree, d Interface) (b bool) {
	defer func() {
		t.height = max(t.rsubtree.Height(), t.lsubtree.Height()) + 1
	}()
	if d.Value() < t.data.Value() {
		b = t.lsubtree.delete(t, d)
		if t.rsubtree.Height()-t.lsubtree.Height() == 2 {
			if t.rsubtree.lsubtree != nil && t.rsubtree.lsubtree.Height() > t.rsubtree.rsubtree.Height() {
				t.doubleRotateRL()
			} else {
				t.singRotateRight()
			}
		}
		return
	} else if d.Value() > t.data.Value() {
		b = t.rsubtree.delete(t, d)
		if t.lsubtree.Height()-t.rsubtree.Height() == 2 {
			if t.lsubtree.rsubtree != nil && t.lsubtree.rsubtree.Height() > t.lsubtree.lsubtree.Height() {
				t.doubleRotateLR()
			} else {
				t.singRotateLeft()
			}
		}
		return
	} else {
		if t.lsubtree != nil && t.rsubtree != nil {
			maxParent, max := t.lsubtree.getLeftSubTreeMax(t)
			if maxParent != t {
				maxParent.rsubtree = max.lsubtree
			} else {
				maxParent.lsubtree = nil
			}
			t.data = max.data
		} else if t.lsubtree != nil {
			tmp := t.lsubtree
			t.data = tmp.data
			t.rsubtree = tmp.rsubtree
			t.lsubtree = tmp.lsubtree
			tmp.data = nil
			tmp.rsubtree = nil
			tmp.lsubtree = nil
		} else if t.rsubtree != nil {
			tmp := t.rsubtree
			t.data = tmp.data
			t.rsubtree = tmp.rsubtree
			t.lsubtree = tmp.lsubtree
			tmp.data = nil
			tmp.rsubtree = nil
			tmp.lsubtree = nil
		} else {
			if parent.data.Value() < d.Value() {
				parent.rsubtree = nil
			} else {
				parent.lsubtree = nil
			}
		}
		return true
	}
}

// Depth 树的深度(根节点深度为0)
func (t *AVLTree) Depth() int {
	var left, right int = 1, 1
	if t.lsubtree != nil {
		left = 1 + t.lsubtree.Depth()
	}
	if t.rsubtree != nil {
		right = 1 + t.rsubtree.Depth()
	}
	if t.isRoot {
		return max(left, right) - 1
	}
	return max(left, right)
}

// Traversal 遍历各个值（深度优先--中序遍历 (从小到大)）
func (t *AVLTree) Traversal(f func(Interface) bool) {
	current := t
	sk := stack.NewStack(int(t.Depth()))
	for current != nil || !sk.IsEmpty() {
		if current != nil {
			sk.Push(current)
			current = current.lsubtree
			continue
		}
		el, _ := sk.Pop()
		node := el.(*AVLTree)
		if !f(node.data) {
			return
		}
		current = node.rsubtree
	}
}

// singRotateLeft 向左旋转(左左)
//
func (t *AVLTree) singRotateLeft() {
	tmp := t.lsubtree
	t.data, tmp.data = tmp.data, t.data
	t.lsubtree = tmp.lsubtree
	tmp.lsubtree = tmp.rsubtree
	tmp.rsubtree = t.rsubtree
	t.rsubtree = tmp
	tmp.height = max(tmp.lsubtree.Height(), tmp.rsubtree.Height()) + 1
	t.height = max(t.lsubtree.Height(), t.rsubtree.Height()) + 1
	return
}

// singRotateRight 向右旋转（右右）
func (t *AVLTree) singRotateRight() {
	tmp := t.rsubtree
	t.data, tmp.data = tmp.data, t.data
	t.rsubtree = tmp.rsubtree
	tmp.rsubtree = tmp.lsubtree
	tmp.lsubtree = t.lsubtree
	t.lsubtree = tmp

	tmp.height = max(tmp.lsubtree.Height(), tmp.rsubtree.Height()) + 1
	t.height = max(t.lsubtree.Height(), t.rsubtree.Height()) + 1
	return
}

// doubleRotateLR 左右情况下的双旋转
func (t *AVLTree) doubleRotateLR() {
	t.lsubtree.singRotateRight()
	t.singRotateLeft()
}

// 右左情况下的双旋转
func (t *AVLTree) doubleRotateRL() {
	t.rsubtree.singRotateLeft()
	t.singRotateRight()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
