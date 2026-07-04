func setZeroes(matrix [][]int) {
	const flagZero = -1
	colZero, rowZero := false, false
    for i:=0; i<len(matrix); i++ {
		for j:=0; j<len(matrix[i]); j++ {
			if matrix[i][j] != 0 {
				continue
			}
			
			if i == 0 && j == 0 {
				colZero = true
				rowZero = true
			} else if i == 0 {
				rowZero = true
			} else if j == 0 {
				colZero = true	
			}

			matrix[0][j] = flagZero
			matrix[i][0] = flagZero
		}
	}

	for i:=1; i<len(matrix); i++ {
		for j:=1; j<len(matrix[i]); j++ {
			if matrix[0][j] == flagZero ||
				matrix[i][0] == flagZero {
					matrix[i][j] = 0
				}
		}
	}

	for i:=0; i<len(matrix); i++ {
		if matrix[i][0] == flagZero || colZero {
			matrix[i][0] = 0
		}
	}

	for i:=0; i<len(matrix[0]); i++ {
		if matrix[0][i] == flagZero || rowZero {
			matrix[0][i] = 0
		}
	}
}
