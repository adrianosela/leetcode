package search_a_2d_matrix_II

func searchMatrix(matrix [][]int, target int) bool {
	// if we have no rows or no columns, the target
	// will not exist in the matrix
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row := len(matrix) - 1
	col := 0
	for row >= 0 && col < len(matrix[0]) {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			row--
		} else {
			col++
		}
	}

	return false
}
