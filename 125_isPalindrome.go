package main

import "fmt"

func main() {
	s := "0P"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	ls := len(s)
	if ls <= 1 {
		return true
	}

	start := 0
	end := ls - 1

	for start < end {
		if !(s[start] >= 'a' && s[start] <= 'z') && !(s[start] >= 'A' && s[start] <= 'Z') && !(s[start] >= '0' && s[start] <= '9') {
			start++
			continue
		}

		if !(s[end] >= 'a' && s[end] <= 'z') && !(s[end] >= 'A' && s[end] <= 'Z') && !(s[end] >= '0' && s[end] <= '9') {
			end--
			continue
		}

		fmt.Println(s[start])
		fmt.Println(s[end])
		if s[start] == s[end] || (byte(s[start]+32) == s[end] && s[start] >= 'A') || (byte(s[end]+32) == s[start] && s[end] >= 'A') {
			start++
			end--
		} else {
			fmt.Println(ls)
			fmt.Println(start)
			fmt.Println(end)
			fmt.Println("haha")
			return false
		}
	}
	return true
}
