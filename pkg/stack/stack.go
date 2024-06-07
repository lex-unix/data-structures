package stack

import "errors"

type Stack struct {
	top      int
	size     int
	elements []int
}

func NewStack(size int) *Stack {
	return &Stack{
		top:      -1,
		size:     size,
		elements: make([]int, size),
	}
}

func (s *Stack) Push(elem int) error {
	if s.top == s.size-1 {
		return errors.New("stack overflow")
	}
	s.top += 1
	s.elements[s.top] = elem
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.top == -1 {
		return -1, errors.New("stack underflow")
	}
	x := s.elements[s.top]
	s.top -= 1
	return x, nil
}
