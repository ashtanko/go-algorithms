package shuffle_array

func shuffle(nums []int, n int) []int {
	for i := 0; i < len(nums)/2; i++ {
		nums[i+n] = nums[i+n] << 10
		nums[i+n] = nums[i+n] | nums[i]
	}

	for i := 0; i < len(nums); i += 2 {
		nums[i] = nums[n+i/2] & 1023
		nums[i+1] = nums[n+(i+1)/2] >> 10
	}
	return nums
}
