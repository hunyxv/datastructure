package heap

import (
	"sync"
)

type node struct {
	parent *node
	child  *node // 任一孩子的指针（所有孩子是一个链表）
	left   *node
	right  *node

	degree int
	marded bool

	value Interface
}

type FibHeap struct {
	main           *node
	cursor         *node
	num            uint
	maxNumOfDegree int

	mux     sync.RWMutex
	compare func(a, b Interface) bool
}

func NewFibHeap(maxDegree int, t T) *FibHeap {
	heap := &FibHeap{
		maxNumOfDegree: maxDegree,
	}
	if t == MinHeap {
		heap.compare = le
	} else {
		heap.compare = ge
	}
	return heap
}

func (h *FibHeap) Insert(val Interface) {
	h.mux.Lock()
	defer h.mux.Unlock()

	if h.main == nil {
		n := &node{
			value: val,
		}
		n.left, n.right = n, n
		h.main = n
		return
	}

	n := &node{
		left:  h.main.left,
		right: h.main,

		value: val,
	}
	h.main.left.right = n
	h.main.left = n
	if h.compare(n.value, h.main.value) {
		h.main = n
	}
}

func (h *FibHeap) Pop() (Interface, error) {
	h.mux.Lock()
	defer h.mux.Unlock()

	if h.main == nil {
		return nil, ErrEmpty
	}

	top := h.main
	if top.child == nil {
		if top.left == top {
			h.main = nil
		} else {
			h.cursor = top.right
			top.left.right = top.right
			top.right.left = top.left
			h.consolidate()
		}
		return top.value, nil
	}

	tmp := top.child
	for head := tmp; tmp.right != head; tmp = tmp.right {
		tmp.parent = nil
	}

	if top.right == top {
		h.cursor = top.child
		h.consolidate()
		return top.value, nil
	}

	top.left.right = top.child
	top.child.left.right = top.right

	top.right.left = top.child.left
	top.child.left = top.left

	h.cursor = top.child
	h.consolidate()

	top.child.parent = nil
	top.child = nil
	top.left = nil
	top.right = nil
	return top.value, nil
}

func (h *FibHeap) consolidate() {
	list := make([]*node, 0, 3)
	main := h.cursor
	for head := h.cursor; h.cursor.right != head; h.cursor = h.cursor.right {
		list = append(list, h.cursor)
	}
	list = append(list, h.cursor)

	table := make(map[int]*node)
	for _, p := range list {
		if !h.compare(main.value, p.value) {
			main = p
		}
		h.union(table, p)
	}

	h.main = main
}

func (h *FibHeap) union(table map[int]*node, cur *node) {
	if n, ok := table[cur.degree]; ok {
		if h.compare(n.value, cur.value) {
			cur.left.right = cur.right
			cur.right.left = cur.left
			h.addChild(n, cur)
			delete(table, cur.degree)
			n.degree++
			h.union(table, n)
		} else {
			n.left.right = n.right
			n.right.left = n.left
			h.addChild(cur, n)
			delete(table, n.degree)
			cur.degree++
			h.union(table, cur)
		}
	} else {
		table[cur.degree] = cur
	}
}

func (h *FibHeap) addChild(parent *node, n *node) {
	n.parent = parent
	if parent.child == nil {
		parent.child = n
		n.left, n.right = n, n
		return
	}

	parent.child.left.right = n
	n.left = parent.child.left
	parent.child.left = n
	n.right = parent.child
}
