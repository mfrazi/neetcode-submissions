func islandPerimeter(grid [][]int) int {
    perimeter := 0

    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            if grid[i][j] != 1{
                continue
            }
            
            if i == 0 || grid[i-1][j] == 0 {
                perimeter++
            } 
            if i == len(grid)-1 || grid[i+1][j] == 0 {
                perimeter++
            }
            if j == 0 || grid[i][j-1] == 0 {
                perimeter++
            }
            if j == len(grid[0])-1 || grid[i][j+1] == 0 {
                perimeter++
            }
        }
    }


    return perimeter
}
