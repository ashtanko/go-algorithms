package koko_eating_bananas

// 875. Koko Eating Bananas
// https://leetcode.com/problems/koko-eating-bananas/description/
func minEatingSpeed(piles []int, hours int) int {
	left := 1
	right := 1000000000
	for left < right {
		mid := (left + right) / 2
		totalHours := 0
		for _, pile := range piles {
			totalHours += (pile + mid - 1) / mid
		}
		if totalHours > hours {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
