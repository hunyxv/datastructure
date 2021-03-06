package set

import (
	"math/rand"
	"testing"
	"time"
)

type Value uint32

func (v Value) Hash() uint32 {
	return uint32(v)
}

func TestBitSet(t *testing.T) {
	set := NewBitSet()

	list := []Value{1, 2, 3, 4, 5, 5, 6, 8, 98, 4, 3, 54, 5, 2, 2, 4, 56, 2, 2, 5, 65, 6, 7}

	for _, v := range list {
		set.Add(v)
	}

	t.Logf("%+v", set.Set())
	set.Remove(Value(54))
	t.Logf("remove 54 --> %+v", set.Set())
}

func TestBitSetRemove(t *testing.T) {
	set := NewBitSet()

	if err := set.Remove(Value(1)); err == nil {
		t.FailNow()
	}
}

func TestBitSetPop(t *testing.T) {
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

func TestBitSetDifference(t *testing.T) {
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

func TestBitSetIntersection(t *testing.T) {
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

func TestBitSetUnion(t *testing.T) {
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

func TestBitSetIsSubSet(t *testing.T) {
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

func TestBitSetIsSuperSet(t *testing.T) {
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

func TestMapSet(t *testing.T) {
	set := NewMapSet()

	list := []int{1, 2, 3, 4, 5, 5, 6, 8, 98, 4, 3, 54, 5, 2, 2, 4, 56, 2, 2, 5, 65, 6, 7}

	for _, v := range list {
		set.Add(v)
	}

	t.Logf("%+v", set.Set())
	set.Remove(54)
	t.Logf("remove 54 --> %+v", set.Set())
}

func TestMapSetRemove(t *testing.T) {
	set := NewMapSet()

	set.Remove(1)
}

func TestMapSetPop(t *testing.T) {
	set := NewMapSet()

	if _, err := set.Pop(); err == nil {
		t.Fatal("set is empty")
	}

	list := []int{1, 2, 3, 4, 5, 5, 6, 8, 98, 4, 3, 54, 5, 2, 2, 4, 56, 2, 2, 5, 65, 6, 7}
	for _, v := range list {
		set.Add(v)
	}

	v, err := set.Pop()
	if err != nil {
		t.Fatal("set is not empty")
	}
	t.Logf("pop --> %+v", v.(int))
}

func TestMapSetDifference(t *testing.T) {
	set := NewMapSet()
	list := []int{1, 2, 3, 4, 5, 8, 0}
	for _, v := range list {
		set.Add(v)
	}

	set2 := NewMapSet()
	list2 := []int{8, 0, 7, 6, 9}
	for _, v := range list2 {
		set2.Add(v)
	}

	diff := set.Difference(set2)
	t.Logf("%+v", diff.Set())
}

func TestMapSetIntersection(t *testing.T) {
	set := NewMapSet()
	list := []int{1, 2, 3, 5, 8, 0}
	for _, v := range list {
		set.Add(v)
	}

	set2 := NewMapSet()
	list2 := []int{8, 0, 7, 6, 9}
	for _, v := range list2 {
		set2.Add(v)
	}

	intersection := set.Intersection(set2)
	t.Logf("%+v", intersection.Set())
}

func TestMapSetUnion(t *testing.T) {
	set := NewMapSet()
	list := []int{1, 2, 3, 5, 8, 0}
	for _, v := range list {
		set.Add(v)
	}

	set2 := NewMapSet()
	list2 := []int{8, 0, 7, 6, 9}
	for _, v := range list2 {
		set2.Add(v)
	}

	intersection := set.Union(set2)
	t.Logf("%+v", intersection.Set())
}

func TestMapSetIsSubSet(t *testing.T) {
	set := NewMapSet()
	list := []int{1, 2, 3, 5, 8, 0}
	for _, v := range list {
		set.Add(v)
	}

	set2 := NewMapSet()
	list2 := []int{1, 2, 3, 5, 8, 0}
	for _, v := range list2 {
		set2.Add(v)
	}

	if !set.IsSubSet(set2) {
		t.FailNow()
	}
}

func TestMapSetIsSuperSet(t *testing.T) {
	set := NewMapSet()
	list := []int{1, 2, 3, 5, 6}
	for _, v := range list {
		set.Add(v)
	}

	set2 := NewMapSet()
	list2 := []int{1, 2, 3}
	for _, v := range list2 {
		set2.Add(v)
	}

	if !set.IsSuperSet(set2) {
		t.FailNow()
	}
}

func BenchmarkBitSetAdd(b *testing.B) {
	set := NewBitSet()
	b.RunParallel(func(pb *testing.PB) {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		for pb.Next() {
			n := r.Uint32()
			set.Add(Value(n))
		}
	})
}

func BenchmarkMapSetAdd(b *testing.B) {
	set := NewMapSet()
	b.RunParallel(func(pb *testing.PB) {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		for pb.Next() {
			n := r.Uint32()
			set.Add(n)
		}
	})
}
