package main

import "fmt"

func main() {
	n := 4
	fmt.Println(solveNQueens(n))
}

func solveNQueens(n int) [][]string {
	inRow := make(map[int]int)
	inCol := make(map[int]int)
	inDale := make(map[int]int)
	inHill := make(map[int]int)
	var result [][]string
	ret := backTrack(0, n, inRow, inCol, inDale, inHill, result)
	return ret
}

func backTrack(row int, n int, inRow, inCol, inDale, inHill map[int]int, result [][]string) [][]string {
	for col := 0; col < n; col++ {
		if isValid(row, col, n, inCol, inDale, inHill) {
			inRow[row] = col
			inCol[col] = 1
			inDale[row-col+n] = 1
			inHill[row+col] = 1
			if row == n-1 {
				var arrTmp []string
				for i := 0; i < n; i++ {
					tmp := ""
					for j := 0; j < inRow[i]; j++ {
						tmp += "."
					}
					tmp += "Q"
					for j := inRow[i] + 1; j < n; j++ {
						tmp += "."
					}

					arrTmp = append(arrTmp, tmp)
				}
				result = append(result, arrTmp)
			} else {
				fmt.Println(inCol)
				return backTrack(row+1, n, inRow, inCol, inDale, inHill, result)
			}
			delete(inRow, row)
			delete(inCol, col)
			delete(inDale, row-col+n)
			delete(inHill, row+col)
		}
	}
	return result
}

func isValid(row, col, n int, inCol, inDale, inHill map[int]int) bool {
	if _, ok := inCol[col]; ok {
		return false
	}

	if _, ok := inDale[row-col+n]; ok {
		return false
	}

	if _, ok := inHill[row+col]; ok {
		return false
	}
	return true
}
