func longestPalindrome(s string) int {
	counter := map[rune]int{}
	for _, c := range s {
		counter[c]++
	}

	maximumLength := 0
	addOne := false
	for _, val := range counter {
		maximumLength += val

		if val % 2 == 1 {
			maximumLength--

			if !addOne {
				addOne = true
			}
		}
	}
	
	if addOne {
		maximumLength++
	}

	return maximumLength
}