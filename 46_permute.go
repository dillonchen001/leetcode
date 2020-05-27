package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	ln := len(nums)
	var result [][]int

	if ln == 0 {
		return result
	}

	if ln == 1 {
		result = append(result, nums)
		return result
	}

	lastRet := permute(nums[0 : ln-1])

	for _, tmp := range lastRet {
		for j := ln - 1; j >= 0; j-- {
			var ttmp []int
			ttmp = append(ttmp, tmp[0:j]...)
			ttmp = append(ttmp, nums[ln-1])

			ttmp = append(ttmp, tmp[j:ln-1]...)
			result = append(result, ttmp)
		}
	}
	return result
}
