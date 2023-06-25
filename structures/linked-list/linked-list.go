package linked_list

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// LinkedList head ->10->5->1->7<- tail
type LinkedList[T any] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length uint
}

func NewLinkList[T any]() *LinkedList[T] {
	ll := &LinkedList[T]{}
	ll.Length = 0

	return ll
}

func (ll *LinkedList[T]) Unshift(value T) {
	node := &Node[T]{Value: value}

	if ll.Length == 0 {
		ll.Head = node
		ll.Tail = node
	} else {
		node.Next = ll.Head
		ll.Head = node
	}

	ll.Length++
}

func (ll *LinkedList[T]) Push(value T) {
	node := &Node[T]{Value: value}

	if ll.Length == 0 {
		ll.Head = node
		ll.Tail = node
	} else {
		ll.Tail.Next = node
		ll.Tail = node
	}

	ll.Length++
}

func (ll *LinkedList[T]) Insert(n uint, value T) {
	node := &Node[T]{Value: value}

	if n > ll.Length {
		ll.Push(value)
		return
	}

	leadNode := ll.getLeadNode(n)
	afterNode := leadNode.Next
	leadNode.Next = node
	node.Next = afterNode
	ll.Length++
}

func (ll *LinkedList[T]) getLeadNode(index uint) *Node[T] {
	leadNode := ll.Head
	var i uint
	for i = 1; i < index; i++ {
		leadNode = leadNode.Next
	}

	return leadNode
}

func (ll *LinkedList[T]) Delete(n uint) {
	if n > ll.Length {
		return
	}

	leadNode := ll.getLeadNode(n)
	deleteNode := leadNode.Next
	leadNode.Next = deleteNode.Next
	ll.Length--
}

func (ll *LinkedList[T]) Reverse() {
	prev := &Node[T]{}
	current := ll.Head
	following := ll.Head

	for {
		if current == nil {
			break
		}

		following = following.Next
		current.Next = prev
		prev = current
		current = following
	}

	ll.Head = prev
}

func (ll *LinkedList[T]) ToList() []T {
	var list []T
	nodeValue := ll.Head
	for {
		list = append(list, nodeValue.Value)
		if nodeValue.Next == nil {
			break
		}
		nodeValue = nodeValue.Next
	}

	return list
}

func main() {
	//head->->->tail
	//head -> 10->5->1->7 <-tail
	//n1 := &Node{Value: nil, Node: nil}
	//n2 := &Node{Value: 1, Node: n1}
	//n3 := &Node{Value: 5, Node: n2}
	//ll := LinkedList[int]{Head: n2, Tail: n3}
	//
	//fmt.Println(ll)
}
