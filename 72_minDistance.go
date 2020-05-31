package main

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	var dp [][]int

	for i := 0; i <= m; i++ {
		var tmp []int
		num := 0
		for j := 0; j <= n; j++ {
			if i == 0 {
				num = j
			}

			if j == 0 {
				num = i
			}

			tmp = append(tmp, num)
		}
		dp = append(dp, tmp)
	}

	var cost int
	for i := 1; i <= m; i++ {
		for j := 1; j < n; j++ {
			cost = 1
			dp[i][j] = dp[i-1][j] + 1

			if dp[i-1][j]+1 < dp[i][j] {
				dp[i][j] = dp[i-1][j] + 1
			}

			if word1[i-1] == word2[j-1] {
				cost = 0
			}

			if cost+dp[i-1][j-1] < dp[i][j] {
				dp[i][j] = cost + dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}
