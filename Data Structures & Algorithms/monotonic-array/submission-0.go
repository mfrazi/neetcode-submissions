func isMonotonic(nums []int) bool {
    if len(nums) == 1 {
		return true
	}
	sign := 0
	for i:=1; i<len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff == 0 {
			continue
		}
		if sign > 0 && diff < 0 ||
			sign < 0 && diff > 0 {
				return false
			}
		sign = diff
	}
	return true
}