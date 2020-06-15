package main

func numDistinct(s string, t string) int {
	ls := len(s)
	lt := len(t)

	dp := make([][]int, ls+1)
	for i := 0; i <= ls; i++ {
		dp[i] = make([]int, lt+1)
		for j := 0; j <= lt; j++ {
			if j == 0 {
				dp[i][0] = 1
				continue
			}

			if i == 0 {
				dp[0][j] = 0
				continue
			}

			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[ls][lt]
}
