package questions

import (
	"reflect"
	"testing"
)

func setZeroes(matrix [][]int) {

	rows, columns := make(map[int]bool, len(matrix)), make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				columns[j] = true
			}
		}
	}

	for r := range rows {
		for c := range matrix[r] {
			matrix[r][c] = 0
		}
	}
	for c := range columns {
		for r := range matrix {
			matrix[r][c] = 0
		}
	}
}

func TestSetZeroes(t *testing.T) {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	expect := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	setZeroes(matrix)
	if !reflect.DeepEqual(matrix, expect) {
		t.Logf("actual matrix: %v", matrix)
	}
}
