func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
    	return fmt.Sprintf("%d%d", nums[i], nums[j]) > fmt.Sprintf("%d%d", nums[j], nums[i])
	})

	result := ""
	notWrite := true
	for _, num := range nums {
		if notWrite && num == 0 {
			continue
		}
		notWrite = false
		result += fmt.Sprintf("%d",num)
	}

	if result == "" {
		return "0"
	}

	return result
}
