func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	for i:=0; i<len(nums); i++ {
		result[i] = 1
	}
	for i:=1; i<len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	product := nums[len(nums)-1]
	for i:=len(nums)-2; i>=0; i-- {
		result[i] *= product
		product *= nums[i]
	}

	return result
}
