package set

import (
	"errors"
	"sync"

	"github.com/hunyxv/datastructure/bitmap"
)

// Interface 实现hash 返回值作为bitmap的key
// 可使用 go/src/runtime/alg.go 、go/src/runtime/hash32.go 中的 memhash 算法
type Interface interface {
	Hash() uint32
}

// BitSet 基于 bitmap 实现的 set
type BitSet struct {
	set    []Interface
	bitmap *bitmap.BitMap
	mux    *sync.RWMutex
}

// NewBitSet return new bitset
func NewBitSet() *BitSet {
	return &BitSet{
		set:    make([]Interface, 0),
		bitmap: bitmap.NewBitMap(),
		mux:    new(sync.RWMutex),
	}
}

// Add Add an element to a set.
func (s *BitSet) Add(el Interface) {
	if s.bitmap.Put(el.Hash()) {
		s.mux.Lock()
		s.set = append(s.set, el)
		s.mux.Unlock()
	}
}

// AddFromList .
func (s *BitSet) AddFromList(els []Interface) {
	for _, el := range els {
		s.Add(el)
	}
}

// Exists .
func (s *BitSet) Exists(el Interface) bool {
	return s.bitmap.Exists(el.Hash())
}

// Set returns all elements in set
func (s *BitSet) Set() []Interface {
	s.mux.RLock()
	defer s.mux.RUnlock()
	set := make([]Interface, len(s.set))
	copy(set, s.set)
	return set
}

// Remove remove an element from a set; it must be a member.
// if the element is not a member, return a error.
func (s *BitSet) Remove(el Interface) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	if !s.bitmap.Exists(el.Hash()) {
		return errors.New("the element is not a member")
	}

	for i, e := range s.set {
		if e == el {
			s.bitmap.Pop(e.Hash())

			s.set[i] = s.set[len(s.set)-1]
			s.set[len(s.set)-1] = nil
			s.set = s.set[:len(s.set)-1]
			return nil
		}
	}
	return nil
}

// Pop remove and return an arbitrary set element.
// return error if the set is empty.
func (s *BitSet) Pop() (Interface, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if len(s.set) == 0 {
		return nil, errors.New("set is empty")
	}
	el := s.set[len(s.set)-1]
	s.set[len(s.set)-1] = nil
	s.set = s.set[:len(s.set)-1]
	s.bitmap.Pop(el.Hash())
	return el, nil
}

// Size returns the number of elements in the set.
func (s *BitSet) Size() int {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return len(s.set)
}

// Difference return the difference of two sets as a new set.
func (s *BitSet) Difference(other *BitSet) *BitSet {
	set := NewBitSet()

	s.mux.RLock()
	defer s.mux.RUnlock()
	for _, el := range s.set {
		if !other.Exists(el) {
			set.Add(el)
		}
	}
	return set
}

// Intersection return the intersection of two sets as a new set.
func (s *BitSet) Intersection(other *BitSet) *BitSet {
	set := NewBitSet()

	s.mux.RLock()
	defer s.mux.RUnlock()
	for _, el := range s.set {
		if other.Exists(el) {
			set.Add(el)
		}
	}

	return set
}

// Union return the union of sets as a new set.
func (s *BitSet) Union(other *BitSet) *BitSet {
	set := NewBitSet()
	set.AddFromList(other.Set())
	s.mux.RLock()
	defer s.mux.RUnlock()
	for _, el := range s.set {
		set.Add(el)
	}
	return set
}

// SysmmetricDifference return the symmetric difference of two sets as a new set.
func (s *BitSet) SysmmetricDifference(other *BitSet) *BitSet {
	set := NewBitSet()
	set.AddFromList(other.Set())

	s.mux.RLock()
	defer s.mux.RUnlock()
	for _, el := range s.set {
		if !other.Exists(el) {
			set.Add(el)
		} else {
			set.Remove(el)
		}
	}
	return set
}

// IsSubSet report whether another set contains this set
func (s *BitSet) IsSubSet(other *BitSet) bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	for _, el := range s.set {
		if !other.Exists(el) {
			return false
		}
	}
	return true
}

// IsSuperSet report whether this set contains another set.
func (s *BitSet) IsSuperSet(other *BitSet) bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	for _, el := range other.Set() {
		if !s.Exists(el) {
			return false
		}
	}

	return true
}
