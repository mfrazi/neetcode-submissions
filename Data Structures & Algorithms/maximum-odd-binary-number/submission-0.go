func maximumOddBinaryNumber(s string) string {
	count := [2]int{}
	for _, c := range s {
		count[c-'0']++
	}
	result := ""
	for i:=0; i<count[1]-1; i++ {
		result += "1"
	}
	for i:=0; i<count[0]; i++ {
		result += "0"
	}
	return result + "1"
}
