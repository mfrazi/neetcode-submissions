func maxScore(s string) int {
	n := len(s)
    countOne := make([]int, n)
	if s[0] == '1' {
		countOne[0] = 1
	}

	for i:=1; i<n; i++ {
		if s[i] == '1' {
			countOne[i] = countOne[i-1] + 1
		} else {
			countOne[i] = countOne[i-1]
		}
	}

	result := 0

	for i:=1; i<n; i++ {
		one := countOne[n-1] - countOne[i-1]
		zero := i - countOne[i-1]
		if one + zero > result {
			result = one + zero
		}
	}

	return result
}