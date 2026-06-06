func isArraySpecial(nums []int) bool {
    if len(nums) == 1 {
		return true
	}
	check := nums[0] % 2
	for i:=1; i<len(nums); i++ {
		check = (check+1)%2
		if nums[i]%2 != check {
			return false
		}
	}
	return true
}