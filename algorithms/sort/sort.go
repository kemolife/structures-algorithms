package main

import "fmt"

type compatible interface {
	~int | ~string
}

// always check all items and did one additional iteration in the end.
func bubbleSortNotOptimal(input []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(input); i++ {
			if input[i-1] > input[i] {
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
	return input
}

// BubbleSort [3,4,6,10,5,1,9,2]
func BubbleSort[T compatible](list []T) []T {
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-i-1; j++ {
			if list[j] > list[j+1] {
				list[j+1], list[j] = list[j], list[j+1]
			}
		}
	}

	return list
}

// SelectionSort check min element and go it to top, iterate for all element in list O(n^2)
func SelectionSort[T compatible](list []T) []T {
	for i := 0; i < len(list)-1; i++ {
		min := list[i]
		for j := i + 1; j < len(list); j++ {
			if min > list[j] {
				min, list[j] = list[j], min
			}
		}
		list[i] = min
	}

	return list
}

// InsertionSort check min element for iteration and insert in right index
func InsertionSort[T compatible](list []T) []T {
	if len(list) < 2 {
		return list
	}
	newList := []T{list[0], list[1]}
	if list[0] > list[1] {
		newList[0], newList[1] = newList[1], newList[0]
	}
	for i := 2; i < len(list); i++ {
		isModifiedList := false
		for j := 0; j < len(newList); j++ {
			if list[i] < newList[j] {
				isModifiedList = true
				if j == 0 {
					newList = append([]T{list[i]}, newList...)
				} else {
					newList = append(newList[:j], append([]T{list[i]}, newList[j:]...)...)
				}
				break
			}
		}
		if !isModifiedList {
			newList = append(newList, list[i])
		}
	}

	return newList
}

func merge[T compatible](left []T, right []T) []T {
	var combine []T
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			combine = append(combine, left[i])
			i++
		} else {
			combine = append(combine, right[j])
			j++
		}
	}
	for ; i < len(left); i++ {
		combine = append(combine, left[i])
	}
	for ; j < len(right); j++ {
		combine = append(combine, right[j])
	}
	return combine
}

func MergeSort[T compatible](list []T) []T {
	if len(list) < 2 {
		return list
	}

	middleIndex := len(list) / 2
	left := list[:middleIndex]
	right := list[middleIndex:]

	list = merge(MergeSort(left), MergeSort(right))

	return list
}

func combine[T compatible](left []T, right []T, pivot T) []T {
	right = append([]T{pivot}, right...)
	if left == nil {
		return right
	} else if right == nil {
		return append(left, pivot)
	} else {
		return append(left, right...)
	}
}

func QuickSort[T compatible](list []T) []T {
	lengthList := len(list)
	if len(list) < 2 {
		return list
	}

	pivot := list[lengthList-1]
	var left []T
	var right []T
	for _, v := range list[:lengthList-1] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return combine(QuickSort(left), QuickSort(right), pivot)
}

func main() {
	l := []int{3, 4, 6, 10, 1, 2, 3, 5, 9, 14, 15}
	fmt.Println(QuickSort(l))
	//fmt.Println(merge(l1, l2))
	//fmt.Println(bubbleSortNotOptimal(l))
}
