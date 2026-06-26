func countStudents(students []int, sandwiches []int) int {
	count := [2]int{}
	for _, student := range students {
		count[student]++
	}

	for _, sandwich := range sandwiches {
		count[sandwich]--
		if count[sandwich] < 0 {
			return count[0] + count[1] + 1
		}
	}

	return 0
}

// [1,1,1,0,0,1]
// [1,0,0,0,1,1]

// [1,1,0,0,1]
// [0,0,0,1,1]

// [1,0,0,1,1]
// [0,0,0,1,1]

// [0,0,1,1,1]
// [0,0,0,1,1]

// [1,1,1]
// [0,1,1]