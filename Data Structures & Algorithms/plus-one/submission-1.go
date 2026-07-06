func plusOne(digits []int) []int {
	i := len(digits)-1
	carry := 0
	
	for i >=0 && (i == len(digits)-1 || carry > 0) {
		digits[i]++
		if digits[i] > 9 {
			carry = 1
			digits[i] = 0
		} else {
			carry = 0
		}
		i--
	}

	if carry > 0 {
		return append([]int{1}, digits...)
	}

	return digits

}
