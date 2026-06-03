func groupAnagrams(strs []string) [][]string {
	groups := map[[26]int][]string{}

	for _, str := range strs {
		group := [26]int{}
		for _, c := range str {
			group[c-'a']++
		}
		if _, found := groups[group]; !found {
			groups[group] = []string{str}
		} else {
			groups[group] = append(groups[group], str)
		}
	}

	result := [][]string{}
	for _, val := range groups {
		result = append(result, val)
	}

	return result
}
