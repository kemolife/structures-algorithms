package graph

type Edge[T comparable] struct {
	value  T
	weight int
}

type Graph[T comparable] struct {
	adjacencyMap map[T][]Edge[T]
	nodesCount   uint
}

func (e *Edge[T]) GetValue() T {
	return e.value
}

func (e *Edge[T]) GetWeight() int {
	return e.weight
}

func (gr *Graph[T]) GetNodeCount() uint {
	return gr.nodesCount
}

func (gr *Graph[T]) GetAdjacencyMap() map[T][]Edge[T] {
	return gr.adjacencyMap
}

func (gr *Graph[T]) GetSiblingValues(key T) []T {
	siblingEdges := gr.adjacencyMap[key]
	siblings := make([]T, len(siblingEdges))
	for key, edge := range siblingEdges {
		siblings[key] = edge.value
	}

	return siblings
}

func (gr *Graph[T]) GetSiblingEdges(key T) []Edge[T] {
	return gr.adjacencyMap[key]
}

func (gr *Graph[T]) HasKey(key T) bool {
	_, ok := gr.adjacencyMap[key]

	return ok
}

func (gr *Graph[T]) AddEdge(sourceNode T, targetNode T, weight int) {
	edgeSource := Edge[T]{targetNode, weight}

	if len(gr.adjacencyMap) == 0 {
		gr.adjacencyMap = make(map[T][]Edge[T])
	}

	if _, ok := gr.adjacencyMap[sourceNode]; !ok {
		gr.adjacencyMap[sourceNode] = []Edge[T]{edgeSource}
		gr.nodesCount++
	} else {
		gr.adjacencyMap[sourceNode] = append(gr.adjacencyMap[sourceNode], edgeSource)
	}

}
