package main

import "fmt"

func main() {
	n := 5
	fmt.Println(generate(n))
}

func generate(numRows int) [][]int {
	var result [][]int
	if numRows <= 0 {
		return result
	}

	result = append(result, []int{1})
	if numRows == 1 {
		return result
	}

	result = append(result, []int{1, 1})
	if numRows == 2 {
		return result
	}

	for i := 3; i <= numRows; i++ {
		var tmp []int
		tmp = append(tmp, 1)
		fmt.Println(result)
		for j := 1; j <= i/2; j++ {
			num := result[i-2][j-1] + result[i-2][j]
			tmp = append(tmp, num)
		}
		fmt.Println(tmp)
		for j := i/2 + 1; j < i; j++ {
			num := tmp[i-1-j]
			tmp = append(tmp, num)
		}
		result = append(result, tmp)

	}
	return result
}
