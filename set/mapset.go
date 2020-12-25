package set

import (
	"errors"
	"sync"
)

// MapSet .
type MapSet struct {
	sync.Map
}

// NewMapSet return new bitset
func NewMapSet() *MapSet {
	return &MapSet{}
}

// Add Add an element to a set.
func (s *MapSet) Add(key interface{}) {
	s.Store(key, struct{}{})
}

// AddFromList .
func (s *MapSet) AddFromList(keys []interface{}) {
	for _, k := range keys {
		s.Store(k, struct{}{})
	}
}

// Exists .
func (s *MapSet) Exists(key interface{}) bool {
	if _, ok := s.Load(key); ok {
		return true
	}
	return false
}

// Set returns all elements in set
func (s *MapSet) Set() []interface{} {
	set := make([]interface{}, 0)
	s.Range(func(key, _ interface{}) bool {
		set = append(set, key)
		return true
	})
	return set
}

// Remove remove an element from a set; it must be a member.
// if the element is not a member, return a error.
func (s *MapSet) Remove(key interface{}) {
	s.Delete(key)
	return
}

// Pop remove and return an arbitrary set element.
// return error if the set is empty.
func (s *MapSet) Pop() (interface{}, error) {
	var ele interface{}

	s.Range(func(key, _ interface{}) bool {
		ele = key
		return false
	})
	if ele == nil {
		return nil, errors.New("set is empty")
	}
	s.Delete(ele)
	return ele, nil
}

// Size returns the number of elements in the set.
func (s *MapSet) Size() int {
	var size int
	s.Range(func(_, _ interface{}) bool {
		size++
		return true
	})
	return size
}

// Difference return the difference of two sets as a new set.
func (s *MapSet) Difference(other *MapSet) *MapSet {
	set := NewMapSet()

	s.Range(func(el, _ interface{}) bool {
		if !other.Exists(el) {
			set.Add(el)
		}
		return true
	})
	return set
}

// Intersection return the intersection of two sets as a new set.
func (s *MapSet) Intersection(other *MapSet) *MapSet {
	set := NewMapSet()

	s.Range(func(el, _ interface{}) bool {
		if other.Exists(el) {
			set.Add(el)
		}
		return true
	})
	return set
}

// Union return the union of sets as a new set.
func (s *MapSet) Union(other *MapSet) *MapSet {
	set := NewMapSet()

	set.AddFromList(other.Set())
	s.Range(func(el, _ interface{}) bool {
		if !other.Exists(el) {
			set.Add(el)
		}
		return true
	})

	return set
}

// SysmmetricDifference return the symmetric difference of two sets as a new set.
func (s *MapSet) SysmmetricDifference(other *MapSet) *MapSet {
	set := NewMapSet()
	set.AddFromList(other.Set())

	s.Range(func(el, _ interface{}) bool {
		if !other.Exists(el) {
			set.Add(el)
		} else {
			set.Remove(el)
		}
		return true
	})
	return set
}

// IsSubSet report whether another set contains this set
func (s *MapSet) IsSubSet(other *MapSet) bool {
	var b bool
	s.Range(func(el, _ interface{}) bool {
		if !other.Exists(el) {
			b = false
			return false
		}
		return true
	})

	return b
}

// IsSuperSet report whether this set contains another set.
func (s *MapSet) IsSuperSet(other *MapSet) bool {
	var b bool
	other.Range(func(el, _ interface{}) bool {
		if !s.Exists(el) {
			b = false
			return false
		}
		return true
	})

	return b
}
