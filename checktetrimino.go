package main


func Validateposition(y int, x int, piece int, tetrominoesList []Tetrominoes) bool {
	for i := 0; i < len(tetrominoesList[piece].form); i++ {
		if len(tetrominoesList[piece].form)+y > len(BOARD) || len(tetrominoesList[piece].form[i])+x > len(BOARD) {
			return false
		}
	}

	for a := y; a < (len(tetrominoesList[piece].form) + y); a++ {
		for b := x; b < (len(tetrominoesList[piece].form[a-y]) + x); b++ {
			if tetrominoesList[piece].form[a-y][b-x] == 46 {
				continue
			}
			if BOARD[a][b] == 0 {
				PlacePiece(a, b, piece, tetrominoesList)
				// BOARD[a][b] = byte(65 + piece)
			} else {
				Clearnonused(y, x, piece, tetrominoesList)
				return false
			}
		}
	}
	return true
}
