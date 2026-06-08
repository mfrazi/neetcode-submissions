func kthDistinct(arr []string, k int) string {
	count := map[string]int{}

	for _, val := range arr {
		if _, found := count[val]; !found {
			count[val] = 1
		} else {
			count[val]++
		}
	}

	for _, val := range arr {
		if c, _ := count[val]; c == 1 {
			k--
			if k == 0 {
				return val
			}
		}
	}

	return ""
}