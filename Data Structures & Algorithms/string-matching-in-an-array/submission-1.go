func stringMatching(words []string) []string {
	setResult := map[string]struct{}{}
	for i:=0; i<len(words); i++ {
		for j:=0; j<len(words); j++ {
			if i == j {
				continue
			}
			if strings.Contains(words[i], words[j]) {
				setResult[words[j]] = struct{}{}
			}
		}
	}

	result := []string{}
	for key, _ := range setResult {
		result = append(result, key)
	}

	return result
}
