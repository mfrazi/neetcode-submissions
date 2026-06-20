func maxLengthBetweenEqualCharacters(s string) int {
    left := [26]int{}
	right := [26]int{}
	for i, j := 0, len(s)-1; i<len(s); i, j = i+1, j-1 {
		if left[s[i]-'a'] == 0 {
			left[s[i]-'a'] = i+1
		}
		if right[s[j]-'a'] == 0 {
			right[s[j]-'a'] = j+1
		}
	}
	result := 0
	for i:=0; i<26; i++ {
		if left[i] > 0 && right[i] > 0 && right[i]-left[i] > result {
			result = right[i]-left[i]
		}
	}

	return result-1
}