package main

func CreateBoard(size int) {
	BOARD = nil
	BOARD = make([][]byte, size)
	for i := range BOARD {
		BOARD[i] = make([]byte, size)
	}
}
