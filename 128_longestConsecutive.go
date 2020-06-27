package main

import "fmt"

func main() {
	nums := []int{1, 2, 0, 1}
	fmt.Println(nums)
	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
	hashNum := make(map[int]int)
	result := 0
	for _, num := range nums {
		if _, ok := hashNum[num]; ok {
			continue
		}

		left := 0
		right := 0

		if llen, ok := hashNum[num-1]; ok {
			left = llen
		}

		if rlen, ok := hashNum[num+1]; ok {
			right = rlen
		}

		curlen := left + right + 1
		if curlen > result {
			result = curlen
		}

		hashNum[num] = curlen
		hashNum[num-left] = curlen
		hashNum[num+right] = curlen
	}

	return result
}
