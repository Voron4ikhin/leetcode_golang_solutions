package _121_best_time_to_buy_sell_stocks

func MaxProfit(prices []int) int {
	var profit int
	var lowest = prices[0]
	for _, price := range prices {
		if price < lowest {
			lowest = price
		}
		if profit < price-lowest {
			profit = price - lowest
		}
	}

	return profit
}
