package leetcode

import (
	"fmt"
	linked_list "github.com/structures-algorithms/structures/linked-list"
	"strconv"
	"strings"
)

func revers(s string) string {
	if len(s) < 2 {
		return s
	}

	var l []byte
	for i := len(s) - 1; i >= 0; i-- {
		l = append(l, s[i])
	}

	return string(l)
}

//Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums.
//If target exists, then return its index. Otherwise, return -1.
//You must write an algorithm with O(log n) runtime complexity.
//Input: nums = [-1,0,3,5,9,12], target = 9
//Output: 4
//Explanation: 9 exists in nums and its index is 4
func search(nums []int, target int) int {
	n := len(nums)
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if nums[h] < target {
			i = h + 1
		} else {
			j = h
		}
	}

	if i > n {
		return -1
	}

	return i
}

//Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums.
//If target exists, then return its index. Otherwise, return -1.
//You must write an algorithm with O(log n) runtime complexity.
//Input: nums = [-1,0,3,5,9,12], target = 9
//Output: 4
//Explanation: 9 exists in nums and its index is 4
func searchInsert(nums []int, target int) int {
	n := len(nums)
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if nums[h] < target {
			i = h + 1
		} else if nums[h] > target {
			j = h
		} else if nums[h] == target {
			i = h
			break
		}
	}

	return i
}

func isBadVersion(n int) bool {
	if n == 10 {
		return true
	}

	return false
}

//You are a product manager and currently leading a team to develop a new product.
//Unfortunately, the latest version of your product fails the quality check.
//Since each version is developed based on the previous version, all the versions after a bad version are also bad.
//Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.
//You are given an API bool isBadVersion(version) which returns whether version is bad.
//Implement a function to find the first bad version. You should minimize the number of calls to the API.
func firstBadVersion(n int) int {
	low := 1
	height := n
	for {
		if low >= height {
			break
		}

		middle := low + (height-low)/2
		if isBadVersion(middle) {
			height = middle
		} else {
			low = middle + 1
		}
	}

	return low
}

//Given an integer array nums sorted in non-decreasing order,
//return an array of the squares of each number sorted in non-decreasing order.
//Input: nums = [-4,-1,0,3,10]
//Output: [0,1,9,16,100]
//Explanation: After squaring, the array becomes [16,1,0,9,100].
//After sorting, it becomes [0,1,9,16,100].
func sortedSquares(nums []int) []int {
	var negativeList []int
	var i int
	for {
		if i == len(nums) || nums[i] >= 0 {
			break
		}
		negativeList = append(negativeList, nums[i])
		i++
	}

	if len(negativeList) == 0 {
		for j, v := range nums {
			nums[j] = v * v
		}

		return nums
	}

	positiveList := nums[i:]

	if len(positiveList) == 0 {
		negativeListSqr := make([]int, len(negativeList))
		for k := len(negativeList) - 1; k >= 0; k-- {
			i = len(negativeList) - 1 - k
			negativeListSqr[i] = negativeList[k] * negativeList[k]
		}

		return negativeListSqr
	}

	for j, v := range positiveList {
		positiveList[j] = v * v
	}

	negativeListSqr := make([]int, len(negativeList))
	for k := len(negativeList) - 1; k >= 0; k-- {
		i = len(negativeList) - 1 - k
		negativeListSqr[i] = negativeList[k] * negativeList[k]
	}

	var combine []int
	i, j := 0, 0
	for i < len(negativeListSqr) && j < len(positiveList) {
		if negativeListSqr[i] < positiveList[j] {
			combine = append(combine, negativeListSqr[i])
			i++
		} else {
			combine = append(combine, positiveList[j])
			j++
		}
	}
	for ; i < len(negativeListSqr); i++ {
		combine = append(combine, negativeListSqr[i])
	}
	for ; j < len(positiveList); j++ {
		combine = append(combine, positiveList[j])
	}

	return combine
}

//Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
//Input: nums = [1,2,3,4,5,6,7], k = 3
//Output: [5,6,7,1,2,3,4]
//Explanation:
//rotate 1 steps to the right: [7,1,2,3,4,5,6]
//rotate 2 steps to the right: [6,7,1,2,3,4,5]
//rotate 3 steps to the right: [5,6,7,1,2,3,4]
func rotate(nums []int, k int) []int {
	if k == 0 {
		return nums
	}
	var rotate []int
	middle := len(nums) / 2

	right := nums[middle+1:]
	left := nums[:middle+1]

	for i := k; i <= len(right); i++ {
		rotate = append(rotate, right[:i]...)
	}

	return append(rotate, left...)
}

//Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//Without making a copy of the array.
//Input: nums = [0,1,0,3,12]
//Output: [1,3,12,0,0]
func moveZeroes(nums []int) {
	zeroCount := 0
	var i int
	for {
		if i >= len(nums) {
			break
		}

		if nums[i] == 0 {
			if i == len(nums)-1 {
				nums = nums[:i]
			} else {
				nums = append(nums[:i], nums[i+1:]...)
			}
			zeroCount++
		} else {
			i++
		}
	}
	if zeroCount != 0 {
		zeroList := make([]int, zeroCount)
		nums = append(nums, zeroList...)
	}
}

//Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number.
//Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
//Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
//The tests are generated such that there is exactly one solution. You may not use the same element twice.
//Your solution must use only constant extra space.
//Input: numbers = [2,7,11,15], target = 9
//Output: [1,2]
//Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
func twoSum(numbers []int, target int) []int {
	result := make([]int, 2)
out:
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				result[0], result[1] = i, j
				break out
			}
		}
	}

	return result
}

func twoSumSorted(numbers []int, target int) []int {
	result := make([]int, 2)
	n := len(numbers)
	i, j := 0, n

	for i < j {
		a := target - numbers[j-1]

		if a == numbers[i] {
			result[0] = i + 1
			result[1] = j
			break
		}

		if a > numbers[i] {
			i++
			continue
		}

		if a < numbers[i] {
			j--
			continue
		}
	}

	return result
}

//Write a function that reverses a string. The input string is given as an array of characters s.
//You must do this by modifying the input array in-place with O(1) extra memory.
//Input: s = ["h","e","l","l","o"]
//Output: ["o","l","l","e","h"]
func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func convertToTitle(columnNumber int) string {
	toChar := func(i int) rune {
		return rune('A' - 1 + i)
	}
	pos := columnNumber % 26

	if columnNumber <= 26 {
		return string(toChar(pos))
	}

	result := ""
	for {
		if columnNumber <= 26 {
			result += string(toChar(pos))
			break
		} else {
			columnNumber = columnNumber / 26
			if pos == 0 {
				columnNumber = columnNumber - 1
				pos = 26
			}
			result += string(toChar(columnNumber % 26))
		}
	}

	return result
}

func convertToTitleByte(columnNumber int) string {
	bA := byte('A')
	res := make([]byte, 7)
	pos := 6
	for columnNumber > 0 {
		columnNumber--
		res[pos] = bA + byte(columnNumber%26)
		pos--
		columnNumber /= 26
	}
	return string(res[pos+1:])
}

func convertToTitleNew(columnNumber int) string {
	const Base26Digits = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var title string
	var quotient, remainder int
	for ok := true; ok; ok = (quotient != 0) {
		quotient, remainder = columnNumber/26, columnNumber%26
		if remainder == 0 {
			quotient, remainder = quotient-1, 26
		}
		fmt.Println(quotient, remainder)
		title = string(Base26Digits[remainder-1]) + title
		columnNumber = quotient
	}
	return title
}

//Given a string s, reverse the order of characters in each word within a
//sentence while still preserving whitespace and initial word order.
//Input: s = "Let's take LeetCode contest"
//Output: "s'teL ekat edoCteeL tsetnoc"
func reverseWords(s string) string {
	result := ""
	words := strings.Split(s, " ")
	for _, w := range words {
		revers := ""
		for _, l := range w {
			revers = string(l) + revers
		}
		result += revers + " "
	}

	return strings.Trim(result, " ")
}

//Given the head of a singly linked list, return the middle node of the linked list.
//If there are two middle nodes, return the second middle node.
//Input: head = [1,2,3,4,5]
//Output: [3,4,5]
//Explanation: The middle node of the list is node 3.
func middleNode(head *linked_list.Node[int]) *linked_list.Node[int] {
	i := 1
	current := head
	for {
		if current.Next == nil {
			break
		}

		current = current.Next
		i++
	}

	middle := i / 2
	i = 0
	current = head
	for {
		if i == middle {
			break
		}

		current = current.Next
		i++
	}

	return current
}

//Given the head of a linked list, remove the nth node from the end of the list and return its head.
//Input: head = [1,2,3,4,5], n = 2
//Output: [1,2,3,5]
func removeNthFromEnd(head *linked_list.Node[int], n int) *linked_list.Node[int] {
	count := 1
	current := head
	for {
		if current.Next == nil {
			break
		}

		current = current.Next
		count++
	}

	if count < n {
		return head
	}

	if count == n {
		head = head.Next
		return head
	}

	i := 1
	j := count - n
	current = head
	next := head
	for {
		if j == i {
			next = current.Next.Next
			current.Next = next
			break
		}

		current = current.Next
		i++
	}

	return head
}

//Given a string s, find the length of the longest substring
// without repeating characters.
//Input: s = "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}

	var cs int
	ss := ""
	for _, b := range s {
		char := string(b)
		if !strings.Contains(ss, char) {
			ss += char
		} else {
			if cs < len(ss) {
				cs = len(ss)
			}
			index := strings.Index(ss, char)
			ss = ss[index+1:] + char
		}
	}

	if cs < len(ss) {
		cs = len(ss)
	}

	return cs
}

//Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
//In other words, return true if one of s1's permutations is the substring of s2.
//Input: s1 = "ab", s2 = "eidbaooo"
//Output: true
//Explanation: s2 contains one permutation of s1 ("ba").
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	//var subs []string
	lenS1 := len(s1)
	lenS2 := len(s2)
	isSub := false
	for i, _ := range s2 {
		end := lenS1 + i
		if end > lenS2 {
			break
		}

		if end <= lenS2 {
			sub := s2[i:end]
			subr := []rune(sub)
			for _, s := range s1 {
				for j, r := range subr {
					if r == s {
						subr = append(subr[0:j], subr[j+1:]...)
						break
					}
				}
			}
			if len(subr) == 0 {
				isSub = true
				break
			}
		}
	}

	return isSub
}

func check(s1, s2 []int) bool {
	for i := 0; i < 26; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func checkInclusionMap(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	// No point in processing if length of s1 is greater than s2.
	if n1 > n2 {
		return false
	}

	// Create a map for s1 once. Need a mapping for each character
	// and its number of occurrences.
	s1Map := make([]int, 26)
	for i := 0; i < n1; i++ {
		s1Map[s1[i]-'a']++
	}

	// Create a similar map for s2. We include all except one character
	// since it will be added in the loop.
	s2Map := make([]int, 26)
	for i := 0; i < n1-1; i++ {
		s2Map[s2[i]-'a']++
	}

	// Sliding window loop.
	for i := n1 - 1; i < n2; i++ {
		// Add a char from s2 to the window
		s2Map[s2[i]-'a']++

		// Check if s1 and s2 maps are equal.
		// If they are equal, we have found a permutation.
		if check(s1Map, s2Map) {
			return true
		}

		// Remove the first character from the window to prepare
		// for the next iteration.
		s2Map[s2[i-n1+1]-'a']--
	}

	// No permutation found.
	return false
}

//An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.
//You are also given three integers sr, sc, and color. You should perform a flood fill on the image starting from the pixel image[sr][sc].
//To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel,
//plus any pixels connected 4-directionally to those pixels (also with the same color), and so on.
//Replace the color of all of the aforementioned pixels with color.
//Return the modified image after performing the flood fill.
//Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2
//Output: [[2,2,2],[2,2,0],[2,0,1]]
//Explanation: From the center of the image with position (sr, sc) = (1, 1) (i.e., the red pixel), all pixels connected by a path of the same color as the starting pixel (i.e., the blue pixels) are colored with the new color.
//Note the bottom corner is not colored 2, because it is not 4-directionally connected to the starting pixel.
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	mainColor := image[sr][sc]
	if mainColor == color {
		return image
	}
	var queue [][]int
	lenCol := len(image[sr])
	lenRow := len(image)
	image[sr][sc] = color
	queue = append(queue, []int{sr, sc})

	for {
		if len(queue) == 0 {
			break
		}
		sr, sc, queue = queue[0][0], queue[0][1], queue[1:]

		if sr-1 >= 0 && image[sr-1][sc] == mainColor {
			image[sr-1][sc] = color
			queue = append(queue, []int{sr - 1, sc})
		}

		if sc+1 < lenCol && image[sr][sc+1] == mainColor {
			image[sr][sc+1] = color
			queue = append(queue, []int{sr, sc + 1})
		}

		if sr+1 < lenRow && image[sr+1][sc] == mainColor {
			image[sr+1][sc] = color
			queue = append(queue, []int{sr + 1, sc})
		}

		if sc-1 >= 0 && image[sr][sc-1] == mainColor {
			image[sr][sc-1] = color
			queue = append(queue, []int{sr, sc - 1})
		}
	}

	return image
}

//You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.)
//You may assume all four edges of the grid are surrounded by water.
//The area of an island is the number of cells with a value 1 in the island.
//Return the maximum area of an island in grid. If there is no island, return 0.
//Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],
//[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],
//[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
//Output: 6
//Explanation: The answer is not 11, because the island must be connected 4-directionally.
func islandCalc(grid [][]int, sr int, sc int, used *map[string]bool) int {
	var count int
	var queue [][]int
	lenRow := len(grid)
	lenCol := len(grid[sr])
	queue = append(queue, []int{sr, sc})

	for {
		if len(queue) == 0 {
			break
		}
		sr, sc, queue = queue[0][0], queue[0][1], queue[1:]

		if grid[sr][sc] == 1 && !(*used)[strconv.Itoa(sr)+strconv.Itoa(sc)] {
			(*used)[strconv.Itoa(sr)+strconv.Itoa(sc)] = true
			count++
		}

		if sr-1 >= 0 && !(*used)[strconv.Itoa(sr-1)+strconv.Itoa(sc)] && grid[sr-1][sc] == 1 {
			queue = append(queue, []int{sr - 1, sc})
		}

		if sc+1 < lenCol && !(*used)[strconv.Itoa(sr)+strconv.Itoa(sc+1)] && grid[sr][sc+1] == 1 {
			queue = append(queue, []int{sr, sc + 1})
		}

		if sr+1 < lenRow && !(*used)[strconv.Itoa(sr+1)+strconv.Itoa(sc)] && grid[sr+1][sc] == 1 {
			queue = append(queue, []int{sr + 1, sc})
		}

		if sc-1 >= 0 && !(*used)[strconv.Itoa(sr)+strconv.Itoa(sc-1)] && grid[sr][sc-1] == 1 {
			queue = append(queue, []int{sr, sc - 1})
		}
	}

	return count
}

