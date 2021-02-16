package num_identical_pairs

func numIdenticalPairs(nums []int) int {
	cnt := make(map[int]int)
	var pairs int
	for _, num := range nums {
		pairs += cnt[num]
		cnt[num]++
	}
	return pairs
}
