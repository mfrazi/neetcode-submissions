func makeEqual(words []string) bool {
    n := len(words)
	count := [26]int{}
	for _, word := range words {
		for i:=0; i<len(word); i++ {
			count[word[i]-'a']++
		}
	}
	for i:=0; i<26; i++ {
		if count[i]%n !=0 {
			return false
		}
	}

	return true
}