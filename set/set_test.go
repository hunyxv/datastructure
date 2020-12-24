package set

import "testing"

type Value uint32

func (v Value) Hash() uint32 {
	return uint32(v)
}

func TestBitmap(t *testing.T) {
	set := NewBitSet()

	list := []Value{1, 2, 3, 4, 5, 5, 6, 8, 98, 4, 3, 54, 5, 2, 2, 4, 56, 2, 2, 5, 65, 6, 7}

	for _, v := range list {
		set.Add(v)
	}

	t.Logf("%+v", set.Set())
	set.Remove(Value(54))
	t.Logf("remove 7 --> %+v", set.Set())
}

func TestBitmapRemove(t *testing.T) {
	set := NewBitSet()

	if err := set.Remove(Value(1)); err == nil {
		t.FailNow()
	}
}

func TestBitmapPop(t *testing.T) {
	set := NewBitSet()

	if _, err := set.Pop(); err == nil {
		t.Fatal("set is empty")
	}

	list := []Value{1, 2, 3, 4, 5, 5, 6, 8, 98, 4, 3, 54, 5, 2, 2, 4, 56, 2, 2, 5, 65, 6, 7}
	for _, v := range list {
		set.Add(v)
	}

	v, err := set.Pop()
	if err != nil {
		t.Fatal("set is not empty")
	}
	t.Logf("pop --> %+v", v.(Value))
}

func TestBitmapDifference(t *testing.T) {
	set := NewBitSet()
	list := []Value{1, 2, 3, 4, 5, 8, 0}
	for _, v := range list {
		set.Add(v)
	}

	set2 := NewBitSet()
	list2 := []Value{8, 0, 7, 6, 9}
	for _, v := range list2 {
		set2.Add(v)
	}

	diff := set.Difference(set2)
	t.Logf("%+v", diff.Set())
}

func TestBitmapIntersection(t *testing.T) {
	set := NewBitSet()
	list := []Value{1, 2, 3, 5, 8, 0}
	for _, v := range list {
		set.Add(v)
	}

	set2 := NewBitSet()
	list2 := []Value{8, 0, 7, 6, 9}
	for _, v := range list2 {
		set2.Add(v)
	}

	intersection := set.Intersection(set2)
	t.Logf("%+v", intersection.Set())
}

func TestBitmapUnion(t *testing.T) {
	set := NewBitSet()
	list := []Value{1, 2, 3, 5, 8, 0}
	for _, v := range list {
		set.Add(v)
	}

	set2 := NewBitSet()
	list2 := []Value{8, 0, 7, 6, 9}
	for _, v := range list2 {
		set2.Add(v)
	}

	intersection := set.Union(set2)
	t.Logf("%+v", intersection.Set())
}

func TestBitmapIsSubSet(t *testing.T) {
	set := NewBitSet()
	list := []Value{1, 2, 3, 5, 8, 0}
	for _, v := range list {
		set.Add(v)
	}

	set2 := NewBitSet()
	list2 := []Value{1, 2, 3, 5, 8, 0}
	for _, v := range list2 {
		set2.Add(v)
	}

	if !set.IsSubSet(set2) {
		t.FailNow()
	}
}

func TestBitmapIsSuperSet(t *testing.T) {
	set := NewBitSet()
	list := []Value{1, 2, 3, 5, 6}
	for _, v := range list {
		set.Add(v)
	}

	set2 := NewBitSet()
	list2 := []Value{1, 2, 3}
	for _, v := range list2 {
		set2.Add(v)
	}

	if !set.IsSuperSet(set2) {
		t.FailNow()
	}
}
