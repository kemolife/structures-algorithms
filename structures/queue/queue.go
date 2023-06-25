package queue

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Queue[T any] struct {
	first  *Node[T]
	last   *Node[T]
	length uint
}

// Add element to Queue to end
func (q *Queue[T]) Enqueue(v T) {
	node := &Node[T]{Value: v}
	if q.length == 0 {
		q.first = node
		q.last = node
	} else {
		q.last.Next = node
		q.last = node
	}

	q.length++
}

// Remove first element from Queue
func (q *Queue[T]) Dequeue() {
	if q.length == 0 {
		return
	}

	if q.first == q.last {
		q.first = nil
	}

	q.first = q.first.Next
	q.length--
}

// Get peek value from Queue and move next to top
func (q *Queue[T]) Read() *Node[T] {
	element := q.Peek()
	q.Dequeue()

	return element
}

// Get peek value from Queue
func (q *Queue[T]) Peek() (node *Node[T]) {
	if q.length == 0 {
		return nil
	}

	return q.first
}

// Check if Queue is empty
func (q *Queue[T]) IsEmpty() bool {
	if q.length == 0 {
		return true
	}

	return false
}

func main() {
	q := &Queue[int]{}
	fmt.Println(q.Peek())
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.last)
}
