package traverse

import (
	"github.com/structures-algorithms/structures/graph"
	"golang.org/x/exp/slices"
	"math"
)

// BreadthFirstSearchGraph is implementation Breadth first search for Graph structure
// Graph : {map[1:[2 3] 2:[1 4] 3:[1 4] 4:[2 3 5] 5:[4 6] 6:[5]] 6}
// Search Result [1 2 3 4 5 6]
func BreadthFirstSearchGraph(g *graph.Graph[int], start int) []int {
	var listResult []int
	if !g.HasKey(start) {
		return listResult
	}
	currentValue := start
	var queue []int

	queue = append(queue, currentValue)
	for {
		if len(queue) == 0 {
			break
		}

		currentValue, queue = queue[0], queue[1:]
		if slices.Contains(listResult, currentValue) {
			continue
		}

		listResult = append(listResult, currentValue)
		queue = append(queue, g.GetSiblingValues(currentValue)...)
	}

	return listResult
}

// DfsInOrderGraph is implementation Deep first search for Graph structure
// DfsInOrderGraph is use iterative method for excluded stack overflow problem
// Graph : {map[1:[2 3] 2:[1 4] 3:[1 4] 4:[2 3 5] 5:[4 6] 6:[5]] 6}
// Search result : [1 3 4 5 6 2]
func DfsInOrderGraph(g *graph.Graph[int], start int) []int {
	var list []int
	var stack []int
	if !g.HasKey(start) {
		return list
	}
	currentNode := start
	stack = append(stack, start)
	for {
		if stack == nil {
			break
		}

		currentNode, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if slices.Contains(list, currentNode) {
			break
		}
		list = append(list, currentNode)

		for _, edge := range g.GetSiblingEdges(currentNode) {
			if !slices.Contains(list, edge.GetValue()) {
				stack = append(stack, edge.GetValue())
			}
		}
	}

	return list
}

func getModifiedWeightMap[T comparable](weightMap map[T]int, edge *graph.Edge[T], currentValue T) map[T]int {
	inf := int(math.Inf(1))
	weightCurrent := edge.GetWeight()
	if weightMap[currentValue] != inf {
		weightCurrent = weightMap[currentValue] + edge.GetWeight()
	}
	if weightMap[edge.GetValue()] != inf && weightCurrent < weightMap[edge.GetValue()] {
		weightMap[edge.GetValue()] = weightCurrent
	} else if weightMap[edge.GetValue()] == inf {
		weightMap[edge.GetValue()] = weightCurrent
	}

	return weightMap
}

// DijkstraSearch is algorithm for search min weight path from one node to another in direction positive weight graph
// Graph: {map[0:[{1 0} {2 2} {3 4}] 1:[{2 1} {3 3}] 2:[{1 10} {3 0}] 3:[{0 0}]] 4}
// Result 1
func DijkstraSearch[T comparable](g *graph.Graph[T], start T, end T) int {
	inf := int(math.Inf(1))
	if !g.HasKey(start) {
		return inf
	}
	var queue []T
	currentValue := start
	visited := make(map[T]bool, g.GetNodeCount())
	weightMap := make(map[T]int, g.GetNodeCount())

	for node, _ := range g.GetAdjacencyMap() {
		weightMap[node] = inf
	}
	queue = append(queue, currentValue)

	for {
		if len(queue) == 0 {
			break
		}

		currentValue, queue = queue[0], queue[1:]
		if visited[currentValue] {
			continue
		}

		visited[currentValue] = true
		for _, edge := range g.GetSiblingEdges(currentValue) {
			if visited[edge.GetValue()] {
				continue
			}
			weightMap = getModifiedWeightMap(weightMap, &edge, currentValue)
			queue = append(queue, edge.GetValue())
		}
	}

	return weightMap[end]
}