func maxAreaOfIsland(grid [][]int) int {
	var queue [][]int
	var islandCount int

	lenRow := len(grid)
	sr := (lenRow - 1) / 2
	lenCol := len(grid[sr])
	sc := (lenCol - 1) / 2
	queue = append(queue, []int{sr, sc})
	used := make(map[string]bool, lenRow*lenCol)

	for i := 0; i < lenRow; i++ {
		for j := 0; j < lenCol; j++ {
			if grid[i][j] == 1 && !used[strconv.Itoa(i)+strconv.Itoa(j)] {
				count := islandCalc(grid, i, j, &used)
				if count > islandCount {
					islandCount = count
				}
			}
		}
	}

	return islandCount
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//You are given two binary trees root1 and root2.
//
//Imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not.
//You need to merge the two trees into a new binary tree.
//The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node.
//Otherwise, the NOT null node will be used as the node of the new tree.
//Return the merged tree.
//Note: The merging process must start from the root nodes of both trees.
//Input: root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
//Output: [3,4,5,5,4,null,7]
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	} else if root1 == nil && root2 != nil {
		return &TreeNode{Val: root2.Val}
	} else if root1 != nil && root2 == nil {
		return root1
	}

	var sum func(root1 *TreeNode, root2 *TreeNode)
	sum = func(root1 *TreeNode, root2 *TreeNode) {
		root1.Val = root1.Val + root2.Val

		if root1.Right != nil && root2.Right != nil {
			sum(root1.Right, root2.Right)
		} else if root1.Right == nil && root2.Right != nil {
			root1.Right = root2.Right
		}

		if root1.Left != nil && root2.Left != nil {
			sum(root1.Left, root2.Left)
		} else if root1.Left == nil && root2.Left != nil {
			root1.Left = root2.Left
		}
	}

	sum(root1, root2)

	return root1
}

// Node You are given a perfect binary tree where all leaves are on the same level,
//and every parent has two children. The binary tree has the following definition:
//Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//Populate each next pointer to point to its next right node.
//If there is no next right node, the next pointer should be set to NULL.
//Initially, all next pointers are set to NULL.
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	root.Next = nil

	var linked func(node *Node)
	linked = func(node *Node) {
		if node.Left != nil {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
			linked(node.Left)
		}

		if node.Right != nil {
			linked(node.Right)
		}
	}

	linked(root)

	return root
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//You are given the heads of two sorted linked lists list1 and list2.
//Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
//Return the head of the merged linked list.
//Input: list1 = [1,2,4], list2 = [1,3,4]
//Output: [1,1,2,3,4,4]
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	current := &ListNode{}

	if list1.Val < list2.Val {
		current = list1
		list1 = list1.Next
	} else {
		current = list2
		list2 = list2.Next
	}

	result := current
	for list1 != nil && list2 != nil {
		if list2.Val <= list1.Val {
			current.Next = list2
			list2 = list2.Next
		} else {
			current.Next = list1
			list1 = list1.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	}

	if list2 != nil {
		current.Next = list2
	}

	return result
}

//Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].
//You may return the answer in any order.
//Input: n = 4, k = 2
//Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
//Explanation: There are 4 choose 2 = 6 total combinations.
//Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.
func combine(n int, k int) [][]int {
	var res [][]int
	var item []int
	for i := 1; i <= n; i++ {
		item = append(item, i)
	}

	if n == k {
		return append(res, item)
	}

	var itemComputation func(item []int, temp []int, start int, end int, index int, k int)
	itemComputation = func(item []int, temp []int, start int, end int, index int, k int) {
		if index == k {
			var t []int
			res = append(res, append(t, temp...))
			return
		}
		i := start
		for i <= end && end-i+1 >= k-index {
			temp[index] = item[i]
			itemComputation(item, temp, i+1, end, index+1, k)
			i++
		}
	}

	temp := make([]int, k)
	itemComputation(item, temp, 0, n-1, 0, k)

	return res
}

//Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
//Input: nums = [1,2,3]
//Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
func permute(nums []int) [][]int {
	var res [][]int

	var itemComputation func(nums []int, n int)
	itemComputation = func(nums []int, n int) {
		if n == 1 {
			var temp []int
			res = append(res, append(temp, nums...))
			return
		}
		for i := 0; i < n; i++ {
			itemComputation(nums, n-1)
			if n%2 == 1 {
				nums[0], nums[n-1] = nums[n-1], nums[0]
			} else {
				nums[i], nums[n-1] = nums[n-1], nums[i]
			}
		}
	}

	length := len(nums)
	itemComputation(nums, length)

	return res
}

