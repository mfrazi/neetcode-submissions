func twoSum(nums []int, target int) []int {
    index := map[int]int{}
	for right, num := range nums {
		if left, found := index[target-num]; found {
			return []int{left, right}
		}
		index[num] = right
	}
	return []int{}
}
