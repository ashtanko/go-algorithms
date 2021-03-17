package largest_altitude

import "github.com/ashtanko/go-algorithms/utils"

func largestAltitude(gain []int) int {
	maxAlt := 0
	currAlt := 0
	for i := 0; i < len(gain); i++ {
		currAlt += gain[i]
		maxAlt = utils.Max(currAlt, maxAlt)
	}
	return maxAlt
}
