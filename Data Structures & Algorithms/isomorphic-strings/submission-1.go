func isIsomorphic(s string, t string) bool {
	pair := map[byte]byte{}
	isUsed := map[byte]struct{}{}

	for i:=0; i<len(s); i++ {
		value, found := pair[s[i]]
		if found {
			if value != t[i] {
				return false
			}
		} else {
			if _, used := isUsed[t[i]]; used {
				return false
			}
			pair[s[i]] = t[i]
			isUsed[t[i]] = struct{}{}
		}
	}
	return true
}
