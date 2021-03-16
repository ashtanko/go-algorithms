package max_profit

func maxProfitFee(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	min, profit := prices[0], 0

	for _, price := range prices {
		if price < min {
			min = price
		} else if price > min+fee {
			profit += price - fee - min
			min = price - fee
		}
	}
	return profit
}
