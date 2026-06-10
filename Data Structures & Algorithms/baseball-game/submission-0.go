func calPoints(operations []string) int {
	scores := make([]int, len(operations))
	iscores := 0

	for _, operation := range operations {
		if operation == "+" {
			scores[iscores] = scores[iscores-1] + scores[iscores-2]
			iscores++
			continue
		}
		if operation == "D" {
			scores[iscores] = 2 * scores[iscores-1]
			iscores++
			continue
		}
		if operation == "C" {
			iscores--
			scores[iscores] = 0
			continue
		}

		s, _ := strconv.Atoi(operation)
		scores[iscores] = s
		iscores++
	}

	finalScore := 0
	for _, score := range scores {
		finalScore += score
	}

	return finalScore
}
