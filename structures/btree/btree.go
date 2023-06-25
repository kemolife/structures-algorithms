package btree

import (
	"fmt"
	"strings"
)

type T interface{}
type Tree interface{}

type Node[T any] struct {
	left  *Node[T]
	right *Node[T]
	value T
}

type BinaryTree[T any] struct {
	h   uint
	top *Node[T]
}

func less(t1, t2 T) bool {
	switch v1 := t1.(type) {
	case string:
		return strings.Compare(v1, t2.(string)) < 0
	case int:
		return v1 < t2.(int)
	}

	return false
}

func equal(t1, t2 T) bool {
	switch v1 := t1.(type) {
	case string:
		return strings.Compare(v1, t2.(string)) == 0
	case int:
		return v1 == t2.(int)
	}

	return false
}

func addNew[T any](tNode *Node[T], node *Node[T]) {
	if less(node.value, tNode.value) {
		if tNode.left == nil {
			tNode.left = node
			return
		}
		addNew(tNode.left, node)
	} else {
		if tNode.right == nil {
			tNode.right = node
			return
		}
		addNew(tNode.right, node)
	}
}

func (bt *Node[T]) GetLeft() *Node[T] {
	return bt.left
}

func (bt *Node[T]) GetRight() *Node[T] {
	return bt.right
}

func (bt *Node[T]) GetValue() (v T) {
	return bt.value
}

func (bt *BinaryTree[T]) GetTop() *Node[T] {
	return bt.top
}

func (bt *BinaryTree[T]) InsertRecursive(value T) {
	node := &Node[T]{value: value}
	if bt.top == nil {
		bt.top = node
	} else {
		addNew(bt.top, node)
	}
}

func (bt *BinaryTree[T]) Insert(value T) {
	node := &Node[T]{value: value}
	if bt.top == nil {
		bt.top = node
	} else {
		tNode := bt.top
		for {
			if less(node.value, tNode.value) {
				if tNode.left == nil {
					tNode.left = node
					break
				}
				tNode = tNode.left
			} else {
				if tNode.right == nil {
					tNode.right = node
					break
				}
				tNode = tNode.right
			}
		}
	}
}

func (bt *BinaryTree[T]) SearchElement(value T) bool {
	response := true
	if bt.top == nil {
		response = false
		return response
	}
	tNode := bt.top
	for {
		if tNode == nil {
			response = false
			break
		}
		if equal(value, tNode.value) {
			break
		}
		if less(value, tNode.value) {
			tNode = tNode.left
		} else {
			tNode = tNode.right
		}
	}

	return response
}

func (bt *BinaryTree[T]) Delete(value T) {
	if bt.top == nil {
		return
	}
	var prevNode *Node[T]
	tNode := bt.top
	for {
		if tNode == nil {
			break
		}
		if equal(value, tNode.value) {
			if tNode.left != nil && tNode.right != nil {
				right := tNode.right
				left := tNode.left
				if left.right != nil {
					prevNode.left = left.right
				} else {
					prevNode.left = left
					left.right = right
				}
			} else if tNode.left == nil && tNode.right != nil {
				prevNode.left = tNode.right
			} else if tNode.left != nil && tNode.right == nil {
				prevNode.left = tNode.left
			} else {
				prevNode.left = nil
			}
			break
		}
		prevNode = tNode
		if less(value, tNode.value) {
			tNode = tNode.left
		} else {
			tNode = tNode.right
		}
	}
}

func main() {
	bt := &BinaryTree[int]{}
	fmt.Println(bt.h)
	bt.Insert(100)
	bt.Insert(50)
	bt.Insert(60)
	bt.Insert(40)
	bt.Insert(51)
	bt.Insert(55)
	bt.Insert(70)
	bt.Insert(70)
	bt.Insert(120)
	bt.Insert(110)
	bt.Insert(130)
	fmt.Println(bt.top)
	fmt.Println(bt.SearchElement(40))
}
