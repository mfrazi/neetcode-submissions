func rob(nums []int) int {
    a, b := 0, 0
	for _, num := range nums {
		a, b = b, max(b, a+num)
	}
	
	return b
}

// 8 1 1  8  9
// 8 8 9 16 18

// 2,9,8,3,6
// 2,9,10,12,