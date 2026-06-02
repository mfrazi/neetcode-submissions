func replaceElements(arr []int) []int {
    n := len(arr)
    left, right, current := n-2, n-2, arr[n-1]
    result := make([]int, n)
    result[n-1] = -1

    for left >= 0 {
        if arr[left] > current || left == 0 {
            for right >= left {
                result[right] = current
                right--
            }
            current = arr[left]
        }
        
        left--
    }

    return result
}
