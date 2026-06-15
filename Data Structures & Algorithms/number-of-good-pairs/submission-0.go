func numIdenticalPairs(nums []int) int {
    count := [100]int{}
	for _, num := range nums {
		count[num-1]++
	}
	sum := 0

	for i:=0; i<100; i++ {
		val := count[i]-1
		sum += int(float64(val + 1 ) / float64(2) * float64(val))
	}

	return sum
}