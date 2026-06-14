func maxNumberOfBalloons(text string) int {
    // balon
	count := [5]int{}
	for _, c := range text {
		switch c {
			case 'b':
				count[0] += 2
			case 'a':
				count[1] += 2
			case 'l':
				count[2] += 1
			case 'o':
				count[3] += 1
			case 'n':
				count[4] += 2
		}
	}
	minVal := count[0]
	for i:=1; i<5; i++ {
		if count[i] < minVal {
			minVal = count[i]
		}
	}
	return minVal/2
}