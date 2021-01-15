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

func (t *RedBlackNode) brother() *RedBlackNode {
	if t.parent == nil {
		return nil
	}
	if t == t.parent.lsubtree {
		return t.parent.rsubtree
	}
	return t.parent.lsubtree
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
		tmp.rsubtree = t.rsubtree
		if tmp.rsubtree != nil {
			tmp.rsubtree.parent = tmp
		}
		t.rsubtree = tmp
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
		parent := t.parent
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
			if parent == grandparent.lsubtree {
				if t == t.parent.lsubtree { // 插入节点是其父节点的左子节点，左左
					parent.isRed = false
					grandparent.isRed = true
					grandparent.singRotateRight()
				} else { // 插入节点是其父节点的右子节点，左右
					parent := t.parent
					parent.singRotateLeft()
					parent.insertRestructuring()
				}
				return
			}
			// 插入节点的父节点是其祖父节点的右子节点
			if parent == grandparent.rsubtree {
				if t == parent.lsubtree { // 插入节点是其父节点的左子节点
					parent := parent
					parent.singRotateRight()
					parent.insertRestructuring()
				} else { // 插入节点是其父节点的右子节点
					parent.isRed = false
					grandparent.isRed = true
					grandparent.singRotateLeft()
				}
				return
			}
		}
	}
}

// Delete 删除节点
func (t *RedBlackNode) Delete(d Interface) {
	if t.data.Value() < d.Value() {
		t.rsubtree.Delete(d)
	} else if t.data.Value() > d.Value() {
		t.lsubtree.Delete(d)
	} else {
		if t.lsubtree != nil && t.rsubtree != nil {
			rightMin := t.rsubtree.getRightSubTreeMin()
			t.data = rightMin.data
			t.freq = rightMin.freq
			rightMin.deleteRestructuring()
			rightMin.deleteNode()
		} else if t.lsubtree != nil {
			t.data = t.lsubtree.data
			t.freq = t.lsubtree.freq
			t.lsubtree.Delete(t.data)
		} else if t.rsubtree != nil {
			t.data = t.rsubtree.data
			t.freq = t.rsubtree.freq
			t.rsubtree.Delete(t.data)
		} else {
			t.deleteRestructuring()
			t.deleteNode()
		}
	}
}

func (t *RedBlackNode) getLeftSubTreeMax() *RedBlackNode {
	node := t.rsubtree
	if node == nil {
		return t
	}

	return node.getLeftSubTreeMax()
}

func (t *RedBlackNode) getRightSubTreeMin() *RedBlackNode {
	node := t.lsubtree
	if node == nil {
		return t
	}

	return node.getRightSubTreeMin()
}

func (t *RedBlackNode) deleteNode() {
	if t.isRoot {
		t.data = nil
		return
	}

	if t.parent == nil {
		return
	}
	if t == t.parent.lsubtree {
		t.parent.lsubtree = nil
	} else {
		t.parent.rsubtree = nil
	}
	t.parent = nil
	return
}

// deleteRestructuring 删除节点后调整至平衡
func (t *RedBlackNode) deleteRestructuring() {
	if t.isRed {
		t.deleteNode()
		return
	}

	if t == t.parent.lsubtree { // 替换节点是其父节点的左子节点
		brother := t.parent.rsubtree
		if brother != nil {
			if brother.isRed { // 兄弟节点是红色
				brother.isRed = false
				t.parent.isRed = true
				t.parent.singRotateLeft()
				t.deleteRestructuring()
			} else if brother.rsubtree != nil {
				if brother.rsubtree.isRed { // 兄弟节点为黑色，其右子节点为红色（左子节点为任意色）
					brother.isRed = t.parent.isRed
					t.parent.isRed = false
					brother.rsubtree.isRed = false
					t.parent.singRotateLeft()
				} else if brother.lsubtree != nil {
					if brother.lsubtree.isRed { // 兄弟节点为黑色，其左子节点为红色（右子节点为任意色）
						brother.lsubtree.isRed = false
						brother.isRed = true
						brother.singRotateRight()
						t.deleteRestructuring()
					} else { // 兄弟节点为黑色，其左右子节点都为黑色
						brother.isRed = true
						t.parent.deleteRestructuring()
					}
				}
			}
		}
	} else { // 替换节点是其父节点的右子节点（和上面类似）
		brother := t.parent.lsubtree
		if brother != nil {
			if brother.isRed { // 替换节点的兄弟节点是红色
				brother.isRed = false
				t.parent.isRed = true
				t.parent.singRotateRight()
				t.deleteRestructuring()
			} else {
				if brother.lsubtree != nil {
					if brother.lsubtree.isRed { // 兄弟节点为黑色，其左子节点是红色（右子节点为任意色）
						brother.isRed = t.parent.isRed
						t.parent.isRed = false
						brother.lsubtree.isRed = false
						t.parent.singRotateRight()
					} else {
						if brother.rsubtree != nil {
							if brother.rsubtree.isRed { // 兄弟节点为黑色，其右子节点是红色（左子节点为任意色）
								brother.isRed = true
								brother.rsubtree.isRed = false
								brother.singRotateLeft()
								t.deleteRestructuring()
							} else { // 兄弟节点为黑色，其左右子节点都为黑色
								brother.isRed = false
								t.parent.deleteRestructuring()
							}
						}
					}
				}
			}
		}
	}
}

// Data .
func (t *RedBlackNode) Data() Interface {
	return t.data
}
