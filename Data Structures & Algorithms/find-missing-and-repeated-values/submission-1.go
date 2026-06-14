func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
	totalSum := int(float64(n*n+1)/float64(2)*float64(n*n))
	sum := 0
	double := 0
	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			val := grid[i][j]
			if val < 0 {
				val *= -1
			}

			sum += val
			val--

			nextI, nextJ := val/n, val%n
			if grid[nextI][nextJ] > 0 {
				grid[nextI][nextJ] *= -1
			} else {
				double = grid[i][j]
				if double < 0 {
					double *= -1
				}
			}
		}
	}

	return []int{double, totalSum - (sum - double)}
}

// 0 1 2
// 3 4 5
// 6 7 8

// 1 2 3
// 4 5 6
// 7 8 9

// 0 = 0 0 
// 1 = 0 1
// 2 = 0 2
// 3 = 1 0
// 4 = 1 1
// 5 = 1 2
// 6 = 2 0
// 7 = 2 1
// 8 = 2 2