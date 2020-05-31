package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 3
	k := 3
	fmt.Println(getPermutation(n, k))
}

func getPermutation(n int, k int) string {

	used := make(map[int]bool)
	var factory []int

	factory = append(factory, 1)
	for i := 1; i < n; i++ {
		factory = append(factory, factory[i-1]*i)
	}
	k--
	var m int
	var digit int
	var j int

	var builder strings.Builder
	for i := n - 1; i >= 0; i-- {
		m = k / factory[i]
		k = k % factory[i]
		digit = 0
		j = 0
		for j <= m {
			digit++
			if _, ok := used[digit]; ok {
				continue
			}
			j++
		}

		used[digit] = true
		builder.WriteString(strconv.Itoa(digit))
	}

	return builder.String()
}
