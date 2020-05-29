package main

import "fmt"

func main() {
	matrix := [][]int{{1, 11}, {2, 12}, {3, 13}, {4, 14}, {5, 15}, {6, 16}, {7, 17}, {8, 18}, {9, 19}, {10, 20}}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	var result []int
	m := len(matrix)
	if m == 0 {
		return result
	}
	n := len(matrix[0])

	if n == 0 {
		return result
	}

	//分层
	for i := 0; i < (m+1)/2 && i < (n+1)/2; i++ {
		if m-2*i == 1 {
			for j := 0; j < n-i*2; j++ {
				result = append(result, matrix[i][i+j])
			}
			break
		}

		if n-2*i == 1 {
			for j := 0; j < m-i*2; j++ {
				result = append(result, matrix[i+j][i])
			}
			break
		}

		//上
		for j := 0; j < n-1-i*2; j++ {
			result = append(result, matrix[i][i+j])
		}

		fmt.Println(n)
		fmt.Println(i)
		fmt.Println("hahah")
		//右
		for j := 0; j < m-1-i*2; j++ {
			result = append(result, matrix[i+j][n-1-i])
		}

		//下
		for j := 0; j < n-1-i*2; j++ {
			result = append(result, matrix[m-1-i][n-1-i-j])
		}

		//左
		for j := 0; j < m-1-i*2; j++ {
			result = append(result, matrix[m-1-j][i])
		}
	}
	return result
}
