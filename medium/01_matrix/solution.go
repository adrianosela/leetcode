package _01_matrix

import "math"

func updateMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	ret := make([][]int, rows)
	for pos := range ret {
		ret[pos] = make([]int, cols)
	}

	maxInt := int(^uint(0) >> 1)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			ret[i][j] = maxInt
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// every time we find a zero, we
			// update our matrix
			if matrix[i][j] == 0 {
				updateMatrixAt(ret, i, j)
			}
		}
	}

	return ret
}

func updateMatrixAt(ret [][]int, i, j int) {
	for a := 0; a < len(ret); a++ {
		for b := 0; b < len(ret[0]); b++ {
			dist := int(math.Abs(float64(i-a)) + math.Abs(float64(j-b)))
			if dist < ret[a][b] {
				ret[a][b] = dist
			}
		}
	}
}
