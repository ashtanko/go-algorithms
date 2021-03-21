package reordered_power_of_two

import "reflect"

func reorderedPowerOf2(N int) bool {
	if N < 1 {
		return false
	}
	arr := count(N)
	for i := 0; i < 31; i++ {
		if reflect.DeepEqual(arr, count(1<<i)) {
			return true
		}
	}
	return false
}

func count(N int) []int {
	ans := make([]int, 10)
	for N > 0 {
		ans[N%10]++
		N /= 10
	}
	return ans
}
