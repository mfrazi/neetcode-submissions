func minOperations(s string) int {
    missCount, position := 0, 0

	for i:=0; i<len(s); i++ {
		c := int(s[i] - '0')
		if position != c {
			missCount++
		}

		position ^= 1
	}

	return min(missCount, len(s)-missCount)
}