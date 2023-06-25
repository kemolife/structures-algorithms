package recursive

import "errors"

// FactorialRecursive 0! = 1
// 1! = 1
// n! = n * (n-1)!
func FactorialRecursive(n uint64) uint64 {
	if n < 2 {
		return 1
	}

	return n * FactorialRecursive(n-1)
}

// FactorialIteration 0! = 1
// 1! = 1
// n! = n * (n-1)!
func FactorialIteration(n uint64) uint64 {
	var result uint64
	result = 1
	for {
		if n < 2 {
			result *= 1
			break
		}

		result *= n
		n--
	}

	return result
}

func fibonacciStep(n uint64) uint64 {
	if n < 2 {
		return n
	}

	return fibonacciStep(n-2) + fibonacciStep(n-1)
}

// FibonacciRecursive O(2^n)
// n < 2: Fn = 1.
// Fn = Fn-1 + Fn-2
func FibonacciRecursive(n uint64) []uint64 {
	var fibonacciList []uint64
	var i uint64
	for i = 0; i <= n; i++ {
		fibonacciList = append(fibonacciList, fibonacciStep(i))
	}

	return fibonacciList
}

// FibonacciInteractive O(n)
// n < 2: Fn = 1.
// Fn = Fn-1 + Fn-2
func FibonacciInteractive(n uint64) []uint64 {
	var i uint64
	l := make([]uint64, n+1)
	l[0] = 0
	l[1] = 1
	for i = 2; i <= n; i++ {
		l[i] = l[i-2] + l[i-1]
	}

	return l
}

// [1, 2, 4, 4, 2, 6, 7, 1, 7]
func recurringCharacters(s []int) (int, error) {
	if len(s) < 2 {
		return -1, errors.New("not valid")
	}

	minIndex := 0
	value := 0
	recurringMap := make(map[int][]int)
	for i := 0; i < len(s); i++ {
		if len(recurringMap[s[i]]) == 0 {
			recurringMap[s[i]] = []int{i}
		} else {
			if recurringMap[s[i]][0] == 0 {
				value = s[i]
				break
			}
			if minIndex != 0 && minIndex > recurringMap[s[i]][0] {
				minIndex = recurringMap[s[i]][0]
				value = s[i]
			}
			recurringMap[s[i]] = append(recurringMap[s[i]], i)
		}
	}

	if value == 0 {
		return -1, errors.New("not found")
	}

	return value, nil
}
