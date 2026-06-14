func check(nums []int) bool {
    firstVal := nums[0]
	foundRotate := false
	for i:=1; i<len(nums); i++ {
		if foundRotate && nums[i] > firstVal {
			return false
		}

		if nums[i] >= nums[i-1] {
			continue
		}

		if foundRotate {
			return false
		}

		foundRotate = true
	}

	return true
}