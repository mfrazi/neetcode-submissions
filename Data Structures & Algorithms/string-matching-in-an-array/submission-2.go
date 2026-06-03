func stringMatching(words []string) []string {
	result := make([]string, 0)

	for i := range words {
		for j := range words {
			if i == j || len(words[i]) > len(words[j]) {
				continue
			}

			if strings.Contains(words[j], words[i]) {
				result = append(result, words[i])
				break
			}
		}
	}

	return result
}