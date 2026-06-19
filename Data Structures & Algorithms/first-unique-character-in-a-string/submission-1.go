func firstUniqChar(s string) int {
	count := [26]int{}
	for i:=0; i<len(s); i++ {
		c := int(s[i]-'a')
		if count[c] < 0 {
			continue
		}
		if count[c] > 0 {
			count[c] *= -1
			continue
		}
		count[c] = i+1
	}

	minIndex := len(s) + 1

	for i:=0; i<26; i++ {
		if count[i] > 0 && minIndex > count[i] {
			minIndex = count[i]
		}
	}

	if minIndex == len(s) + 1 {
		return -1
	}

	return minIndex - 1
}
