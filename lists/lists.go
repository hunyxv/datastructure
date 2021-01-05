package lists

import (
	"errors"
)

// Tag .
type Tag int

const (
	// Elem 元素
	Elem Tag = iota + 1
	// SubList 子表
	SubList
)

var (
	// ErrEmpty 空表
	ErrEmpty = errors.New("lists is empty")
)

// Interface .
type Interface interface {
	GetTag() Tag
	GetElem() interface{}
	Next() (Interface, bool)
	CopyGList() Interface
	GListLength() int
	GListDepth() int
	GListEmpth() bool
	GetHead() (Interface, error)
	GetTail() (Interface, error)
	Insert(int, Interface) bool
	Delete(int)
	Add(interface{})
	Traverse(func(interface{}) bool)
}

var _ Interface = (*GLNode)(nil)

// GLNode 广义表 节点结构
type GLNode struct {
	ElemTag    Tag
	Atom       interface{}
	hptr, tptr *GLNode
}

// GetTag 返回节点 tag
func (n *GLNode) GetTag() Tag {
	return n.ElemTag
}

// GetElem 返回该节点值，元素/子表
func (n *GLNode) GetElem() interface{} {
	if n.ElemTag == Elem {
		return n.Atom
	}
	return n.hptr
}

func (n *GLNode) next() (*GLNode, bool) {
	if n.tptr == nil {
		return nil, false
	}
	return n.tptr, true
}

// Next 下一个节点
func (n *GLNode) Next() (Interface, bool) {
	return n.next()
}

func (n *GLNode) copyElem() *GLNode {
	if n.ElemTag == Elem {
		node := *n
		return &node
	}

	head := &GLNode{ElemTag: SubList}
	node := n.hptr
	if node == nil {
		return head
	}

	head.hptr = node.copyElem()
	ptr := head.hptr
	for node, has := node.next(); has; node, has = node.next() {
		ptr.tptr = node.copyElem()
		ptr = ptr.tptr
	}
	return head
}

// CopyGList 返回 lists 的复制
func (n *GLNode) CopyGList() Interface {
	head := &GLNode{ElemTag: SubList}
	ptr := head

	node := n.hptr
	if node == nil {
		return head
	}

	ptr.hptr = node.copyElem()
	ptr = ptr.hptr

	for node, has := node.next(); has; node, has = node.next() {
		ptr.tptr = node.copyElem()
		ptr = ptr.tptr
	}
	return head
}

// GListLength 求广义表的长度(即第一层的元素个数)
func (n *GLNode) GListLength() int {
	head := n.hptr
	if head == nil {
		return 0
	}
	var i = 1
	for next, has := head.next(); has; next, has = next.next() {
		i++
	}
	return i
}

// GListDepth 返回广义表的深度（展开后所含括号的层数）
func (n *GLNode) GListDepth() int {
	head, err := n.GetHead()
	if err != nil {
		return 1
	}

	var max, dep int = 1, 1
	for next, has := head, true; has; next, has = next.Next() {
		if next.GetTag() != SubList {
			continue
		}
		dep = 1 + next.GListDepth()
		if dep > max {
			max = dep
		}
	}
	return max
}

// GListEmpth 广义表是否为空
func (n *GLNode) GListEmpth() bool {
	return n.tptr == nil && n.hptr == nil
}

// GetHead 获取表头 （空表无表头表尾）
func (n *GLNode) GetHead() (Interface, error) {
	if n.hptr != nil {
		return n.hptr, nil
	}
	return nil, ErrEmpty
}

// GetTail 获取表尾 （空表无表头表尾）
func (n *GLNode) GetTail() (Interface, error) {
	if head := n.hptr; head != nil {
		if tail, has := head.next(); has {
			return &GLNode{ElemTag: SubList, hptr: tail}, nil
		}
		return nil, ErrEmpty
	}
	return &GLNode{ElemTag: SubList}, nil
}

// Insert 在 i 索引下插入一个节点
func (n *GLNode) Insert(i int, elem Interface) bool {
	node, ok := elem.(*GLNode)
	if !ok {
		return false
	}
	if i == 0 {
		n.hptr, node.tptr = node, n.hptr
		return true
	}

	preOrder := n.hptr
	for j := 0; preOrder != nil && j < i-2; j++ {
		preOrder = preOrder.tptr
	}
	if preOrder != nil {
		postOrder := preOrder.tptr
		node.tptr = postOrder
	}
	preOrder.tptr = node

	return true
}

// Add 从头部加入一个节点
func (n *GLNode) Add(e interface{}) {
	if sublist, ok := e.(*GLNode); ok {
		n.Insert(0, sublist)
		return
	}

	node := &GLNode{Atom: e, ElemTag: Elem}
	n.Insert(0, node)
}

// Delete 删除 i 节点
func (n *GLNode) Delete(i int) {
	if i == 0 && n.hptr != nil {
		n.hptr = n.hptr.tptr
		return
	}

	preOrder := n.hptr
	for j := 0; preOrder != nil && j < i-2; j++ {
		preOrder = preOrder.tptr
	}
	if preOrder == nil || preOrder.tptr == nil {
		return
	}
	preOrder.tptr = preOrder.tptr.tptr
	return
}

// Traverse 遍历广义表，使用函数f处理每个元素
func (n *GLNode) Traverse(f func(interface{}) bool) {
	if n.ElemTag == Elem {
		if !f(n.Atom) {
			return
		}
	} else {
		if n.hptr != nil {
			n.hptr.Traverse(f)
		}
	}
	if next, has := n.next(); has {
		next.Traverse(f)
	}

	return
}
