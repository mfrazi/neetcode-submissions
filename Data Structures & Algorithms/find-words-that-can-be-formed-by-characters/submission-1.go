func countCharacters(words []string, chars string) int {
    countChars := [26]int{}
	for i:=0; i<len(chars); i++ {
		countChars[chars[i]-'a']++
	}
	
	result := 0
	for _, word := range words {
		good := true
		countWord := [26]int{}
		word = strings.TrimSpace(word)
		for i:=0; i<len(word); i++ {
			c := word[i]-'a'
			countWord[c]++
			if countWord[c] > countChars[c] {
				good = false
				break
			}
		}
		if good {
			result += len(word)
		}
	}

	return result
}