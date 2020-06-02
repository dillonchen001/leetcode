package main

import "fmt"

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCB"
	fmt.Println(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	var marked [][]bool
	m := len(board)
	if m == 0 {
		return false
	}

	n := len(board[0])
	for i := 0; i < m; i++ {
		var tmp []bool
		for j := 0; j < n; j++ {
			tmp = append(tmp, false)
		}

		marked = append(marked, tmp)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrace(board, word, 0, i, j, marked, m, n) {
				return true
			}
		}
	}
	return false
}

func backtrace(board [][]byte, word string, index int, startI int, startJ int, marked [][]bool, m int, n int) bool {
	if index == len(word)-1 {
		return board[startI][startJ] == word[index]
	}

	if board[startI][startJ] == word[index] {
		marked[startI][startJ] = true

		if startI > 0 && !marked[startI-1][startJ] && backtrace(board, word, index+1, startI-1, startJ, marked, m, n) {
			return true
		}

		if startJ+1 < n && !marked[startI][startJ+1] && backtrace(board, word, index+1, startI, startJ+1, marked, m, n) {
			return true
		}

		if startI+1 < m && !marked[startI+1][startJ] && backtrace(board, word, index+1, startI+1, startJ, marked, m, n) {
			return true
		}

		if startJ > 0 && !marked[startI][startJ-1] && backtrace(board, word, index+1, startI, startJ-1, marked, m, n) {
			return true
		}
	}

	marked[startI][startJ] = false
	return false
}
