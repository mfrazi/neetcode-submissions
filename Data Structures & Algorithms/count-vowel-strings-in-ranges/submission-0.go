func vowelStrings(words []string, queries [][]int) []int {
    prefixSum := make([]int, len(words))
	if startEndVowel(words[0]) {
		prefixSum[0] = 1
	}
	for i:=1; i<len(words); i++ {
		if startEndVowel(words[i]) {
			prefixSum[i] = prefixSum[i-1] + 1
		} else {
			prefixSum[i] = prefixSum[i-1]
		}
	}

	result := []int{}
	for _, query := range queries {
		if query[0] == 0 {
			result = append(result, prefixSum[query[1]])
			continue
		}
		result = append(result, prefixSum[query[1]]-prefixSum[query[0]-1])
	}
	return result
}

func startEndVowel(word string) bool {
	flag := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	_, start := flag[word[0]]
	_, end := flag[word[len(word)-1]]
	
	return start && end
}