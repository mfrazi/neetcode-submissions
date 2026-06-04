func reverseString(s []byte) {
    for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
        tmp := s[left]
        s[left] = s[right]
        s[right] = tmp
    }
}