//Given a string s, you can transform every letter individually to be lowercase or uppercase to create another string.
//Return a list of all possible strings we could create. Return the output in any order.
//Input: s = "a1b2"
//Output: ["a1b2","a1B2","A1b2","A1B2"]
func letterCasePermutation(s string) []string {
	var resp []string
	a := []rune(s)
	if len(a) == 0 {
		return resp
	}

	var permutationHelper func(s []rune, index int, current []rune)
	permutationHelper = func(s []rune, index int, current []rune) {
		if index == len(s) {
			resp = append(resp, string(current))
			return
		}

		if s[index] >= 65 && s[index] <= 122 {
			// Generate two permutations for each letter: lowercase and uppercase
			if s[index] >= 97 && s[index] <= 122 {
				permutationHelper(s, index+1, append(current, s[index]))
				permutationHelper(s, index+1, append(current, s[index]-32))
			} else {
				permutationHelper(s, index+1, append(current, s[index]+32))
				permutationHelper(s, index+1, append(current, s[index]))
			}
		} else {
			// If the current character is not a letter, just add it to the current permutation
			permutationHelper(s, index+1, append(current, s[index]))
		}
	}

	permutationHelper(a, 0, []rune{})

	return resp
}

//You are climbing a staircase. It takes n steps to reach the top.
//Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
// Fibonacci numbers :
// n - 1 out - 1 (1)
// n - 2 out - 2 (1 + 1, 2)
// n - 3 out - 3 (1 + 1 + 1, 2 + 1, 1 + 2)
// n - 4 out - 5 (1 + 1 + 1 + 1, 1 + 1 + 2, 1 + 2 + 1, 2 + 1 + 1, 2 + 2)
// n - 5 out - 8 (1 + 1 + 1 + 1 + 1, 1 + 1 + 2 + 1, 1 + 2 + 1 + 1, 2 + 1 + 1 + 1, 1 + 1 + 1 + 2, 2 + 2 + 1, 1 + 2 + 2, 2 + 1 + 2)
func climbStairs(n int) int {
	n++
	cache := make(map[int]int, n)
	var climbStairsStep func(n int, cache *map[int]int) int
	climbStairsStep = func(n int, cache *map[int]int) int {
		if n < 2 && n > 0 {
			(*cache)[n] = n
			return n
		}

		if len(*cache) > n {
			return (*cache)[n]
		} else {
			(*cache)[n] = climbStairsStep(n-2, cache) + climbStairsStep(n-1, cache)
			return (*cache)[n]
		}
	}

	climbStairsStep(n, &cache)

	return cache[n]
}

//Given a triangle array, return the minimum path sum from top to bottom.
//For each step, you may move to an adjacent number of the row below.
//More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.
//Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
//Output: 11
//Explanation: The triangle looks like:
//   -1
//  2 3
// 1 -1 -3
//4 1 8 3
//The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).
func minimumTotal(triangle [][]int) int {
	sum := 0
	length := len(triangle)
	if length == 0 {
		return sum
	}
	cache := make(map[string]int, length)

	var getKey func(i int, j int) string
	getKey = func(i int, j int) string {
		return strconv.Itoa(i-1) + strconv.Itoa(j)
	}

	var rowMin func(row []int, i int)
	rowMin = func(row []int, i int) {
		for j := 0; j < len(row); j++ {
			prev := 0
			if i-1 >= 0 {
				if j == 0 {
					prev = cache[getKey(i-1, j)]
				} else if j == len(row)-1 {
					prev = cache[getKey(i-1, j-1)]
				} else {
					prev = cache[getKey(i-1, j-1)]
					if cache[getKey(i-1, j)] < cache[getKey(i-1, j-1)] {
						prev = cache[getKey(i-1, j)]
					}
				}
			}

			cache[getKey(i, j)] = row[j] + prev
		}
	}

	for i, v := range triangle {
		rowMin(v, i)
	}

	min := cache[getKey(length-1, 0)]
	for i := 1; i < len(triangle[length-1]); i++ {
		current, ok := cache[getKey(length-1, i)]
		if ok && min > current {
			min = cache[getKey(length-1, i)]
		}
	}

	return min
}

//Given an integer n, return true if it is a power of two. Otherwise, return false.
//An integer n is a power of two, if there exists an integer x such that n == 2x.
//Input: n = 16
//Output: true
//Explanation: 24 = 16
func isPowerOfTwo(n int) bool {
	res := true
	for {
		if n <= 2 {
			break
		}

		if n <= 0 || n%2 != 0 {
			res = false
			break
		}

		n = n / 2
	}

	return res
}

func isPowerOfTwoBits(n int) bool {
	if n%2 == 1 && n != 1 || n <= 0 {
		return false
	}

	for offset := 0; offset < 31; offset++ {
		if n == 1<<offset {
			return true
		}

	}

	return false

}

