package main

import "fmt"

func main() {
	m := 3
	n := 7
	fmt.Println(uniquePaths(m, n))
}

func uniquePaths(m int, n int) int {
	var cur []int
	for i := 0; i < n; i++ {
		cur = append(cur, 1)
	}
	for k := 1; k < m; k++ {
		for j := 1; j < n; j++ {
			cur[j] += cur[j-1]
			fmt.Println(j)
			ft.Println(cur[j])
		}
	}

	return cur[n-1]
}
