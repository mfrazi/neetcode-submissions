func isPerfectSquare(num int) bool {
	left := 1
	right := num

	for left < right {
		mid := (left+right)/2
		val := mid*mid
		if val == num {
			return true
		}
		if val < num {
			left = mid+1
		} else {
			right = mid
		}
	}

	return left*left == num
}
