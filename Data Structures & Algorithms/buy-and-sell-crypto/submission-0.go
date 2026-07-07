func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxSell := 0
	for i:=1; i<len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		maxSell = max(maxSell, prices[i]-minPrice)
	}

	return maxSell
}
