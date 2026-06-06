func maxDifference(s string) int {
	freq := [26]int{}
	for _, c := range s {
		freq[int(c-'a')]++
	}

	maxOdd, minEven := 0, 101
	for i:=0; i<26; i++ {
		if freq[i] == 0 {
			continue
		}

		if freq[i] % 2 == 1 {
			maxOdd = max(maxOdd, freq[i])
		} else {
			minEven = min(minEven, freq[i])
		}
	}

	return maxOdd-minEven
}
