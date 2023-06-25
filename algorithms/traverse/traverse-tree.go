package traverse

import "github.com/structures-algorithms/structures/btree"

func traverseInOrder[T comparable](node *btree.Node[T], list *[]T) *[]T {
	if node.GetLeft() != nil {
		traverseInOrder(node.GetLeft(), list)
	}

	*list = append(*list, node.GetValue())
	if node.GetRight() != nil {
		traverseInOrder(node.GetRight(), list)
	}
	return list
}

func traversePreOrder[T comparable](node *btree.Node[T], list *[]T) *[]T {
	*list = append(*list, node.GetValue())
	if node.GetLeft() != nil {
		traversePreOrder(node.GetLeft(), list)
	}

	if node.GetRight() != nil {
		traversePreOrder(node.GetRight(), list)
	}
	return list
}

func traversePostOrder[T comparable](node *btree.Node[T], list *[]T) *[]T {
	if node.GetLeft() != nil {
		traversePostOrder(node.GetLeft(), list)
	}

	if node.GetRight() != nil {
		traversePostOrder(node.GetRight(), list)
	}

	*list = append(*list, node.GetValue())
	return list
}

// BinaryTree
//			100
//		50		120
//	  30  55 110   130

// BreadthFirstSearch [100, 50, 120, 30, 55, 110, 130]
func BreadthFirstSearch(t *btree.BinaryTree[int]) []int {
	currentNode := t.GetTop()
	var listResult []int
	var queue []*btree.Node[int]

	queue = append(queue, currentNode)
	for {
		if len(queue) == 0 {
			break
		}

		currentNode, queue = queue[0], queue[1:]
		listResult = append(listResult, currentNode.GetValue())

		if currentNode.GetLeft() != nil {
			queue = append(queue, currentNode.GetLeft())
		}

		if currentNode.GetRight() != nil {
			queue = append(queue, currentNode.GetRight())
		}
	}

	return listResult
}

// BinaryTree
//			100
//		50		120
//	  30  55 110   130

// DeepFirstSearch in order list [30, 50, 55, 100, 110, 120, 130]
func DfsInOrder(t *btree.BinaryTree[int]) *[]int {
	node := t.GetTop()
	var list []int
	return traverseInOrder(node, &list)
}

// BinaryTree
//			100
//		50		120
//	  30  55 110   130

// DeepFirstSearch pre order list [100, 50, 30, 55, 120, 110, 130]
func DfsPreOrder(t *btree.BinaryTree[int]) *[]int {
	node := t.GetTop()
	var list []int
	return traversePreOrder(node, &list)
}

// BinaryTree
//			100
//		50		120
//	  30  55 110   130

// DeepFirstSearch post order list [30, 55, 50, 110, 130, 120, 100]
func DfsPostOrder(t *btree.BinaryTree[int]) *[]int {
	node := t.GetTop()
	var list []int
	return traversePostOrder(node, &list)
}
