package heap

import (
	"errors"
	"math"
	"sync"
)

type node struct {
	parent *node
	child  *node // 任一孩子的指针（所有孩子是一个链表）
	left   *node
	right  *node

	degree int
	marked bool

	value Interface
	key   interface{}
	v     float64
}

type FibHeap struct {
	main   *node
	cursor *node
	num    uint
	index  map[interface{}]*node

	t       T
	mux     sync.RWMutex
	compare func(a, b float64) bool
}

// NewFibHeap 初始化 fibonacci heap
func NewFibHeap(t T) *FibHeap {
	heap := &FibHeap{
		index: make(map[interface{}]*node),

		t: t,
	}
	if t == MinHeap {
		heap.compare = func(a, b float64) bool {
			return a <= b
		}
	} else {
		heap.compare = func(a, b float64) bool {
			return a >= b
		}
	}
	return heap
}

// T heap 的类型
//	最小堆/最大堆
func (h *FibHeap) T() T {
	return h.t
}

// Insert 插入一个元素
//	元素的排序指标范围为 (-inf, +inf)
func (h *FibHeap) Insert(val Interface) error {
	if math.IsInf(val.Value(), 0) {
		return errors.New("infinity is for internal use only")
	}

	h.mux.Lock()
	if _, ok := h.index[val.Key()]; ok {
		h.mux.Unlock()
		return errors.New("duplicate key is not allowed")
	}

	h.insertValue(val)
	h.mux.Unlock()
	return nil
}

func (h *FibHeap) insertValue(val Interface) {
	n := &node{
		value: val,
		key:   val.Key(),
		v:     val.Value(),
	}
	h.index[val.Key()] = n
	if h.main == nil {
		n.left, n.right = n, n
		h.main = n
		return
	}

	h.insertNode(n)
	if h.compare(n.v, h.main.v) {
		h.main = n
	}
}

func (h *FibHeap) insertNode(n *node) {
	n.left = h.main.left
	n.right = h.main
	h.main.left.right = n
	h.main.left = n
}

// Pop 返回并移除堆顶元素
func (h *FibHeap) Pop() (val Interface, err error) {
	if h.main == nil {
		return nil, ErrEmpty
	}

	h.mux.Lock()
	node, err := h.popNode()
	h.mux.Unlock()
	if err != nil {
		return nil, err
	}
	return node.value, nil
}

func (h *FibHeap) popNode() (*node, error) {
	if h.main == nil {
		return nil, ErrEmpty
	}

	top := h.main
	delete(h.index, top.value.Key())

	if top.child == nil {
		if top.left == top {
			h.main = nil
		} else {
			h.cursor = top.right
			top.left.right = top.right
			top.right.left = top.left
			h.consolidate()
		}
		return top, nil
	}

	tmp := top.child
	for head := top.child; tmp.right != head; tmp = tmp.right {
		tmp.parent = nil
	}
	tmp.parent = nil

	if top.right == top {
		h.cursor = top.child
		top.child = nil
		h.consolidate()
		return top, nil
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
	return top, nil
}

func (h *FibHeap) moveChildren2Root(p *node) {
	p.degree = 0
	if p.child == nil {
		return
	}
	main := h.main

	main.left.right = p.child
	p.child.left.right = main.right
	main.right.left = p.child.right
	p.child.left = main.left
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
		if !h.compare(main.v, p.v) {
			main = p
		}
		h.unionSubtree(table, p)
	}

	h.main = main
}

func (h *FibHeap) unionSubtree(table map[int]*node, cur *node) {
	if n, ok := table[cur.degree]; ok {
		if h.compare(n.v, cur.v) {
			cur.left.right = cur.right
			cur.right.left = cur.left
			h.addChild(n, cur)
			delete(table, cur.degree)
			n.degree++
			h.unionSubtree(table, n)
		} else {
			n.left.right = n.right
			n.right.left = n.left
			h.addChild(cur, n)
			delete(table, n.degree)
			cur.degree++
			h.unionSubtree(table, cur)
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

// Peek 返回堆顶元素（不移除）
func (h *FibHeap) Peek() (val Interface, err error) {
	h.mux.RLock()

	if h.main == nil {
		h.mux.RUnlock()
		return nil, ErrEmpty
	}

	val = h.main.value
	h.mux.RUnlock()
	return
}

// Union 合并另一个堆
func (h *FibHeap) Union(target *FibHeap) error {
	h.mux.Lock()

	for k := range target.index {
		if _, exists := h.index[k]; exists {
			h.mux.Unlock()
			return errors.New("duplicate tag is found in the target heap")
		}
	}

	for _, node := range target.index {
		h.insertValue(node.value)
	}

	h.mux.Unlock()
	return nil
}

// UpdateValue 根据元素的 key 更新其值
func (h *FibHeap) UpdateValue(val Interface) {
	h.mux.Lock()
	if p, ok := h.index[val.Key()]; ok {
		p.value = val
		p.v = val.Value()
		if val.Value() < p.value.Value() {
			h.decreaseValue(p)
		} else {
			h.increaseValue(p)
		}
	}
	h.mux.Unlock()
}

func (h *FibHeap) decreaseValue(p *node) {
	parent := p.parent
	if h.t == MinHeap {
		if parent == nil { // 是根链表节点
			if h.compare(p.v, h.main.v) {
				h.main = p
			}
			return
		}

		if !h.compare(p.v, parent.v) { // 没有破坏最x堆性质
			return
		}

		h.cut(p)
		if h.compare(p.v, h.main.v) {
			h.main = p
		}
		h.cascadingCut(parent)
	} else {
		h.moveChildren2Root(p)
		h.cut(p)
		h.cascadingCut(parent)
	}
}

func (h *FibHeap) increaseValue(p *node) {
	parent := p.parent
	if h.t == MinHeap {
		h.moveChildren2Root(p)
		h.cut(p)
		h.cascadingCut(parent)
	} else {
		if parent == nil {
			if h.compare(p.v, h.main.v) {
				h.main = p
				return
			}
		}

		if !h.compare(p.v, parent.v) {
			return
		}

		h.cut(p)
		if h.compare(p.v, h.main.v) {
			h.main = p
			return
		}
		h.cascadingCut(parent)
	}
}

func (h *FibHeap) cut(n *node) {
	if n.parent == nil {
		return
	}

	n.marked = false
	if n.right == n {
		n.parent.degree = 0
		n.parent.child = nil
	} else {
		n.parent.degree--
		n.parent.child = n.right
		n.left.right = n.right
		n.right.left = n.left
	}
	n.parent = nil
	h.insertNode(n)
}

func (h *FibHeap) cascadingCut(parent *node) {
	if parent == nil {
		return
	}

	if !parent.marked {
		parent.marked = true
		return
	}

	parent.marked = false
	parent.left.right = parent.right
	parent.right.left = parent.left
	h.insertNode(parent)

	h.cascadingCut(parent.parent)
}

// Delete 删除元素
// 	先将元素的排序指标更新为无穷大/无穷小，然后 pop 出堆顶元素
func (h *FibHeap) Delete(key interface{}) Interface {
	h.mux.Lock()
	if node, exists := h.index[key]; exists {
		h.deleteNode(node)
		h.mux.Unlock()
		return node.value
	}
	h.mux.Unlock()
	return nil
}

func (h *FibHeap) deleteNode(n *node) {
	if h.t == MinHeap {
		n.v = math.Inf(-1)
		h.decreaseValue(n)
	} else {
		n.v = math.Inf(1)
		h.increaseValue(n)
	}
	h.popNode()
}
