package main

import "fmt"

func main() {
	n := 4
	fmt.Println(getRow(n + 1))
}

func getRow(rowIndex int) []int {
	var result []int
	if rowIndex == 1 {
		return []int{1, 1}
	}
	result = append(result, 1)
	for i := 1; i <= rowIndex/2; i++ {
		num := 1
		for j := 1; j <= i; j++ {
			num = result[j-1] * (rowIndex + 1 - j) / (j)
		}
		result = append(result, num)
	}

	for i := rowIndex/2 + 1; i <= rowIndex; i++ {
		num := result[rowIndex-i]
		result = append(result, num)
	}
	return result
}
