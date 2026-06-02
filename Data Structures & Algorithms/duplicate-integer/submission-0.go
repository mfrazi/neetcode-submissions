func hasDuplicate(nums []int) bool {
    exist := map[int]struct{}{}
    for _, num := range nums {
        if _, found := exist[num]; found {
            return true
        }
        exist[num] = struct{}{}
    }
    return false
}
