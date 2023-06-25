// Package set s1 := &set.Set[string]{}
// s1.New([]string{"a", "b", "c", "d"})
// s2 := &set.Set[string]{}
// s2.New([]string{"a", "d", "e", "f"})
// s1.Subtraction(s2) // &{[b c]}
// s1.Intersection(s2) // &{[a d]}
// s1.Union(s2) // &{[a b c d a d e f]}
// s1.Complement(s2) // &{[e f]}
package set

import (
	"fmt"
	"golang.org/x/exp/slices"
)

type Set[T comparable] struct {
	data []T
}

// New create Set from list
func (set *Set[T]) New(list []T) *Set[T] {
	itemMap := make(map[T]bool, len(list))
	for _, item := range list {
		if !itemMap[item] {
			set.data = append(set.data, item)
			itemMap[item] = true
		}
	}

	return set
}

func (set *Set[T]) GetData() []T {
	return set.data
}

// Push add new item to Set in the end and return changed Set or Error if element exist in Set
func (set *Set[T]) Push(item T) *Set[T] {
	for _, element := range set.data {
		if element == item {
			return set
		}
	}

	slice := append(set.data, item)
	set.data = slice

	return set
}

// Delete remove item from Set and return changed Set and error if index doesn't exist in Set
func (set *Set[T]) Delete(value T) (*Set[T], error) {
	deleted := false
	for i, item := range set.data {
		if item == value {
			slice := append(set.data[:i], set.data[i+1:]...)
			set.data = slice
			deleted = true
		}
	}

	if !deleted {
		return set, fmt.Errorf("%s dosn't exist in Set", value)
	}

	return set, nil
}

// Size get size of Set
func (set *Set[T]) Size() int {
	return len(set.data)
}

// In check value exist in Set
func (set *Set[T]) In(value T) bool {
	checked := false
	if set.Size() == 0 {
		return false
	}

	for _, it := range set.data {
		if it == value {
			checked = true
			break
		}
	}

	return checked
}

// Intersection returns a new set with the common elements in the set s and the passed set
func (set *Set[T]) Intersection(other *Set[T]) *Set[T] {
	var iSet Set[T]
	if set.Size() == 0 || other.Size() == 0 {
		return &iSet
	}
	for _, item := range set.data {
		if other.In(item) {
			iSet.Push(item)
		}
	}

	return &iSet
}

// Subtraction returns a new set with the elements in the current Set but not in the passed set
func (set *Set[T]) Subtraction(other *Set[T]) *Set[T] {
	var sSet Set[T]
	if set.Size() == 0 {
		return &sSet
	}
	for _, item := range set.data {
		if !other.In(item) {
			sSet.Push(item)
		}
	}

	return &sSet
}

// Union returns a new set with the all elements in the set s and the passed set
func (set *Set[T]) Union(other *Set[T]) *Set[T] {
	if other.Size() == 0 {
		return set
	}

	items := make([]T, set.Size())
	copy(items, set.data)
	for _, item := range other.data {
		items = append(items, item)
	}

	return &Set[T]{data: items}
}

// Complement returns a set of all elements that are in the universal set but are not in base Set
func (set *Set[T]) Complement(universal *Set[T]) *Set[T] {
	if universal.Size() == 0 {
		return nil
	}

	if set.Size() == 0 {
		return universal
	}

	items := make([]T, universal.Size())
	copy(items, universal.data)
	for _, item := range set.data {
		if !universal.In(item) {
			continue
		}

		if i := slices.Index(items, item); i != -1 {
			items = append(items[:i], items[i+1:]...)
		}
	}

	return &Set[T]{data: items}
}
