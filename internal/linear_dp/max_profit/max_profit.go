package max_profit

func maxProfit(prices []int) int {
	minValue, profit := 0, 0
	for i, price := range prices {
		if i == 0 {
			minValue = price
			continue
		}

		if diff := price - minValue; diff > profit {
			profit = diff
		}
		if price < minValue {
			minValue = price
		}
	}
	return profit
}
