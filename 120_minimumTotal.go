package main

import "fmt"

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}

func minimumTotal(triangle [][]int) int {
	level := len(triangle)

	if level == 0 {
		return 0
	}
	if level == 1 {
		return triangle[0][0]
	}
	var dp [][]int
	dp = append(dp, []int{triangle[0][0]})

	result := 0
	for i := 1; i < level; i++ {
		var tmp []int
		if i == level-1 {
			result = dp[i-1][0] + triangle[i][0]
		} else {
			tmp = append(tmp, dp[i-1][0]+triangle[i][0])
		}
		for j := 1; j < i; j++ {
			num := 0
			if dp[i-1][j-1] < dp[i-1][j] {
				num = dp[i-1][j-1] + triangle[i][j]
			} else {
				num = dp[i-1][j] + triangle[i][j]
			}

			if i == level-1 {
				if num < result {
					result = num
				}
			} else {
				tmp = append(tmp, num)
			}
		}

		if i == level-1 {
			if dp[i-1][i-1]+triangle[i][i] < result {
				result = dp[i-1][i-1] + triangle[i][i]
			}
		} else {
			tmp = append(tmp, dp[i-1][i-1]+triangle[i][i])
			dp = append(dp, tmp)
		}
	}

	return result
}
