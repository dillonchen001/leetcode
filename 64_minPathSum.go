package main

import "fmt"

func main() {
	grid := [][]int{{1, 2, 4, 3}, {5, 6, 4, 4}}
	fmt.Println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}

	dp := make([][]int, rows, rows)
	dp[0] = append(dp[0], grid[0][0])
	for j := 1; j < cols; j++ {
		dp[0] = append(dp[0], dp[0][j-1]+grid[0][j])
	}

	fmt.Println(dp)
	for i := 1; i < rows; i++ {
		dp[i] = append(dp[i], dp[i-1][0]+grid[i][0])
		for j := 1; j < cols; j++ {
			if dp[i-1][j] > dp[i][j-1] {
				dp[i] = append(dp[i], dp[i][j-1]+grid[i][j])
			} else {
				dp[i] = append(dp[i], dp[i-1][j]+grid[i][j])
			}
		}
	}
	return dp[rows-1][cols-1]
}
