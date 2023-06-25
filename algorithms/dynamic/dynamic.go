package dynamic

type Cache struct {
	data  map[int]int
	Count int
}

// 0, 1, 1, 2, 3, 5, 8, 13, 22, 35...
func Fibonacci(n int, cache *Cache) int {
	cache.Count++
	if n < 2 {
		if len(cache.data) == 0 {
			cache.data = make(map[int]int)
		}
		cache.data[n] = n
		return n
	}

	if len(cache.data) > n {
		return cache.data[n]
	} else {
		cache.data[n] = Fibonacci(n-2, cache) + Fibonacci(n-1, cache)
		return cache.data[n]
	}
}
