func isPathCrossing(path string) bool {
	currentPoint := [2]int{0, 0}
	visited := map[[2]int]struct{}{
		currentPoint: {},
	}

	for _, p := range path {
		switch p {
		case 'N':
			currentPoint[1]++
		case 'S':
			currentPoint[1]--
		case 'E':
			currentPoint[0]++
		case 'W':
			currentPoint[0]--
		}

		if _, found := visited[currentPoint]; found {
			return true
		}

		visited[currentPoint] = struct{}{}
	}

	return false
}