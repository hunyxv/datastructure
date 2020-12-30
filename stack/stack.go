package stack

import (
	"errors"
	"sync"
)

var (
	// ErrEmpty stack is empty
	ErrEmpty = errors.New("stack is empty")
	// ErrFull stack is full
	ErrFull = errors.New("stack is full")
)

// Stack stack
type Stack struct {
	stack []interface{}
	base  int
	top   int
	size  int
	mux   *sync.RWMutex
}

// NewStack create new stack
func NewStack(size int) *Stack {
	return &Stack{
		stack: make([]interface{}, size),
		size:  size,
		top:   -1,
		mux:   new(sync.RWMutex),
	}
}

// ClearStack clear the stack
func (s *Stack) ClearStack() {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.stack = make([]interface{}, s.size)
}

// IsEmpty return true if stack is empty
func (s *Stack) IsEmpty() bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.top == -1
}

// StackLength returns the number of elements
func (s *Stack) StackLength() int {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.top + 1
}

// GetTop return the top element of the stack if
// the stack is not empty
func (s *Stack) GetTop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmpty
	}
	s.mux.RLock()
	defer s.mux.RUnlock()
	if s.top == -1 {
		return nil, ErrEmpty
	}
	return s.stack[s.top], nil
}

// Push insert element into the top of the stack
func (s *Stack) Push(ele interface{}) error {
	if s.top == s.size {
		return ErrFull
	}

	s.mux.Lock()
	defer s.mux.Unlock()
	if s.top == s.size {
		return ErrFull
	}
	s.top++
	s.stack[s.top] = ele
	return nil
}

// Pop delete and return the element at the top of the stack
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmpty
	}

	s.mux.Lock()
	defer s.mux.Unlock()
	if s.top == -1 {
		return nil, ErrEmpty
	}
	el := s.stack[s.top]
	s.stack[s.top] = nil
	s.top--
	return el, nil
}

// StackTraverse calls f sequentially for element present in the stack. If f
// returns false, StackTraverse stops the iteration.
func (s *Stack) StackTraverse(f func(el interface{}) bool) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	for i := 0; i <= s.top; i++ {
		if !f(s.stack[i]) {
			break
		}
	}
}
