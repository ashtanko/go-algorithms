package find_duplicate

import (
	"sort"
)

type FindDuplicate interface {
	Sort() int
	Set() int
}

type FindDuplicateData struct {
	nums []int
}

func (d FindDuplicateData) Sort() int {
	nums := d.nums
	n := len(nums)
	sort.Ints(nums)
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return 0
}
