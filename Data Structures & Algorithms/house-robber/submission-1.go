func rob(nums []int) int {
    if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	nums[1] = max(nums[0], nums[1])
	for i:=2; i<len(nums); i++ {
		nums[i] = max(nums[i-1], nums[i-2]+nums[i])
	}
	
	return nums[len(nums)-1]
}

// 8 1 1  8  9
// 8 8 9 16 18

// 2,9,8,3,6
// 2,9,10,12,