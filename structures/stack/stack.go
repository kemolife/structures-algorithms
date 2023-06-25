package stack

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Stack[T any] struct {
	top    *Node[T]
	length uint
}

func (s *Stack[T]) Push(v T) {
	node := &Node[T]{Value: v}
	if s.length == 0 {
		s.top = node
	} else {
		node.Next = s.top
		s.top = node
	}

	s.length++
}

func (s *Stack[T]) Pop() {
	if s.length == 0 {
		return
	}

	s.length--
	if s.length == 0 {
		s.top = nil
		return
	}

	s.top = s.top.Next
}

func (s *Stack[T]) Peek() (v *T) {
	if s.length == 0 {
		return nil
	}

	return &s.top.Value
}

func (s *Stack[T]) IsEmpty() bool {
	if s.length == 0 {
		return true
	}

	return false
}

func main() {
	s := &Stack[int]{}
	fmt.Println(s.Peek())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s.IsEmpty())
	s.Pop()
	s.Pop()
	s.Pop()
	fmt.Println(s.top)
	fmt.Println(s.length)
	s.Pop()
	fmt.Println(s.top)
	fmt.Println(s.length)
	fmt.Println(s.Peek())
	fmt.Println(s.IsEmpty())
}
