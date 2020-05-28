package main

import "fmt"

func main() {
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	lm := len(matrix)

	if lm <= 1 {
		return
	}

	tmp := 0
	for i := 0; i < lm/2; i++ {
		//4个角交换
		for j := 0; j < lm-1-i*2; j++ {
			tmp = matrix[i][i+j]
			matrix[i][i+j] = matrix[lm-1-i-j][i]
			matrix[lm-1-i-j][i] = matrix[lm-1-i][lm-1-i-j]
			matrix[lm-1-i][lm-1-i-j] = matrix[i+j][lm-1-i]
			matrix[i+j][lm-1-i] = tmp
		}
		fmt.Println(matrix)
	}
}
