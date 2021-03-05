package find_duplicate

import (
	"testing"
)

type TestCase struct {
	nums     FindDuplicateData
	expected int
}

func TestFindDuplicate(t *testing.T) {
	//data := []struct {
	//	tc TestCase
	//}{
	//	{
	//		nums:     FindDuplicateData{[]int{1, 3, 4, 2, 2}},
	//		expected: 2,
	//	},
	//	{
	//		nums:     FindDuplicateData{[]int{3, 1, 3, 4, 2}},
	//		expected: 3,
	//	},
	//	{
	//		nums:     []int{1, 1},
	//		expected: 1,
	//	},
	//	{
	//		nums:     []int{1, 1, 2},
	//		expected: 1,
	//	},
	//	{
	//		nums:     []int{},
	//		expected: 0,
	//	},
	//}
	//
	//testCases := []struct {
	//	performer FindDuplicate
	//	data TestCase
	//}{
	//	{},
	//}
	//
	//for _, tc := range testCases {
	//	name := fmt.Sprintf("%d", tc.nums)
	//	t.Run(name, func(t *testing.T) {
	//		actual := tc.performer.Sort()
	//		utils.Checkf(t, is.Equal(tc.expected, actual), tc)
	//	})
	//}
}
