func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		checkIndex := nums[i]
		if checkIndex < 0 {
			checkIndex = -checkIndex
		}

		if nums[checkIndex-1] > 0 {
			nums[checkIndex-1] *= -1
		}
	}

	result := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}