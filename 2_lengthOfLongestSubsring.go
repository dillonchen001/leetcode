package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
}

func lengthOfLongestSubstring(s string) int {
	location := make(map[rune]int)

	start := 0
	max := 0
	for i, ch := range s {
		if lastI, ok := location[ch]; ok && start <= lastI {
			start = lastI + 1
		}

		if i-start+1 > max {
			max = i - start + 1
		}

		location[ch] = i
	}
	return max
}
