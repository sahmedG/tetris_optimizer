package main

func PlacePiece(x int, y int, piece int, tetrominoesList []Tetrominoes) {
	tetromino := tetrominoesList[piece].form
	for a := 0; a < len(tetromino); a++ {
		for b := 0; b < len(tetromino[a]); b++ {
			if tetromino[a][b] == 35 {
				BOARD[x][y] = byte(65 + piece)
			}
		}
	}
}