//Write a function that takes the binary representation of an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).
//Input: n = 00000000000000000000000000001011
//Output: 3
func hammingWeight(num uint64) int {
	count := 0
	for num != 0 {
		count++
		num &= num - 1

	}

	return count
}

//Reverse bits of a given 32 bits unsigned integer.
//Input: n = 00000010100101000001111010011100
//Output:    964176192 (00111001011110000010100101000000)
//Explanation: The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596,
//so return 964176192 which its binary representation is 00111001011110000010100101000000.
func reverseBits(num uint32) uint32 {
	var rev uint32
	i := 0
	for num > 0 {
		rev = rev << 1
		if num&1 == 1 {
			rev = rev ^ 1
		}
		num = num >> 1
		i++
	}
	if i < 32 {
		rev = rev << (32 - i)
	}

	return rev
}

//Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
//You must implement a solution with a linear runtime complexity and use only constant extra space.

//Input: nums = [2,2,1]
//Output: 1
func singleNumber(nums []int) int {
	var item int
	m := make(map[int]bool, len(nums))
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			m[v] = true
		} else {
			m[v] = false
		}
	}

	for k, v := range m {
		if !v {
			item = k
			break
		}
	}

	return item
}

//Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
//You must implement a solution with a linear runtime complexity and use only constant extra space.

//Input: nums = [2,2,1]
//Output: 1
func singleNumberXor(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}

func rob(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	var amount int
	cache := make(map[int]int, length)

	for i := 0; i < length; i++ {
		if i-2 < 0 {
			cache[i] = nums[i]
			if amount < cache[i] {
				amount = cache[i]
			}
			continue
		}
		localAmount := 0
		for j := i - 2; j >= 0; j-- {
			if localAmount < cache[j] {
				localAmount = cache[j]
			}
		}
		cache[i] = nums[i] + localAmount
		if amount < cache[i] {
			amount = cache[i]
		}
	}

	return amount
}

//Given the head of a singly linked list, reverse the list, and return the reversed list.
//Input: head = [1,2,3,4,5]
//Output: [5,4,3,2,1]
func reverseList(head *ListNode) *ListNode {
	current := head
	var prev *ListNode
	var next *ListNode
	for {
		if current == nil {
			break
		}

		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

//Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
//The distance between two adjacent cells is 1.
//Input: mat = [[0,0,0],[0,1,0],[0,0,0]]
//Output: [[0,0,0],[0,1,0],[0,0,0]]
func updateMatrix(mat [][]int) [][]int {
	i, j := 0, 0
	lenCel := len(mat[0])
	lenRow := len(mat)
	list := make([]int, lenCel*lenRow)
	//current := mat[i][j]
	var step func(i int, j int)
	step = func(i int, j int) {
		if i+1 > lenRow || j+1 > lenCel {
			return
		}
		list = append(list, mat[i][j])
		step(i, j+1)
		step(i+1, j)
	}

	step(i, j)

	return mat
}

//You are given an m x n grid where each cell can have one of three values:
//0 representing an empty cell,
//1 representing a fresh orange, or
//2 representing a rotten orange.
//Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.
//Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.
//Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
//Output: 4
func orangesRotting(grid [][]int) int {
	lenRow := len(grid)
	if lenRow == 0 {
		return 0
	}
	lenCol := len(grid[0])
	sr, sc := 0, 0
outFirst:
	for i := 0; i < lenRow; i++ {
		for j := 0; j < lenCol; j++ {
			if grid[i][j] == 2 {
				sr, sc = i, j
				break outFirst
			}
		}
	}
	var queue [][]int
	queue = append(queue, []int{sr, sc})
	timeCount := 0

	for {
		localCount := 0
		if len(queue) == 0 {
			break
		}
		sr, sc, queue = queue[0][0], queue[0][1], queue[1:]
		if sr == 7 && sc == 7 {
			fmt.Println(1)
		}

		if grid[sr][sc] == 0 {
			continue
		}

		if (sr+1 < lenRow && sc+1 < lenCol) &&
			((grid[sr+1][sc] == 2 && grid[sr][sc+1] == 2 && grid[sr-1][sc] == 2 && grid[sr][sc-1] == 2) ||
				(grid[sr+1][sc] == 0 && grid[sr][sc+1] == 0 && grid[sr-1][sc] == 0 && grid[sr][sc-1] == 0)) {
			continue
		}

		if sc+1 < lenCol && grid[sr][sc+1] == 1 {
			grid[sr][sc+1] = 2
			queue = append(queue, []int{sr, sc + 1})
			localCount++
		}

		if sr+1 < lenRow && grid[sr+1][sc] == 1 {
			grid[sr+1][sc] = 2
			queue = append(queue, []int{sr + 1, sc})
			localCount++
		}

		if sr-1 >= 0 && grid[sr-1][sc] == 1 {
			grid[sr-1][sc] = 2
			queue = append(queue, []int{sr - 1, sc})
			localCount++
		}

		if sc-1 >= 0 && grid[sr][sc-1] == 1 {
			grid[sr][sc-1] = 2
			queue = append(queue, []int{sr, sc - 1})
			localCount++
		}
		if localCount > 0 {
			timeCount++
			grid[sr][sc] = 2
		}
	}
out:
	for i := 0; i < lenRow; i++ {
		for j := 0; j < lenCol; j++ {
			if grid[i][j] == 1 {
				timeCount = -1
				break out
			}
		}
	}

	return timeCount
}

//You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n,
//representing the number of elements in nums1 and nums2 respectively.
//Merge nums1 and nums2 into a single array sorted in non-decreasing order.
//The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
//To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged,
//and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
//Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//Output: [1,2,2,3,5,6]
//Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
//The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1Fil := make([]int, m)
	copy(nums1Fil, nums1[:m])
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1Fil[i] < nums2[j] {
			nums1[k] = nums1Fil[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}
	for ; i < m; i++ {
		nums1[k] = nums1Fil[i]
		k++
	}
	for ; j < n; j++ {
		nums1[k] = nums2[j]
		k++
	}
}

//Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
//The order of the elements may be changed.
//Then return the number of elements in nums which are not equal to val.
//Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:
//Change the array nums such that the first k elements of nums contain the elements which are not equal to val.
//The remaining elements of nums are not important as well as the size of nums.
//Return k.
func removeElement(nums []int, val int) int {
	var i int
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}

	return i
}

//Input: nums = [1,1,2]
//Output: 2, nums = [1,2,_]
//Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
//It does not matter what you leave beyond the returned k (hence they are underscores).
func removeDuplicates(nums []int) int {
	var index int
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if prev != nums[i] {
			index++
			prev = nums[i]
			nums[index] = prev
		} else {
			nums[index] = nums[i]
		}
	}

	if nums[index] <= prev {
		nums[index] = prev
		index++
	}

	return index
}

