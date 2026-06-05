func isAlienSorted(words []string, order string) bool {
    originalOrder := map[byte]byte{}
    for i:=0; i<len(order); i++ {
        originalOrder[order[i]] = byte(i + 'a')
    }

    for i:=0; i<len(words); i++ {
    	var newWord strings.Builder
        for j:=0; j<len(words[i]); j++ {
            newWord.WriteByte(originalOrder[words[i][j]])
        }
        words[i] = newWord.String()
    }

    for i:=1; i<len(words); i++ {
        if words[i] < words[i-1] {
            return false
        }
    }

    return true
}