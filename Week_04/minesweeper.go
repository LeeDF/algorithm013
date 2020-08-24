package Week_04

import "strconv"

func updateBoard(board [][]byte, click []int) [][]byte {
	x := click[0]
	y := click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}
	_dfs(board, x, y)
	return board
}

func _dfs(board [][]byte, x, y int) {
	xs := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	ys := []int{-1, -1, -1, 0, 0, 1, 1, 1}

	nu := 0
	for i := range xs {
		tx := x + xs[i]
		ty := y + ys[i]
		if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) {
			continue
		}
		if board[tx][ty] == 'M' {
			nu++
		}
	}

	if nu > 0 {
		board[x][y] = []byte(strconv.Itoa(nu))[0]
	} else {
		board[x][y] = 'B'
		for i := range xs {
			tx := x + xs[i]
			ty := y + ys[i]
			if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) || board[tx][ty] != 'E' {
				continue
			}
			_dfs(board, tx, ty)
		}
	}
}
