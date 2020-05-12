package main

func intToRoman(num int) string {
	if num < 0 || num > 3999 {
		return "您输入的数字不合法，请输入1~3999"
	}

	m5 := make(map[int]string)
	m1 := make(map[int]string)

	m5[100] = "D"
	m5[10] = "L"
	m5[1] = "V"

	m1[1000] = "M"
	m1[100] = "C"
	m1[10] = "X"
	m1[1] = "I"

	reminder := 0
	binary := 1
	result := ""
	tmpStr := ""

	for num > 0 {
		reminder = num % 10
		num = num / 10
		tmpStr = ""

		if reminder == 9 {
			result = m1[binary] + m1[binary*10] + result
		} else if reminder >= 5 {
			tmpStr = tmpStr + m5[binary]
			for i := 0; i < reminder-5; i++ {
				tmpStr = tmpStr + m1[binary]
			}
			result = tmpStr + result
		} else if reminder == 4 {
			result = m1[binary] + m5[binary] + result
		} else {
			for i := 0; i < reminder; i++ {
				tmpStr = tmpStr + m1[binary]
			}
			result = tmpStr + result
		}

		binary = binary * 10
	}
	return result
}
