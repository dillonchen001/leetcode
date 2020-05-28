package main

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)

	length := len(str)
	start := 0
	end := length - 1

	for start <= end {
		if str[start] != str[end] {
			return false
		} else {
			start++
			end--
		}
	}
	return true
}