//Given an integer array nums sorted in non-decreasing order,
//remove some duplicates in-place such that each unique element appears at most twice.
//The relative order of the elements should be kept the same.
//Input: nums = [1,1,1,2,2,3]
//Output: 5, nums = [1,1,2,2,3,_]
//Explanation: Your function should return k = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.
//It does not matter what you leave beyond the returned k (hence they are underscores).
//
//Approach
//I have two pointers:
//First pointer is the result which contains at least two duplicate elements
//Second pointer is the iteration from 1 -> n
//I move the first pointer when:
// 1. The element from the second pointer is different to the first pointer element.
// 2. Or the element from the second pointer equals to the first pointer element,
//but the counter is still avalible (less then at least k duplicate elements)
func removeDuplicatesTwo(nums []int) int {
	k := 0
	length := len(nums)
	count := 1

	for i := 1; i < length; i++ {
		if nums[i] != nums[k] {
			k++
			count = 1
			nums[k] = nums[i]
		} else if count < 2 {
			count++
			k++
			nums[k] = nums[i]
		}

	}
	return k + 1
}

//Given an array nums of size n, return the majority element.
//The majority element is the element that appears more than ⌊n / 2⌋ times.
//You may assume that the majority element always exists in the array.
//Input: nums = [3, 3, 3, 3, 1, 5, 3, 5, 3, 1, 3, 5, 2, 2, 3]
//Output: 2
func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1
	for _, num := range nums[1:] {
		if num == candidate {
			count++
		} else {
			count--
		}
		if count == 0 {
			candidate = num
			count = 1
		}
	}
	return candidate
}

//Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
//You must implement a solution with a linear runtime complexity and use only constant extra space.
//Input: nums = [4,1,2,1,2]
//Output: 4
func addBinary(a string, b string) string {
	lenA := len(a)
	lenB := len(b)
	lenC := len(a)
	if lenA < lenB {
		lenC = lenB
	}
	i, j := lenA-1, lenB-1
	k := i
	if j > i {
		k = j
	}
	c := make([]rune, lenC)
	var step uint8
	for i >= 0 && j >= 0 {
		sum := a[i] + b[j] + step
		if sum == 98 || sum == 99 {
			c[k] = rune(sum - 50)
			step = 1
		}
		if sum <= 97 {
			c[k] = rune(sum - 48)
			step = 0
		}
		i--
		j--
		k--
	}

	for i >= 0 {
		sum := a[i] + step
		step = 0
		c[i] = rune(sum)
		if sum == 50 {
			c[i] = rune(sum - 2)
			step = 1
		}

		i--
	}

	for j >= 0 {
		sum := b[j] + step
		step = 0
		c[j] = rune(sum)
		if sum == 50 {
			c[j] = rune(sum - 2)
			step = 1
		}
		j--
	}

	if step == 1 {
		c = append([]rune("1"), c[0:]...)
	}

	return string(c)
}

