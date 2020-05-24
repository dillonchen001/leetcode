package main

func solveSudoku(board [][]byte) {
	row := make(map[int]map[byte]bool)
	col := make(map[int]map[byte]bool)
	block := make(map[int]map[byte]bool)
	var tbyte byte
	for i := 0; i < 9; i++ {
		tmp := make(map[byte]bool)
		tmp1 := make(map[byte]bool)
		tmp2 := make(map[byte]bool)

		for j := 1; j <= 9; j++ {
			tbyte = byte('0' + j)
			tmp[tbyte] = true
			tmp1[tbyte] = true
			tmp2[tbyte] = true
		}

		row[i] = tmp
		col[i] = tmp1
		block[i] = tmp2
	}

	var empty [][]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				tbyte = board[i][j]
				delete(row[i], tbyte)
				delete(col[j], tbyte)
				delete(block[(i/3)*3+j/3], tbyte)
			} else {
				empty = append(empty, []int{i, j})
			}
		}
	}
	backTrack(0, empty, row, col, block, board)
}

func backTrack(iterm int, empty [][]int, row, col, block map[int]map[byte]bool, board [][]byte) bool {
	if len(empty) == iterm {
		return true
	}

	i := empty[iterm][0]
	j := empty[iterm][1]
	b := (i/3)*3 + j/3

	var tbyte byte
	for k := 1; k <= 9; k++ {
		tbyte = byte('0' + k)
		if _, ok := row[i][tbyte]; !ok {
			continue
		}

		if _, ok := col[j][tbyte]; !ok {
			continue
		}

		if _, ok := block[b][tbyte]; !ok {
			continue
		}

		delete(row[i], tbyte)
		delete(col[j], tbyte)
		delete(block[b], tbyte)
		board[i][j] = tbyte
		if backTrack(iterm+1, empty, row, col, block, board) {
			return true
		}

		row[i][tbyte] = true
		col[j][tbyte] = true
		block[b][tbyte] = true
		board[i][j] = '.'
	}
	return false
}
