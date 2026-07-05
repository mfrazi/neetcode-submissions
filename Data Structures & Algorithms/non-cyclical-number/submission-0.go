func isHappy(n int) bool {
    visited := map[int]struct{}{}
	for n != 1 {
		if _, found := visited[n]; found {
			return false
		}
		visited[n] = struct{}{}

		sum := 0
		for n > 0 {
			sum += (n%10) * (n%10)
			n /= 10
		}

		n = sum
	}

	return true
}
