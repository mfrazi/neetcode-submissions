func appendCharacters(s string, t string) int {
    ps, pt := 0, 0

    for ; ps < len(s) && pt < len(t); ps++ {
        if s[ps] == t[pt] {
            pt++
        }
    }

    return len(t) - pt
}