func countConsistentStrings(allowed string, words []string) int {
    allowedChar := [26]bool{}
	for i:=0; i<len(allowed); i++ {
		allowedChar[allowed[i]-'a'] = true
	}

	result := 0
	for _, word := range words {
		i := 0
		for ; i<len(word); i++ {
			c := word[i]-'a'
			if !allowedChar[c] {
				break
			}
		}
		if i == len(word) {
			result++
		}
	}

	return result
}