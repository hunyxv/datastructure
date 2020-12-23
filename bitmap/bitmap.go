package bitmap

import (
	"math"
	"sync"
)

// BitMap .
type BitMap struct {
	mux   *sync.RWMutex
	size  int
	array []byte
}

// NewBitMap .
func NewBitMap() *BitMap {
	return &BitMap{mux: new(sync.RWMutex), array: make([]byte, 0)}
}

// Put 将key记录在 bitmap 中
func (b *BitMap) Put(key uint32) bool {
	index := int(key / 8)
	position := key % 8
	b.mux.Lock()
	defer b.mux.Unlock()

	if index >= len(b.array) {
		b.grow(index)
		b.array[index] |= 1 << position
		b.size++
		return true
	}

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

	if index >= len(b.array) || b.array[index]>>position&1 != 1 {
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

func (b *BitMap) grow(index int) {
	if index < 1024 {
		b.array = append(b.array, make([]byte, 2*index-len(b.array)+1)...)
		return
	}
	newcap := uint(float32(index) * 1.25)
	if newcap > math.MaxUint32/8 {
		newcap = math.MaxUint32 / 8
	}
	b.array = append(b.array, make([]byte, int(newcap)-len(b.array)+1)...)
}
