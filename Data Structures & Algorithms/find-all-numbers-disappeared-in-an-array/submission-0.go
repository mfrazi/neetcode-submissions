func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		j := nums[i] - 1
		for j >= 0 && nums[j] != j+1 {
			nextJ := nums[j] - 1
			nums[j] = j + 1
			j = nextJ
		}
		if i != j {
			nums[i] = -1
		}
	}

	result := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == -1 {
			result = append(result, i+1)
		}
	}
	return result
}