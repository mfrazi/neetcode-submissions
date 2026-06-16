func containsNearbyDuplicate(nums []int, k int) bool {
	index := map[int]int{}
	for i, num := range nums {
		valI, found := index[num]
		if found && i - valI <= k{
			return true
		}
		index[num] = i
	}
	return false
}
