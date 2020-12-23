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

func TestBitMapPop(t *testing.T) {
	list := []int{3, 7, 4, 9, 1, 0}

	bitmap := NewBitMap()
	for _, v := range list {
		bitmap.Put(uint32(v))
	}

	if !bitmap.Pop(0) {
		t.Fatal("0 should exist")
		return
	}

	if !bitmap.Put(0) {
		t.Fatal("0 should not exist")
	}

	for _, v := range []int{3, 7, 4, 9, 1} {
		if bitmap.Put(uint32(v)) {
			t.Fatalf("%d should exist", v)
		}
	}
}
