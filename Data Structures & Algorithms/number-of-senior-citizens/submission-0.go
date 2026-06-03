func countSeniors(details []string) int {
    total := 0

	for _, d := range details {
		if (d[11] - '0') * 10 + (d[12] - '0') > 60 {
			total++
		}
	}

	return total
}
