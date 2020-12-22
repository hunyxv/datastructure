package bitmap

import "testing"

func TestBitMap(t *testing.T) {
	list := []int{16, 1, 2, 3, 4, 5, 6, 7, 2, 3, 7, 1, 4}

	bitmap := NewBitMap()
	for _, v := range list {
		bitmap.Put(uint32(v))
	}

	if bitmap.Size() != 8 || len(bitmap.array) != 5 {
		t.FailNow()
	}
}
