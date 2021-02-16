package running_sum

func runningSum(nums []int) []int {
	i := 1
	for i < len(nums) {
		nums[i] += nums[i-1]
		i++
	}
	return nums
}
