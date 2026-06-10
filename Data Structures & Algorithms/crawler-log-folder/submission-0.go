func minOperations(logs []string) int {
	depth := 0
	for _, log := range logs {
		if log == "./" {
			continue
		}
		if log == "../" {
			depth--
			if depth < 0 {
				depth = 0
			}
			continue
		}
		depth++
	}
	return depth
}
