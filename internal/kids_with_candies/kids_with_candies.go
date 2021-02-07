package kids_with_candies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var resultSlice = make([]bool, len(candies))
	var max int
	for i, j := range candies {
		if i == 0 || max < j {
			max = j
		}
	}

	for c := range candies {
		if candies[c]+extraCandies >= max {
			resultSlice[c] = true
		} else {
			resultSlice[c] = false
		}
	}
	return resultSlice
}
