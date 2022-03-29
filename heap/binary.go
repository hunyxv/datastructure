package heap

import (
	"errors"
	"sync"
)

// T 堆类型
type T int

const (
	// MaxHeap 最大堆
	MaxHeap T = iota
	// MinHeap 最小堆
	MinHeap
)

var (
	// ErrEmpty heap is empty
	ErrEmpty = errors.New("heap is empty")
	// ErrExceed index out of range
	ErrExceed = errors.New("index out of range")
)

// Interface 用于两个 Interface 比较
type Interface interface {
	Value() float64
}

// BinaryHeap .
type BinaryHeap struct {
	heap []Interface

	t   T
	mux sync.RWMutex
}

// NewBinaryHeap 创建堆
func NewBinaryHeap(t T) *BinaryHeap {
	return &BinaryHeap{
		heap: make([]Interface, 0),
		t:    t,
	}
}

func (h *BinaryHeap) shiftUp(index int) {
	switch h.t {
	case MaxHeap:
		for h.heap[(index-1)/2].Value() < h.heap[index].Value() {
			h.heap[(index-1)/2], h.heap[index] = h.heap[index], h.heap[(index-1)/2]
			index = (index - 1) / 2
		}
	case MinHeap:
		for h.heap[(index-1)/2].Value() > h.heap[index].Value() {
			h.heap[(index-1)/2], h.heap[index] = h.heap[index], h.heap[(index-1)/2]
			index = (index - 1) / 2
		}
	}
}

func (h *BinaryHeap) shiftDown(index int) {
	left := index*2 + 1
	right := index*2 + 2
	var target int
	for left < len(h.heap) {
		switch h.t {
		case MaxHeap:
			if right < len(h.heap) && ge(h.heap[right], h.heap[left]) {
				target = right
			} else {
				target = left
			}
			if h.heap[index].Value() > h.heap[target].Value() {
				return
			}
		case MinHeap:
			if right < len(h.heap) && le(h.heap[right], h.heap[left]) {
				target = right
			} else {
				target = left
			}
			if h.heap[index].Value() < h.heap[target].Value() {
				return
			}
		}
		h.heap[index], h.heap[target] = h.heap[target], h.heap[index]
		index = target
		left = index*2 + 1
		right = index*2 + 2
	}
}

// Insert 入堆
func (h *BinaryHeap) Insert(val Interface) {
	h.mux.Lock()
	defer h.mux.Unlock()

	h.heap = append(h.heap, val)
	index := len(h.heap) - 1

	h.shiftUp(index)
}

// Pop 返回堆顶元素并删除
func (h *BinaryHeap) Pop() (val Interface, err error) {
	h.mux.Lock()
	defer h.mux.Unlock()

	if len(h.heap) == 0 {
		return nil, ErrEmpty
	}

	val = h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]

	h.shiftDown(0)
	return
}

// Peek 返回堆顶元素不删除
func (h *BinaryHeap) Peek() (val Interface, err error) {
	h.mux.RLock()
	defer h.mux.RUnlock()
	if len(h.heap) == 0 {
		return nil, ErrEmpty
	}
	val = h.heap[0]
	return
}

// PopByIndex 返回 index 索引下的值 并删除
func (h *BinaryHeap) PopByIndex(index int) (val Interface, err error) {
	h.mux.Lock()
	defer h.mux.Unlock()

	if index >= len(h.heap) {
		return nil, ErrExceed
	}

	val = h.heap[index]
	h.heap[index] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.shiftDown(index)
	return
}

// Replace 替换 index 位置的值
func (h *BinaryHeap) Replace(index int, val Interface) error {
	h.mux.Lock()
	defer h.mux.Unlock()

	if index >= len(h.heap) {
		return ErrExceed
	}

	h.heap[index] = val
	h.shiftUp(index)
	h.shiftDown(index)
	return nil
}

// Size .
func (h *BinaryHeap) Size() int {
	h.mux.RLock()
	defer h.mux.RUnlock()
	return len(h.heap)
}

func le(a, b Interface) bool {
	if a.Value() < b.Value() {
		return true
	}
	return false
}

func ge(a, b Interface) bool {
	if a.Value() > b.Value() {
		return true
	}
	return false
}
