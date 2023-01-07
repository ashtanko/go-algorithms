package insert_interval

import "sort"

func insertInterval(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)

	//append and sort
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	res = append(res, intervals[0])

	//process the merge
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1]) //update result slice case
		} else {
			res = append(res, intervals[i]) //append result slice case
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
