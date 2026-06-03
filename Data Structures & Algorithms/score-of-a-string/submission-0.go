func scoreOfString(s string) int {
    score := 0
    for i:=1; i<len(s); i++ {
        currentScore := int(s[i]) - int(s[i-1])
        if currentScore < 0 {
            currentScore *= -1
        }
        score += currentScore
    }
    return score
}
