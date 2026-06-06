func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextGreater := map[int]int{}
	stack := []int{0}

	for i:=1; i<len(nums2); i++ {
		for len(stack) > 0 && nums2[stack[len(stack)-1]] < nums2[i] {
			nextGreater[nums2[stack[len(stack)-1]]] = nums2[i]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	for len(stack) > 0 {
		nextGreater[nums2[stack[len(stack)-1]]] = -1
		stack = stack[:len(stack)-1]
	}

	result := make([]int, len(nums1))
	for i:=0; i<len(nums1); i++ {
		result[i] = nextGreater[nums1[i]]
	}

	return result
}
