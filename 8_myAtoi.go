package main

func myAtoi(str string) int {
	len1 := len(str)

	if len1 < 1 {
		return 0
	}

	INT_MAX := 1<<31 - 1
	INT_MIN := -1 << 31

	start := 0
	flag := 0
	result := 0

	for i := 0; i < len1; i++ {
		if str[i] >= '0' && str[i] <= '9' {
			tmpNum, _ := strconv.Atoi(string(str[i]))
			if flag == 1 {
				if result < (INT_MIN+tmpNum)/10 {
					return INT_MIN
				}
				result = result*10 - tmpNum
			} else {
				if result > (INT_MAX-tmpNum)/10 {
					return INT_MAX
				}
				result = result*10 + tmpNum
			}
			start = 1
		} else {
			if start == 0 {
				if str[i] == ' ' {
					continue
				} else if str[i] == '+' {
					start = 1
				} else if str[i] == '-' {
					start = 1
					flag = 1
				} else {
					return result
				}
			} else {
				return result
			}
		}
	}
	return result
}
