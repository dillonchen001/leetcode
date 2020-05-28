package main

func reverse(x int) int {
	MAX_NUM := 1<<31 - 1
	MIN_NUM := -1 << 32

	tmp := 0
	ins := 0
	for x != 0 {
		if tmp > (MAX_NUM-x%10)/10 || tmp < (MIN_NUM-x%10)/10 {
			return 0
		}
		tmp *= 10
		ins = x % 10
		x = x / 10
		tmp += ins
	}
}
