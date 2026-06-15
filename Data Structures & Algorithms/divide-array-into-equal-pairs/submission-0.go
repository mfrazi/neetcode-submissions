func divideArray(nums []int) bool {
    count := [500]int{}

	for _, num := range nums {
		count[num-1]++
	}
	for i:=0; i<500; i++ {
		if count[i]%2!=0 {
			return false
		}
	}
	return true
}