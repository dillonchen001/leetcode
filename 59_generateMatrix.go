package main

import "fmt"

func main() {
	n := 4
	fmt.Println(generateMatrix(n))
}

func generateMatrix(n int) [][]int {
	var ret [][]int
	flag := n % 2
	serial := 0
	//n为奇数时，中心点是一个点，n为偶数时，中心点是4个数
	for i := 0; i < n; i++ {
		var tmp []int
		for j := 0; j < n; j++ {
			//i, j离中心点的距离最大值为层数,根据层数获取序号
			if flag == 1 {
				serial = getOddSerial(i, j, n)
			} else {
				serial = getEvenSerial(i, j, n)
			}

			tmp = append(tmp, serial+1)
		}
		ret = append(ret, tmp)
	}
	return ret
}

func getOddSerial(i, j, n int) int {
	level := 0
	lser := 1
	mid := n / 2
	serial := 0
	if i <= mid {
		level = mid - i
		lser = 1
	} else {
		level = i - mid
		lser = 3
	}

	if j <= mid {
		if mid-j > level {
			level = mid - j
			lser = 4
		}
	} else {
		if j-mid > level {
			level = j - mid
			lser = 2
		}
	}

	for m := n / 2; m > level; m-- {
		serial += 4 * (m * 2)
	}

	if lser == 1 {
		serial += j - (n/2 - level)
	} else if lser == 2 {
		serial += level*2 + i - (n/2 - level)
	} else if lser == 3 {
		serial += level*4 + (n - 1 - j) - (n/2 - level)
	} else {
		serial += level*6 + (n - 1 - i) - (n/2 - level)
	}
	return serial
}

func getEvenSerial(i, j, n int) int {
	level := 0
	lser := 1
	min := n/2 - 1
	max := n / 2
	serial := 0
	if i <= min {
		level = min - i
		lser = 1
	} else {
		level = i - max
		lser = 3
	}

	if j <= min {
		if min-j > level {
			level = min - j
			lser = 4
		}
	} else {
		if j-max > level {
			level = j - max
			lser = 2
		}
	}

	for m := n/2 - 1; m > level; m-- {
		serial += 4 * (m*2 + 1)
	}

	if lser == 1 {
		serial += j - (n/2 - 1 - level)
	} else if lser == 2 {
		serial += (level*2 + 1) + i - (n/2 - 1 - level)
	} else if lser == 3 {
		serial += (level*2+1)*2 + (n - 1 - j) - (n/2 - 1 - level)
	} else {
		serial += (level*2+1)*3 + (n - 1 - i) - (n/2 - 1 - level)
	}

	if i == 1 && j == 1 {
		fmt.Println(level)
		fmt.Println(lser)
		fmt.Println(serial)
	}

	return serial
}
