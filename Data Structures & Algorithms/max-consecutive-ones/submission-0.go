func findMaxConsecutiveOnes(nums []int) int {
	count, maximumValue := 0, 0

	for _, num := range nums {
		if num == 0 {
			if count > maximumValue {
				maximumValue = count
			}
			count = 0
		} else {
			count++
		}
	}

	return max(count, maximumValue)
}
