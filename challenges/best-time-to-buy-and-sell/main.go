package main

func maxProfit(prices []int) int {
	count := len(prices)
    if count < 2 {
        return 0
    }

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices[1:] {
		if price < minPrice {
			minPrice = price
		} else {
			profit := price - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	
	return maxProfit
}