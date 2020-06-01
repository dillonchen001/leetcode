package main

import "fmt"

func main() {
	n := 4
	k := 2
	fmt.Println(combine(n, k))
}

func combine(n int, k int) [][]int {
	var ret [][]int
	var curr []int
	backtrack(1, n, k, curr, &ret)
	return ret
}

func backtrack(first int, n int, k int, curr []int, ret *[][]int) {
	if len(curr) == k {
		var tmp []int
		for _, k := range curr {
			tmp = append(tmp, k)
		}
		*ret = append(*ret, tmp)
	}

	for i := first; i <= n; i++ {
		curr = append(curr, i)
		backtrack(i+1, n, k, curr, ret)
		curr = curr[0 : len(curr)-1]
	}
}
