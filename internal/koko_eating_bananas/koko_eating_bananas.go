package koko_eating_bananas

// 875. Koko Eating Bananas
// https://leetcode.com/problems/koko-eating-bananas/description/
func minEatingSpeed(piles []int, h int) int {
	l := 1
	r := 1000000000
	for l < r {
		m := (l + r) / 2
		total := 0
		for i := range piles {
			p := piles[i]
			total += (p + m - 1) / m
		}
		if total > h {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
