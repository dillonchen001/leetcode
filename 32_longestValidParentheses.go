package main

func longestValidParentheses(s string) int {
	ln := len(s)
	if ln < 2 {
		return 0
	}
	return myMax(cal(s, 0, 1, ln, '('), cal(s, ln-1, -1, -1, ')'))
}

func cal(s string, start int, flag int, end int, cTem byte) int {
	sum := 0
	currLen := 0
	max := 0
	validLen := 0
	for i := start; i != end; i += flag {
		if s[i] == cTem {
			sum += 1
		} else {
			sum += -1
		}
		currLen++

		if sum < 0 {
			max = myMax(max, validLen)
			sum = 0
			currLen = 0
			validLen = 0
		}

		if sum == 0 {
			validLen = currLen
		}
	}

	return myMax(max, validLen)
}

func myMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
