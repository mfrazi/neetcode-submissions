func arrangeCoins(n int) int {
	i, result := 1, 0
	for n > 0 {
		if n >= i {
			result++
		}
		n -= i
		i++
	}
	return result
}
