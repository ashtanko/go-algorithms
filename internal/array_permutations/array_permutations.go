package array_permutations

func permute(nums []int) [][]int {
	backtracks := map[int]bool{}
	var f func(int, int, []int, [][]int) [][]int
	f = func(i, walked int, row []int, out [][]int) [][]int {

		if backtracks[i] {
			return out
		}

		backtracks[i] = true

		defer func() {
			backtracks[i] = false
		}()

		row = append(row, nums[i])
		walked++
		if walked == len(nums) {
			out = append(out, append([]int(nil), row...))
			return out
		}
		for j := 0; j < len(nums); j++ {
			out = f(j, walked, row, out)
		}
		return out
	}

	var out [][]int
	for i := 0; i < len(nums); i++ {
		out = f(i, 0, nil, out)
	}
	return out
}
