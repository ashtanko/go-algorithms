package check_array

func check(nums []int) bool {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[(i+1)%len(nums)] {
			k++
		}
		if k > 1 {
			return false
		}
	}
	return true
}
