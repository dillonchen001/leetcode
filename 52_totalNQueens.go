package main

import "fmt"

func main() {
	n := 4
	fmt.Println(totalNQueens(n))
}

func totalNQueens(n int) int {
	var ret int
	dfs(n, 0, 0, 0, 0, &ret)
	return ret
}

func dfs(n, row, col, ld, rd int, ret *int) {
	if row >= n {
		*ret++
		return
	}

	var pick int
	bits := ^(col | ld | rd) & ((1 << n) - 1)
	for bits > 0 {
		pick = bits & -bits
		dfs(n, row+1, col|pick, (ld|pick)<<1, (rd|pick)>>1, ret)
		bits &= bits - 1
	}
}
