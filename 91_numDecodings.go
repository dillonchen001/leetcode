package main

import "fmt"

func numDecodings(s string) int {
	ls := len(s)

	if s[0] == '0' {
		return 0
	}
	pre := 1
	curr := 1

	for i := 1; i < ls; i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				curr = pre
			} else {
				return 0
			}
		} else if s[i] <= byte('0'+6) {
			if s[i-1] == '1' || s[i-1] == '2' {
				pre, curr = curr, pre+curr
			} else {
				pre = curr
			}
		} else {
			if s[i-1] == '1' {
				pre, curr = curr, pre+curr
			} else {
				pre = curr
			}
		}
	}

	return curr
}
