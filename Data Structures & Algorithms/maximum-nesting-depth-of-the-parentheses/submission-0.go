func maxDepth(s string) int {
	result := 0
	stack := 0

	for i:=0; i<len(s); i++ {
		if s[i] == '(' {
			stack++
		} else if s[i] == ')' {
			stack--
		}

		result = max(result, stack)
	}

	return result
}
