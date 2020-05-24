package main

func isValidSudoku(board [][]byte) bool {
	mrow := make(map[int]map[byte]int)
	mcol := make(map[int]map[byte]int)
	mcel := make(map[int]map[byte]int)

	for i := 0; i < 9; i++ {
		tmp := map[byte]int{
			'1': 0,
			'2': 0,
			'3': 0,
			'4': 0,
			'5': 0,
			'6': 0,
			'7': 0,
			'8': 0,
			'9': 0,
		}
		mrow[i] = tmp

		tmp1 := map[byte]int{
			'1': 0,
			'2': 0,
			'3': 0,
			'4': 0,
			'5': 0,
			'6': 0,
			'7': 0,
			'8': 0,
			'9': 0,
		}
		mcol[i] = tmp1

		tmp2 := map[byte]int{
			'1': 0,
			'2': 0,
			'3': 0,
			'4': 0,
			'5': 0,
			'6': 0,
			'7': 0,
			'8': 0,
			'9': 0,
		}
		mcel[i] = tmp2
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			tByte := board[i][j]
			if tByte == '.' {
				continue
			}

			if mrow[i][tByte] > 0 {
				return false
			} else {
				mrow[i][tByte]++
			}

			if mcol[j][tByte] > 0 {
				return false
			} else {
				mcol[j][tByte]++
			}

			if mcel[i/3*3+j/3][tByte] > 0 {
				return false
			} else {
				mcel[i/3*3+j/3][tByte]++
			}

		}
	}
	return true
}
