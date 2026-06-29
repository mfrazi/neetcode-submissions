func buyChoco(prices []int, money int) int {
	l1, l2 := prices[0], prices[1]
	if l1 > l2 {
		l1, l2 = l2, l1
	}

	for i:=2; i<len(prices); i++ {
		if l2 > prices[i] {
			l2 = prices[i]
		}
		if l1 > l2 {
			l1, l2 = l2, l1
		}
	}

	if l1+l2 > money {
		return money
	}
	return money-(l1+l2)
}
