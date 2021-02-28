package min_rotated_sorted_array

func findMinRotatedSorted(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		if nums[start] < nums[end] {
			return nums[start]
		}
		mid := (start + end) / 2
		if nums[mid] >= nums[start] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return nums[start]
}
