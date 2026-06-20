func intersection(nums1 []int, nums2 []int) []int {
    flag := map[int]struct{}{}
	for _, num := range nums1 {
		if _, found := flag[num]; !found {
			flag[num] = struct{}{}
		}
	}

	result := []int{}
	for _, num := range nums2 {
		if _, found := flag[num]; found {
			result = append(result, num)
			delete(flag, num)
		}
	}

	return result
}
