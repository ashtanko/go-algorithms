package distribute_candies

import (
	"math"
	"sort"
)

func distributeCandies(candyType []int) int {
	sort.Ints(candyType)
	uniqueCandies := 1
	for i := 1; i < len(candyType) && uniqueCandies/2 < len(candyType); i++ {
		if candyType[i] != candyType[i-1] {
			uniqueCandies++
		}
	}
	return int(math.Min(float64(uniqueCandies), float64(len(candyType)/2)))
}
