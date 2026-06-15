func largestGoodInteger(num string) string {
	current := -1
	for i:=0; i<len(num)-2; i++ {
		if num[i] == num[i+1] && num[i+1] == num[i+2] {
			if int(num[i]-'0') > current {
				current = int(num[i]-'0')
			}
			i+=2
		}
	}
	if current == -1 {
		return ""
	}
	return fmt.Sprintf("%d%d%d", current, current, current)
}
