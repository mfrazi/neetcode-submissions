func isSubsequence(s string, t string) bool {
    if len(s) > len(t) {
        return false
    }

    ps := 0
    for pt := 0; pt<len(t) && ps<len(s); pt++ {
        if s[ps] == t[pt] {
            ps++
        }
    }

    return ps == len(s)
}
