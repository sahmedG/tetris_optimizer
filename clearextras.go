package main

func Clearnonused(y int, x int, piece int, tetrominoesList []Tetrominoes) {
	for a := y; a < (len(tetrominoesList[piece].form) + y); a++ {
		for b := x; b < (len(tetrominoesList[piece].form[a-y]) + x); b++ {
			if (tetrominoesList[piece].form[a-y][b-x]) == 46 {
				continue
			}
			if BOARD[a][b] == byte(65+piece) {
				BOARD[a][b] = 0
			}
		}
	}
}
