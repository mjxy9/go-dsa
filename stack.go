package main

import "fmt"

type Stack struct {
	Rear int
	Size int
	Elements []int
}

func NewStack(size int) *Stack {
	
	return &Stack{
		Size: size,
		Rear: 0,
		Elements: make([]int, size),
	}
}

func (s *Stack) Push(element int) {
	
	if s.IsFull() {
		panic("stack is full")
	}

	s.Elements[s.Rear] = element
	s.Rear += 1
}


func (s *Stack) Pop() {
	
	if s.IsEmpty() {
		panic("stack is empty")
	}

	s.Rear -= 1
	s.Elements[s.Rear] = 0
}

func (s *Stack) IsFull() bool {
	
	return s.Rear == s.Size
}

func (s *Stack) IsEmpty() bool {
	return s.Rear == 0
}

func (s *Stack) Peek() int {
	
	if s.IsEmpty() {
		panic("stack is empty")
	}

	return s.Elements[s.Rear]
}

func (s *Stack) Print() {
	
	for _, e := range s.Elements {
		fmt.Printf("[%d]", e)
	}
}


