func findLucky(arr []int) int {
	count := [500]int{}

	for _, a := range arr {
		count[a-1]++
	}

	result := -1
	for i:=0; i<500; i++ {
		if count[i] == i+1 {
			result = i+1
		}
	}

	return result
}
