package sum_range

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{[]int{}}
	}
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	return NumArray{sums: sums}
}

func (a *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return a.sums[j]
	}
	return a.sums[j] - a.sums[i-1]
}
