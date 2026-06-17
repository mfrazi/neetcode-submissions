func findErrorNums(nums []int) []int {
    n := len(nums)
	duplicate, sum := 0, 0
	maxVal := int(float64(n+1) / float64(2) * float64(n))
	for i:=0; i<n; i++ {
		placeIndex := nums[i]
		if placeIndex < 0 {
			placeIndex *= -1
		}
		sum += placeIndex
		placeIndex--

		if nums[placeIndex] < 0 {
			duplicate = placeIndex + 1
		} else {
			nums[placeIndex] *= -1
		}
	}

	return []int{duplicate, maxVal-(sum-duplicate)}

}