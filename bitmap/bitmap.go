package bitmap

import (
	"math"
	"sync"
)

// BitMap .
type BitMap struct {
	mux   *sync.RWMutex
	size  int
	array [math.MaxUint32 / 8]byte
}

// NewBitMap .
func NewBitMap() *BitMap {
	return &BitMap{mux: new(sync.RWMutex)}
}

// Put 将key记录在 bitmap 中
func (b *BitMap) Put(key uint32) bool {
	index := int(key / 8)
	position := key % 8
	b.mux.Lock()
	defer b.mux.Unlock()

	if b.array[index]>>position&1 == 1 {
		return false
	}

	b.array[index] |= 1 << position
	b.size++
	return true
}

// Exists 判断 key 是否存在与 bitmap
func (b *BitMap) Exists(key uint32) bool {
	index := int(key / 8)
	position := key % 8
	b.mux.RLock()
	defer b.mux.RUnlock()
	return b.array[index]>>position&1 == 1
}

// Pop 从 bitmap 中删除某 key
func (b *BitMap) Pop(key uint32) bool {
	index := int(key / 8)
	position := key % 8
	b.mux.Lock()
	defer b.mux.Unlock()

	if b.array[index]>>position&1 != 1 {
		return false
	}

	b.array[index] ^= 1 << position
	b.size--
	return true
}

// Size 返回 bitmap 已使用的大小
func (b *BitMap) Size() int {
	b.mux.RLock()
	defer b.mux.RUnlock()
	return b.size
}
