# structures-algorithms
Implementing all basic structures and algorithm and solving leetcode challenges used Golang.

```
firstBadVersion(5)
var slice []string
array := structures.Array[string]{Data: slice}
array.Add("a")
array.Add("b", "c", "d", "a")
array.Update(2, "x")
array.Add(1)
array.Pop(1)
errR := array.Remove("a")
if errR != nil {
    fmt.Println(errR)
}
i, _ := array.Search("c")
fmt.Println(i)
value, err := array.Get(1)
if err != nil {
    fmt.Println(err)
}
fmt.Println(value)
array.Clear()
fmt.Println(array)
s := "T"
fmt.Println(revers(s))
hm := structures.New[string, string](100)
hm.Set("test", "100")
hm.Set("test1", "120")
println(hm.Data)
println(hm.Get("test"))  //"100"
println(hm.Get("test1")) //"100"

s := []int{1, 2, 3, 4, 6, 55, 34, 7, 2}
response, err := recurringCharacters(s)
//
if err != nil {
    fmt.Println(err)
}
//
fmt.Println(response)

//linked list
//head->->->tail
//10->5->1->nil
n1 := &structures.Node{Value: nil, Node: nil}
n2 := &structures.Node{Value: 1, Node: n1}
n3 := &structures.Node{Value: 5, Node: n2}
ll := structures.LinkedList{Head: 10, Node: n3}
//
ll := structures.NewLinkList[int]()
ll.Push(1)
ll.Push(2)
ll.Push(3)
ll.Push(4)
ll.Reverse()
fmt.Println(ll.ToList())
//
bt := btree.BinaryTree[int]{}
bt.Insert(100)
bt.Insert(50)
bt.Insert(30)
bt.Insert(55)
bt.Insert(120)
bt.Insert(110)
bt.Insert(130)
//
fmt.Println(search.BreadthFirstSearch(&bt))
fmt.Println(traverse.DfsInOrder(&bt))
fmt.Println(traverse.DfsPreOrder(&bt))
fmt.Println(traverse.DfsPostOrder(&bt))
//
c := &dynamic.Cache{}
fmt.Println(dynamic.Fibonacci(35, c))
fmt.Println(c.Count)
//
g := graph.Graph[int]{}
g.AddEdge(0, 1, 0)
g.AddEdge(0, 2, 2)
g.AddEdge(0, 3, 4)
g.AddEdge(1, 2, 1)
g.AddEdge(1, 3, 3)
g.AddEdge(2, 1, 10)
g.AddEdge(2, 3, 0)
g.AddEdge(3, 0, 0)
//
fmt.Println(traverse.BreadthFirstSearchGraph(&g, 1))
fmt.Println(traverse.DfsInOrderGraph(&g, 1))
fmt.Println(traverse.DijkstraSearch(&g, 0, 3))
fmt.Println(g)
nums := []int{1, 3, 5, 6}
fmt.Println(search(nums, 7))
//
nums := []int{-1, -1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
moveZeroes(twoSumSorted(nums, -2))
//
fmt.Println(convertToTitle(702))
fmt.Println(reverseWords("Let's take LeetCode contest"))
ll := &linked_list.LinkedList[int]{}
ll.Push(1)
//
println(removeNthFromEnd(ll.Head, 1).Value)
s1 := "hello"
s2 := "ooolleoooleh"
fmt.Println(lengthOfLongestSubstring(s))
image := [][]int{{0, 0, 0}, {0, 0, 0}}
fmt.Println(floodFill(image, 0, 0, 0))
grid := [][]int{
	{0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 1, 1, 1, 1},
	{0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 0},
	{0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1},
	{0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 0, 1, 1},
	{1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 0, 0, 1},
}
fmt.Println(maxAreaOfIsland(grid))
root1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
root2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}
mergeTrees(root1, root2)
```