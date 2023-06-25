package array

import (
	"errors"
	"fmt"
)

type Array[T comparable] struct {
	Data []T
}

func (a *Array[T]) Get(index int) (T, error) {
	if (len(a.Data) - 1) < index {
		var empty T
		return empty, errors.New("index out of range")
	}

	return a.Data[index], nil
}

func (a *Array[T]) Add(value ...T) {
	a.Data = append(a.Data, value...)
}

func (a *Array[T]) Update(i int, v T) error {
	_, err := a.Get(i)
	if err != nil {
		return err
	}

	a.Data[i] = v

	return nil
}

func (a *Array[T]) Pop(i int) (v T, e error) {
	value, err := a.Get(i)
	if err != nil {
		var empty T
		return empty, err
	}

	a.Data = append(a.Data[:i], a.Data[i+1:]...)

	return value, nil
}

func (a *Array[T]) Remove(value T) error {
	beforeLen := len(a.Data)
	for i, v := range a.Data {
		if v == value {
			a.Pop(i)
			break
		}
	}

	if len(a.Data) == beforeLen {
		return fmt.Errorf("%v not in the list", value)
	}

	return nil
}

func (a *Array[T]) Search(value T) (int, error) {
	var i int
	for i = 0; i < len(a.Data)-1; i++ {
		if a.Data[i] == value {
			break
		}
	}

	return i, nil
}

func (a *Array[T]) Clear() {
	a.Data = nil
}
