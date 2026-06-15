func getRow(rowIndex int) []int {
    result := []int{1}

	for i:=1; i<=rowIndex; i++ {
		tmpResult := make([]int, len(result)+1)
		for j:=0; j<=i; j++ {
			if j == 0 || j == i {
				tmpResult[j] = 1
				continue
			}
			tmpResult[j] = result[j-1]+result[j]
		}
		result = tmpResult
	}

	return result
}