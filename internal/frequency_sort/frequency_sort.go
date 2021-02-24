package frequency_sort

import "sort"

func frequencySort(nums []int) []int {
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if frequency[nums[i]] == frequency[nums[j]] {
			return nums[i] > nums[j]
		}
		return frequency[nums[i]] < frequency[nums[j]]
	})
	return nums
}
