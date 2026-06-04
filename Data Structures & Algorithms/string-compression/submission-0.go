func compress(chars []byte) int {
	total := 0

	for i := 0; i < len(chars); i++ {
		countChar := 1

		for i < len(chars)-1 && chars[i] == chars[i+1] {
			countChar++
			i++
		}

		chars[total] = chars[i]
		total++

		if countChar == 1 {
			continue
		}

		tmp := []byte{}
		for countChar > 0 {
			tmp = append(tmp, byte(countChar%10+'0'))
			countChar /= 10
		}
		for j:=len(tmp)-1; j>=0; j-- {
			chars[total] = tmp[j]
			total++
		}
	}

	return total
}