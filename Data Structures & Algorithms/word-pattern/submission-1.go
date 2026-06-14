func wordPattern(pattern string, s string) bool {
    listOfWord := strings.Split(s, " ")
	
	if len(listOfWord) != len(pattern) {
		return false
	}

	mapWord := map[byte]string{}
	wordExist := map[string]struct{}{}

	for i:=0; i<len(pattern); i++ {
		word, foundWord := mapWord[pattern[i]]
		_, isWordUsed := wordExist[listOfWord[i]]

		if !foundWord {
			if isWordUsed {
				return false
			}

			mapWord[pattern[i]] = listOfWord[i]
			wordExist[listOfWord[i]] = struct{}{}
			continue
		}
		
		if word != listOfWord[i] {
			return false
		}
	}

	return true
}