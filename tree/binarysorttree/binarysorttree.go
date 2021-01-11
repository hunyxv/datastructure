package binarysorttree

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
				maxParent.lsubtree = nil
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
	var max, left, right uint = 1, 1, 1
	if t.lsubtree != nil {
		left = max + t.lsubtree.Depth()
	}
	if t.rsubtree != nil {
		right = max + t.rsubtree.Depth()
	}

	if left > right {
		return left
	}
	return right
}

// Data 当前节点的保存的数据
func (t *BSTNode) Data() Interface {
	return t.data
}
