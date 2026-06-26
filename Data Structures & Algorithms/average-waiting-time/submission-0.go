func averageWaitingTime(customers [][]int) float64 {
    totalWait := 0
	currentTime := 0

	for _, customer := range customers {
		if currentTime < customer[0] {
			currentTime = customer[0]
		}

		currentTime += customer[1]
		totalWait += currentTime - customer[0]
	}


	return float64(totalWait)/float64(len(customers))
}