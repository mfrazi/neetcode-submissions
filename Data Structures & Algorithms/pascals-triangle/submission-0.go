func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}
	for i := 2; i <= numRows; i++ {
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				result[i-1] = append(result[i-1], 1)
				continue
			}
			result[i-1] = append(result[i-1], result[i-2][j-1]+result[i-2][j])
		}
	}

	return result
}