package main

func romanToInt(s string) int {
	m := make(map[rune]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	result := 0
	lastNum := 0
	for _, chr := range s {
		result = result + m[chr]
		if m[chr] > lastNum {
			result = result - lastNum*2
		}
		lastNum = m[chr]
	}
	return result
}
