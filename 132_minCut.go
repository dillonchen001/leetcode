package main

func minCut(s string) int {
	ls := len(s)
	if ls < 2 {
		return 0
	}

	dp := make([]int, ls)
	for i := 0; i < ls; i++ {
		dp[i] = i
	}

	for i := 1; i < ls; i++ {
		if isPalindrome(s[0 : i+1]) {
			dp[i] = 0
			continue
		}

		for j := 0; j < i; j++ {
			if isPalindrome(s[j+1:i+1]) && dp[j]+1 < dp[i] {
				dp[i] = dp[j] + 1

			}
		}
	}
	return dp[ls-1]
}

func isPalindrome(s string) bool {
	ls := len(s)
	if ls <= 1 {
		return true
	}

	start := 0
	end := ls - 1

	for start < end {
		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}
	return true
}
