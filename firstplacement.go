package main

import "os"

func Initialpositions(piece int, tetrominoesList []Tetrominoes, size int) {
	if piece == len(tetrominoesList) {
		PrintBoard()
		os.Exit(0)
	}
	tetromino := tetrominoesList[piece]
	for y := 0; y < size-len(tetromino.form)+1; y++ {
		for x := 0; x < size-len(tetromino.form[0])+1; x++ {
			if Validateposition(y, x, piece, tetrominoesList) {
				Initialpositions(piece+1, tetrominoesList, size)
				Clearnonused(y, x, piece, tetrominoesList)
			}
		}
	}
	if piece == 0 {
		increaseSize := size + 1
		CreateBoard(increaseSize)
		Initialpositions(0, tetrominoesList, increaseSize)
	}
}
