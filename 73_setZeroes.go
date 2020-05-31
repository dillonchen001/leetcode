package main

import "fmt"

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	m := len(matrix)

	if m == 0 {
		return
	}

	n := len(matrix[0])
	if n == 0 {
		return
	}

	lm := make(map[int]bool)
	ln := make(map[int]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				lm[i] = true
				ln[j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if _, ok := lm[i]; ok {
				matrix[i][j] = 0
			}

			if _, ok := ln[j]; ok {
				matrix[i][j] = 0
			}
		}
	}

	return

}
