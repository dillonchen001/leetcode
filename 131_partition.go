package main

import "fmt"

func main() {
	s := "aab"
	fmt.Println(partition(s))
}

func partition(s string) [][]string {
	ls := len(s)
	var result [][]string
	var tmp []string
	dfs(0, ls-1, s, tmp, &result)
	return result
}

func dfs(start int, end int, s string, tmp []string, result *[][]string) {
	if start > end {
		var tmpSli []string
		for _, v := range tmp {
			tmpSli = append(tmpSli, v)
		}
		*result = append(*result, tmpSli)
		return
	}

	for i := start; i <= end; i++ {
		tmpStr := s[start : i+1]
		if isPalindrome(tmpStr) {
			tmp = append(tmp, tmpStr)
			dfs(i+1, end, s, tmp, result)
			tmp = tmp[0 : len(tmp)-1]
		}
	}
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
