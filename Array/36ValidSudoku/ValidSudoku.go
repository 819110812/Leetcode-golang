package _6ValidSudoku

func isValidSudoku(board [][]byte) bool {
	row := make([][]bool, 9)
	col := make([][]bool, 9)
	box := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = make([]bool, 9)
		col[i] = make([]bool, 9)
		box[i] = make([]bool, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - '1'
			k := i/3*3 + j/3
			if row[i][num] || col[j][num] || box[k][num] {
				return false
			} else {
				row[i][num] = true
				col[j][num] = true
				box[k][num] = true
			}
		}
	}
	return true
}
