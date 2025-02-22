package kids_with_candies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var resultSlice = make([]bool, len(candies))
	var maximum int
	for i, j := range candies {
		if i == 0 || maximum < j {
			maximum = j
		}
	}

	for c := range candies {
		if candies[c]+extraCandies >= maximum {
			resultSlice[c] = true
		} else {
			resultSlice[c] = false
		}
	}
	return resultSlice
}
