func rob(nums []int) int {
    if len(nums) == 1 {
		return nums[0]
	}
	
	travel := func(houses []int) int {
		a, b := 0, 0
		for i:=0; i<len(houses); i++ {
			a, b = b, max(b, a+houses[i])
		}

		return b
	}

	return max(travel(nums[1:]), travel(nums[:len(nums)-1]))
}