//Given an integer array nums where every element appears three times except for one, which appears exactly once.
//Find the single element and return it.
//You must implement a solution with a linear runtime complexity and use only constant extra space.
//use implementation from https://lenchen.medium.com/leetcode-137-single-number-ii-31af98b0f462
//Input: nums = [2,2,3,2]
//Output: 3
func singleNumberTwo(nums []int) int {
	ones, twos := 0, 0

	for _, n := range nums {
		ones = (ones ^ n) & ^twos
		twos = (twos ^ n) & ^ones
		fmt.Println(1)
	}

	return ones

}

//Given two integers left and right that represent the range [left, right],
//return the bitwise AND of all numbers in this range, inclusive.
//Example 1:
//Input: left = 5, right = 7
//Output: 4
//
//Example 2:
//Input: left = 0, right = 0
//Output: 0
func rangeBitwiseAnd(left int, right int) int {
	if left == right {
		return left
	}
	if left == 1 && right == 2 {
		return 0
	}
	find := left
	count := 0
	for find != 0 {
		find >>= 1
		count++
	}
	start := 1 << count

	if start <= right {
		return 0
	}

	count = left
	min := start >> 1
	for right > left {
		if count == min {
			break
		}
		count &= right
		right--
	}

	return count
}

func largestElement(slice []int, index int) (int, int) {
	if len(slice) == 1 {
		return index, slice[0]
	}
	var max int
	i := index
	max = slice[0]
	maxJump := max + index

	for j := 1; j < len(slice); j++ {
		current := slice[j]
		if current+index+j >= maxJump {
			max = current
			maxJump = max + index + j
			i = index + j
		}
	}
	return i, max
}

func canJump(nums []int) bool {
	length := len(nums)
	if length == 1 {
		return true
	}
	res := false
	i := 0
	countSteps := nums[i]
	for {
		if countSteps >= length-i-1 {
			res = true
			break
		}
		if countSteps == 0 {
			break
		}
		i, countSteps = largestElement(nums[i+1:countSteps+i+1], i+1)
	}

	return res
}

func jump(nums []int) int {
	length := len(nums)
	count := 0
	if length == 1 {
		return count
	}
	i := 0
	countSteps := nums[i]
	for {
		count++
		if countSteps >= length-i-1 {
			break
		}
		if nums[countSteps] >= length-countSteps-1 {
			count++
			break
		}
		if countSteps == 0 {
			break
		}
		i, countSteps = largestElement(nums[i+1:countSteps+i+1], i+1)
	}

	return count
}

func romanToInt(s string) int {
	r := []rune(s)
	length := len(r)
	mapRoman := map[rune]int{
		73: 1,
		86: 5,
		88: 10,
		76: 50,
		67: 100,
		68: 500,
		77: 1000,
	}
	res := 0

	for i := 0; i < length; i++ {
		num := mapRoman[r[i]]
		if i+1 < length && (mapRoman[r[i]] < mapRoman[r[i+1]]) {
			num = mapRoman[r[i+1]] - mapRoman[r[i]]
			i++
		}

		res += num
	}

	return res
}

func intToRoman(num int) string {
	max := 10000
	i := max
	var numList []rune
	mapRoman := map[int]rune{
		1:    73,
		5:    86,
		10:   88,
		50:   76,
		100:  67,
		500:  68,
		1000: 77,
	}

	for {
		if i == 1 {
			break
		}
		step := i / 10
		if num/step != 0 {
			count := num % i / step
			if count <= 3 {
				for j := 1; j <= count; j++ {
					numList = append(numList, mapRoman[step])
				}
			}

			if count == 4 && step < max {
				numList = append(numList, mapRoman[step])
				numList = append(numList, mapRoman[5*step])
			}

			if count == 5 && step < max {
				numList = append(numList, mapRoman[5*step])
			}

			if count > 5 && count < 9 && step < max {
				numList = append(numList, mapRoman[5*step])
				for j := 1; j <= count-5; j++ {
					numList = append(numList, mapRoman[step])
				}
			}

			if count == 9 && step < max {
				numList = append(numList, mapRoman[step])
				numList = append(numList, mapRoman[10*step])
			}
		}

		i /= 10
	}

	return string(numList)
}

// ArrayDiff implement a difference function, which subtracts one list from another and returns the result.
//It should remove all values from list a, which are present in list b keeping their order.
// ArrayDiff([]int{1,2,3}, []int{2}) = []int{1,3}
// ArrayDiff([]int{1,2,2},[]int{2}) = []int{1}
//If a value is present in b, all of its occurrences must be removed from the other:
func ArrayDiff(a, b []int) []int {
	lenA := len(a)
	lenB := len(b)
	if lenB == 0 {
		return a
	}
	if lenA == 0 {
		return []int{}
	}

	var check func(i int)
	check = func(i int) {
		if i+1 < len(a) && a[i] == a[i+1] {
			a = append(a[:i], a[i+1:]...)
			check(i)
		}
	}

	i, j := 0, 0
	for i < lenA && j < lenB {
		if a[i] == b[j] {
			check(i)
			a = append(a[:i], a[i+1:]...)
			j++
		} else {
			i++
		}
	}

	return a
}
