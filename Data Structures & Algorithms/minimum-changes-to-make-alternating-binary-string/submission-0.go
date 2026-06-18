func minOperations(s string) int {
    zeroOne, oneZero := 0, 1
	zeroOneCount, oneZeroCount := 0, 0

	for i:=0; i<len(s); i++ {
		c := int(s[i] - '0')
		if zeroOne != c {
			zeroOneCount++
		}
		if oneZero != c {
			oneZeroCount++
		}

		zeroOne ^= 1
		oneZero ^= 1
	}

	return min(zeroOneCount, oneZeroCount)
}