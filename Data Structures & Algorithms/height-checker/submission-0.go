func heightChecker(heights []int) int {
    count := [100]int{}

	for _, height := range heights {
		count[height-1]++
	}

	unexpected, iter := 0, 0
	for i:=0; i<100; i++ {
		for j:=0; j<count[i]; j++ {
			if heights[iter] != i+1 {
				unexpected++
			}
			iter++
		}
	}

	return unexpected
}