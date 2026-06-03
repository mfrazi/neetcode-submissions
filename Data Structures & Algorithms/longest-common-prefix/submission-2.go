func longestCommonPrefix(strs []string) string {
    if len(strs) == 1 {
		return strs[0]
	}
	
	var commonPrefix strings.Builder
	index := 0
	
	for {
		for i:=1; i<len(strs); i++ {
			if index == len(strs[i]) || index == len(strs[i-1]) ||
				strs[i][index] != strs[i-1][index] {
				return commonPrefix.String()
			}
		}

		commonPrefix.WriteByte(strs[0][index])
		index++
	}

	return commonPrefix.String()
}
