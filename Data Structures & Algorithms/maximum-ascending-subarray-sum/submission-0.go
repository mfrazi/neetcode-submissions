func maxAscendingSum(nums []int) int {
    result := 0

	for i:=0; i<len(nums); i++ {
		sum := nums[i]
		for i < len(nums)-1 && nums[i] < nums[i+1] {
			sum += nums[i+1]
			i++
		}
		result = max(result, sum)
	}

	return result
}