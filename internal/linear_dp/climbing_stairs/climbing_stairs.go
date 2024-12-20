package climbing_stairs

func climbingStairsBottomUpDp(n int) int {
	if n == 1 {
		return 1
	}
	prev, curr := 1, 1
	for i := 2; i <= n; i++ {
		curr = prev + curr
		prev = curr - prev
	}
	return curr
}
