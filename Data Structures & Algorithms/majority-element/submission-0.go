func majorityElement(nums []int) int {
	currentMajority, power := nums[0], 1

	for i:=1; i<len(nums); i++ {
		num := nums[i]
		if num == currentMajority {
			power++
			continue
		}

		power--
		if power < 0 {
			currentMajority = num
			power = 1
		}
	}

	return currentMajority
}
