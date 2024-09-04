package permutation_sequence

import "strconv"

func getPermutation(n int, k int) string {
	res, str := "", ""
	visited := make(map[int]bool)
	backtrack(n, &k, visited, str, &res)
	return res
}

func backtrack(n int, k *int, visited map[int]bool, current string, result *string) {
	if len(current) == n {
		*k--
		if *k == 0 { // k == 0 means find the answer
			*result = current
			return
		}
	}

	for i := 1; i <= n; i++ {
		if visited[i] {
			continue
		}

		if factorial(n-len(current)-1) < *k {
			*k -= factorial(n - len(current) - 1)
			continue
		}
		current = current + strconv.Itoa(i)
		visited[i] = true
		backtrack(n, k, visited, current, result)
		if *result != "" {
			return
		}
		visited[i] = false
		current = current[:len(current)-1]
	}
}

func factorial(n int) int {
	result := 1
	for i := n; i > 1; i-- {
		result *= i
	}
	return result
}
