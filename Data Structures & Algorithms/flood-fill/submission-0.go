func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    originalColor := image[sr][sc]
    if originalColor == color {
        return image
    }

    queue := [][2]int{{sr, sc}}
    image[sr][sc] = color

    for len(queue) > 0 {
        x, y := queue[0][0], queue[0][1]

        if x > 0 && image[x-1][y] == originalColor {
            image[x-1][y] = color
            queue = append(queue, [2]int{x-1, y})
        }
        if x < len(image)-1 && image[x+1][y] == originalColor {
            image[x+1][y] = color
            queue = append(queue, [2]int{x+1, y})
        }
        if y > 0 && image[x][y-1] == originalColor {
            image[x][y-1] = color
            queue = append(queue, [2]int{x, y-1})
        }
        if y < len(image[0])-1 && image[x][y+1] == originalColor {
            image[x][y+1] = color
            queue = append(queue, [2]int{x, y+1})
        }

        queue = queue[1:]
    }

    return image
}

// [1,1,1]
// [1,1,0]
// [1,0,1]
