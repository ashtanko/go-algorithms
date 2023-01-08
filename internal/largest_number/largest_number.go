package largest_number

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(strs, func(i, j int) bool {
		return (strs[i] + strs[j]) > (strs[j] + strs[i])
	})
	numsStr := strings.Join(strs, "")
	numsStr = strings.TrimLeft(numsStr, "0")
	if numsStr == "" {
		return "0"
	}
	return numsStr
}

type ByLength []string

func (a ByLength) Len() int           { return len(a) }
func (a ByLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLength) Less(i, j int) bool { return (a[i] + a[j]) > (a[j] + a[i]) }

func largestNumber2(nums []int) string {
	num := make([]string, 0)
	for i := 0; i < len(nums); i++ {
		num = append(num, strconv.Itoa(nums[i]))
	}
	sort.Sort(ByLength(num))

	numStr := strings.Join(num, "")
	numStr = strings.TrimLeft(numStr, "0")
	if numStr == "" {
		return "0"
	}
	return numStr
}
