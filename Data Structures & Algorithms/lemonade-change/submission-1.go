func lemonadeChange(bills []int) bool {
	c5, c10 := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			c5++
		} else if bill == 10 {
			c10++
			c5--
			if c5 < 0 {
				return false
			}
		} else {
			if c10 >= 1 && c5 >= 1 {
				c10--
				c5--
			} else if c5 >= 3 {
				c5 -= 3
			} else {
				return false
			}
		}
	}

	return true
}
