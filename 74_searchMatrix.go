package main

import "fmt"

func main() {
	target := 4
	matrix := [][]int{{1, 3, 5}}
	fmt.Println(searchMatrix(matrix, target))
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])
	if n == 0 {
		return false
	}

	start := 0
	end := m - 1
	mid := 0
	for start <= end {
		mid = (start + end) / 2
		if matrix[mid][0] > target {
			end = mid - 1
		} else if matrix[mid][0] == target {
			return true
		} else {
			if matrix[mid][n-1] < target {
				start = mid + 1
			} else if matrix[mid][n-1] == target {
				return true
			} else {
				i := 1
				j := n - 2
				fmt.Println(i)
				fmt.Println(j)

				rmid := 0
				for i <= j {
					rmid = (i + j) / 2
					fmt.Println(rmid)

					if matrix[mid][rmid] == target {
						return true
					} else if matrix[mid][rmid] > target {
						j = rmid - 1
					} else {
						i = rmid + 1
					}
				}

				return false
			}
		}

	}

	return false
}
