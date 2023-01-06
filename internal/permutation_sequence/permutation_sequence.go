package permutation_sequence

import "strconv"

func getPermutation(n int, k int) string {
	res, str := "", ""
	visited := make(map[int]bool)
	backTrack(n, &k, visited, str, &res)
	return res
}

func backTrack(n int, k *int, visited map[int]bool, str string, res *string) {
	if len(str) == n {
		*k--
		if *k == 0 { // k == 0 means find the answer
			*res = str
			return
		}
	}

	for i := 1; i <= n; i++ {
		if visited[i] {
			continue
		}

		if f(n-len(str)-1) < *k {
			*k -= f(n - len(str) - 1) // if the number of permutations is less than k, we can substract the number from k directly
			continue
		}
		str = str + strconv.Itoa(i)
		visited[i] = true
		backTrack(n, k, visited, str, res)
		if *res != "" {
			return
		}
		visited[i] = false
		str = str[:len(str)-1]

	}

}

func f(n int) int {
	res := 1
	for i := n; i > 1; i-- {
		res *= i
	}
	return res
}
