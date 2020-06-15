package main

import "fmt"

func main() {
	s := "02852"
	fmt.Println(restoreIpAddresses(s))
}

func restoreIpAddresses(s string) []string {
	ls := len(s)

	var result []string
	if ls < 4 || ls > 12 {
		return result
	}

	var tmp string
	dps(tmp, s, 4, 0, ls-1, &result)
	return result
}

func dps(tmp string, s string, n int, start int, end int, result *[]string) {
	if n*3 < end-start+1 || n > end-start+1 {
		return
	}

	if n == 0 {
		*result = append(*result, tmp)
		return
	}

	for i := 1; i <= 3 && start+i <= end+1; i++ {
		if isValidIpNum(s, start, i) {
			for j := 0; j < i; j++ {
				tmp = tmp + string(s[start+j])
			}
			flag := 0
			if n > 1 {
				flag = 1
				tmp = tmp + "."
			}
			dps(tmp, s, n-1, start+i, end, result)
			tmp = tmp[0 : len(tmp)-i-flag]
		}
	}
}

func isValidIpNum(s string, start int, i int) bool {
	if s[start] == '0' && i > 1 {
		return false
	}

	if s[start] == '2' && i == 3 {
		if s[start+1] == '5' && s[start+2] > '5' {
			return false
		}

		if s[start+1] > '5' {
			return false
		}
	}

	if s[start] >= '3' && i == 3 {
		return false
	}
	return true
}
