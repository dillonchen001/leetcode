package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 1, 2}
	sort.Ints(nums)
	fmt.Println(permuteUnique(nums))
}

func permuteUnique(nums []int) [][]int {
	ln := len(nums)
	var result [][]int

	if ln == 0 {
		return result
	}

	if ln == 1 {
		result = append(result, nums)
		return result
	}

	lastRet := permuteUnique(nums[0 : ln-1])
	fmt.Println(lastRet)

	for _, tmp := range lastRet {
		for j := ln - 1; j >= 0; j-- {
			var ttmp []int
			ttmp = append(ttmp, tmp[0:j]...)
			ttmp = append(ttmp, nums[ln-1])

			ttmp = append(ttmp, tmp[j:ln-1]...)
			result = append(result, ttmp)
			if j > 0 && tmp[j-1] == nums[ln-1] {
				break
			}
		}
	}
	return result
}
