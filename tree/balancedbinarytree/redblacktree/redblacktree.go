package redblacktree

// Interface .
type Interface interface {
	Value() int
}

// RedBlackNode 红黑树节点
type RedBlackNode struct {
	data     Interface
	freq     int
	isRed    bool
	isRoot   bool
	parent   *RedBlackNode
	lsubtree *RedBlackNode
	rsubtree *RedBlackNode
}

// NewReadBlackTree 创建一棵红黑树
func NewReadBlackTree(d Interface) *RedBlackNode {
	return &RedBlackNode{
		data:   d,
		freq:   1,
		isRoot: true,
	}
}

func newNode(parent *RedBlackNode, d Interface) *RedBlackNode {
	return &RedBlackNode{
		data:   d,
		freq:   1,
		isRed:  true,
		parent: parent,
	}
}

// Insert 插入数据
func (t *RedBlackNode) Insert(d Interface) {
	if t.data.Value() > d.Value() {
		if t.lsubtree != nil {
			t.lsubtree.Insert(d)
		} else {
			t.lsubtree = newNode(t, d)
			t.lsubtree.insertRestructuring()
		}
	} else if t.data.Value() < d.Value() {
		if t.rsubtree != nil {
			t.rsubtree.Insert(d)
		} else {
			t.rsubtree = newNode(t, d)
			t.rsubtree.insertRestructuring()
		}
	} else {
		t.freq++
	}
}

func (t *RedBlackNode) grandparent() *RedBlackNode {
	if t.parent == nil {
		return nil
	}
	return t.parent.parent
}

func (t *RedBlackNode) uncle() *RedBlackNode {
	if t.grandparent() == nil {
		return nil
	}
	if t.parent == t.grandparent().lsubtree {
		return t.grandparent().rsubtree
	}
	return t.grandparent().lsubtree
}

// 左旋转
func (t *RedBlackNode) singRotateLeft() {
	if !t.isRoot {
		right := t.rsubtree
		if t.parent.lsubtree == t {
			t.parent.lsubtree = right
		} else {
			t.parent.rsubtree = right
		}
		right.parent = t.parent
		t.rsubtree = right.lsubtree
		right.lsubtree = t
		t.parent = right
	} else {
		tmp := t.rsubtree
		t.data, tmp.data = tmp.data, t.data
		t.isRed, tmp.isRed = tmp.isRed, t.isRed
		t.freq, tmp.freq = tmp.freq, t.freq
		t.rsubtree = tmp.rsubtree
		t.rsubtree.parent = t
		tmp.rsubtree = tmp.lsubtree
		if tmp.rsubtree != nil {
			tmp.rsubtree.parent = tmp
		}
		tmp.lsubtree = t.lsubtree
		if tmp.lsubtree != nil {
			tmp.lsubtree.parent = tmp
		}
		t.lsubtree = tmp
	}
}

// 右旋转
func (t *RedBlackNode) singRotateRight() {
	if !t.isRoot {
		left := t.lsubtree
		if t.parent.lsubtree == t {
			t.parent.lsubtree = left
		} else {
			t.parent.rsubtree = left
		}
		left.parent = t.parent
		t.lsubtree = left.rsubtree
		left.rsubtree = t
		t.parent = left
	} else {
		tmp := t.lsubtree
		t.data, tmp.data = tmp.data, t.data
		t.isRed, tmp.isRed = tmp.isRed, t.isRed
		t.freq, tmp.freq = tmp.freq, t.freq
		t.lsubtree = tmp.lsubtree
		t.lsubtree.parent = t
		tmp.lsubtree = tmp.rsubtree
		if tmp.lsubtree != nil {
			tmp.lsubtree.parent = tmp
		}

		tmp.rsubtree = t.rsubtree
		if tmp.rsubtree != nil {
			tmp.rsubtree.parent = tmp
		}
		t.rsubtree = tmp
	}
}

// insertRestructuring 插入新节点后调整至平衡
func (t *RedBlackNode) insertRestructuring() {
	if t.isRoot {
		t.isRed = false
		return
	}

	// 父节点为红色
	if t.parent.isRed {
		uncle := t.uncle()
		grandparent := t.grandparent()
		// 叔叔节点不为空并且是红色
		if uncle != nil && uncle.isRed {
			t.parent.isRed = false
			uncle.isRed = false
			grandparent.isRed = true
			grandparent.insertRestructuring()
			return
		}

		// 叔叔节点不存在或是黑色节点
		if uncle == nil || !uncle.isRed {
			// 插入节点的父节点是其祖父节点的左子节点
			if t.parent == grandparent.lsubtree {
				if t == t.parent.lsubtree { // 插入节点是其父节点的左子节点
					t.parent.isRed = false
					grandparent.isRed = true
					grandparent.singRotateRight()
				} else { // 插入节点是其父节点的右子节点
					t.parent.singRotateLeft()
					t.parent.insertRestructuring()
				}
				return
			}
			// 插入节点的父节点是其祖父节点的右子节点
			if t.parent == grandparent.rsubtree {
				if t == t.parent.lsubtree { // 插入节点是其父节点的左子节点
					t.parent.singRotateRight()
					t.parent.insertRestructuring()
				} else { // 插入节点是其父节点的右子节点
					t.parent.isRed = false
					grandparent.isRed = true
					grandparent.singRotateLeft()
				}
				return
			}
		}
	}
}

// Data .
func (t *RedBlackNode) Data() Interface {
	return t.data
}
