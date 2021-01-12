package binarysorttree

import (
	"github.com/hunyxv/datastructure/stack"
)

// Interface .
type Interface interface {
	Value() int
}

// BSTNode binary sort tree node
type BSTNode struct {
	data     Interface
	freq     uint
	isRoot   bool
	lsubtree *BSTNode
	rsubtree *BSTNode
}

// NewBSTree 创建跟节点
func NewBSTree(d Interface) *BSTNode {
	return &BSTNode{
		data:   d,
		freq:   1,
		isRoot: true,
	}
}

func newNode(d Interface) *BSTNode {
	return &BSTNode{data: d, freq: 1}
}

// Insert 插入
func (t *BSTNode) Insert(d Interface) {
	if d.Value() > t.data.Value() {
		if t.rsubtree == nil {
			t.rsubtree = newNode(d)
			return
		}
		t.rsubtree.Insert(d)
	} else if d.Value() < t.data.Value() {
		if t.lsubtree == nil {
			t.lsubtree = newNode(d)
			return
		}
		t.lsubtree.Insert(d)
	} else {
		t.freq++
	}
}

// Find 查找
func (t *BSTNode) Find(d Interface) *BSTNode {
	if d.Value() > t.data.Value() && t.rsubtree != nil {
		return t.rsubtree.Find(d)
	} else if d.Value() < t.data.Value() && t.lsubtree != nil {
		return t.lsubtree.Find(d)
	} else if d.Value() == t.data.Value() {
		return t
	}

	return nil
}

func (t *BSTNode) getLeftSubTreeMax(parent *BSTNode) (*BSTNode, *BSTNode) {
	node := t.rsubtree
	if node == nil {
		return parent, t
	}

	return node.getLeftSubTreeMax(t)
}

func (t *BSTNode) getRightSubTreeMin(parent *BSTNode) (*BSTNode, *BSTNode) {
	node := t.lsubtree
	if node == nil {
		return parent, t
	}

	return node.getRightSubTreeMin(t)
}

// Delete 删除
func (t *BSTNode) Delete(d Interface) bool {
	current := t
	parent := t
	for current != nil {
		if d.Value() > current.data.Value() {
			parent = current
			current = current.rsubtree
		} else if d.Value() < current.data.Value() {
			parent = current
			current = current.lsubtree
		} else {
			if current.rsubtree != nil && current.lsubtree != nil {
				maxParent, max := current.lsubtree.getLeftSubTreeMax(current)
				if maxParent != current {
					maxParent.rsubtree = max.lsubtree
				} else {
					maxParent.lsubtree = nil
				}
				
				current.data = max.data
			} else if current.lsubtree != nil {
				tmp := current.lsubtree
				current.data = tmp.data
				current.rsubtree = tmp.rsubtree
				current.lsubtree = tmp.lsubtree
				tmp.data = nil
				tmp.rsubtree = nil
				tmp.lsubtree = nil
			} else if current.rsubtree != nil {
				tmp := current.rsubtree
				current.data = tmp.data
				current.rsubtree = tmp.rsubtree
				current.lsubtree = tmp.lsubtree
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
	return false
}

// Depth 树的深度
func (t *BSTNode) Depth() uint {
	var  left, right uint =  1, 1
	if t.lsubtree != nil {
		left = 1 + t.lsubtree.Depth()
	}
	if t.rsubtree != nil {
		right = 1 + t.rsubtree.Depth()
	}

	if t.isRoot {
		return max(left, right) -1
	}
	return max(left, right)
}

// Data 当前节点的保存的数据
func (t *BSTNode) Data() Interface {
	return t.data
}

// Traversal 遍历各个值（深度优先--中序遍历 (从小到大)）
func (t *BSTNode) Traversal(f func(Interface) bool) {
	current := t
	sk := stack.NewStack(int(t.Depth()))
	for current != nil || !sk.IsEmpty() {
		if current != nil {
			sk.Push(current)
			current = current.lsubtree
			continue
		}
		el, _ := sk.Pop()
		node := el.(*BSTNode)
		if !f(node.data) {
			return
		}
		current = node.rsubtree
	}
}

func max(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}