package fibonacci_number

import "math"

// Approach 1: Recursion
func recursive(n int) int {
	if n <= 1 {
		return n
	}
	return recursive(n-1) + recursive(n-2)
}

// Approach 2: Bottom-Up Approach using Memoization
func bottomUp(n int) int {
	if n <= 1 {
		return n
	}
	return bottomUpMemoize(n)
}

func bottomUpMemoize(n int) int {
	cache := map[int]int{0: 0, 1: 1}
	for i := 2; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[n]
}

// Approach 3: Top-Down Approach using Memoization
var cache = map[int]int{0: 0, 1: 1}

func topDown(n int) int {
	if n <= 1 {
		return n
	}
	return topDownMemoize(n)
}

func topDownMemoize(N int) int {
	if _, ok := cache[N]; ok {
		return cache[N]
	}
	cache[N] = topDownMemoize(N-1) + topDownMemoize(N-2)
	return topDownMemoize(N)
}

// Approach 6: Math
func fibMath(n int) int {
	var goldenRatio = (1 + math.Sqrt(5)) / 2
	return int(math.Round(math.Pow(goldenRatio, float64(n)) / math.Sqrt(5)))
}
