func replaceElements(arr []int) []int {
    maxVal := -1

    for i:=len(arr)-1; i>=0; i-- {
        tmp := arr[i]
        arr[i] = maxVal

        if tmp > maxVal {
            maxVal = tmp
        }
    }

    return arr
}
