package dlist

import "fmt"

type Node[T any] struct {
	Value    T
	Next     *Node[T]
	Previous *Node[T]
}

// DoublyLinkedList head -> 1 -> <- 0 -> <- 1 -> <- 2 -> 4 <- tail
type DoublyLinkedList[T any] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length uint
}

func New[T any]() *DoublyLinkedList[T] {
	dll := &DoublyLinkedList[T]{}
	dll.Length = 0

	return dll
}

func (dll *DoublyLinkedList[T]) Unshift(value T) {
	node := &Node[T]{Value: value}

	if dll.Length == 0 {
		dll.Head = node
		dll.Tail = node
	} else {
		node.Next = dll.Head
		dll.Head.Previous = node
		dll.Head = node
	}

	dll.Length++
}

func (dll *DoublyLinkedList[T]) Push(value T) {
	node := &Node[T]{Value: value}

	if dll.Length == 0 {
		dll.Head = node
		dll.Tail = node
	} else {
		dll.Tail.Next = node
		node.Previous = dll.Tail
		dll.Tail = node
	}

	dll.Length++
}

func (dll *DoublyLinkedList[T]) Insert(n uint, value T) {
	node := &Node[T]{Value: value}

	if n > dll.Length {
		dll.Push(value)
		return
	}

	leadNode := dll.getLeadNode(n)
	afterNode := leadNode.Next
	leadNode.Next = node
	node.Next = afterNode
	node.Previous = leadNode
	afterNode.Previous = node
	dll.Length++
}

func (dll *DoublyLinkedList[T]) getLeadNode(index uint) *Node[T] {
	leadNode := dll.Head
	var i uint
	for i = 1; i < index; i++ {
		leadNode = leadNode.Next
	}

	return leadNode
}

func (dll *DoublyLinkedList[T]) Delete(n uint) {
	if n > dll.Length {
		return
	}

	leadNode := dll.getLeadNode(n)
	deleteNode := leadNode.Next
	afterNode := deleteNode.Next
	leadNode.Next = afterNode
	afterNode.Previous = leadNode
	dll.Length--
}
func (dll *DoublyLinkedList[T]) ToList() []T {
	var list []T
	node := dll.Head
	for {
		list = append(list, node.Value)
		if node.Next == nil {
			break
		}
		node = node.Next
	}

	return list
}

func (dll *DoublyLinkedList[T]) ToRevertList() []T {
	var list []T
	node := dll.Tail
	for {
		list = append(list, node.Value)
		if node.Previous == nil {
			break
		}
		node = node.Previous
	}

	return list
}

func main() {
	dll := New[int]()
	dll.Push(1)
	dll.Push(2)
	dll.Push(4)
	dll.Unshift(0)
	dll.Unshift(-1)
	dll.Insert(4, 3)
	dll.Push(5)
	dll.Unshift(-2)
	dll.Delete(2)
	fmt.Println(dll.ToList())
	fmt.Println(dll.ToRevertList())
}
