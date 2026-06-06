func longestMonotonicSubarray(nums []int) int {
    result := 1

	for i:=0; i<len(nums); i++ {
		cnt := 1
		for i < len(nums)-1 && nums[i] < nums[i+1] {
			i++
			cnt++
		}
		result = max(result, cnt)
	}

	for i:=0; i<len(nums); i++ {
		cnt := 1
		for i < len(nums)-1 && nums[i] > nums[i+1] {
			i++
			cnt++
		}
		result = max(result, cnt)
	}

	return result
}
