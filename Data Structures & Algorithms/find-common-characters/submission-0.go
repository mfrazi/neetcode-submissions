func commonChars(words []string) []string {
	setResult := make([][26]int, len(words))

	for i, word := range words {
		setResult[i] = [26]int{}
		for j := 0; j < len(word); j++ {
			setResult[i][int(word[j]-'a')]++
		}
	}

	result := []string{}
	for i := 0; i < 26; i++ {
		totalCommon := 1000
		for j := 0; j < len(setResult); j++ {
			totalCommon = min(totalCommon, setResult[j][i])
		}
		for n := 0; n < totalCommon; n++ {
			result = append(result, string('a'+i))
		}
	}

	return result
}