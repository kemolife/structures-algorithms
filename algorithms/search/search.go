package search

import (
	"fmt"
)

type compatible interface {
	~int | ~string
}

func LinerSearch[T comparable](list []T, value T) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}

	return -1
}

func Include[T comparable](list []T, value T) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}

	return false
}

func BinarySearch[T compatible](list []T, value T) bool {
	lengthList := len(list)
	if value > list[lengthList-1] {
		return false
	}
	if value < list[0] {
		return false
	}

	search := false
	var middle int
	for {
		middle = len(list) / 2
		if len(list) == 1 && list[middle] != value {
			break
		}
		if list[middle] == value {
			search = true
			break
		}

		if value < list[middle] {
			list = list[:middle]
		} else {
			list = list[middle+1:]
		}
	}

	return search
}

func main() {
	list := []int{1, 3, 4, 6, 7, 8, 10, 13, 14, 16, 18, 19, 34, 67, 78, 101, 145, 170, 201, 210, 220}
	fmt.Println(BinarySearch(list, -1))
}
