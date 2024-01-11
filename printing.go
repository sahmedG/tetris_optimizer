package main

import "fmt"

func PrintBoard() {
	for i := 0; i < len(BOARD); i++ {
		for k := 0; k < len(BOARD); k++ {
			if BOARD[i][k] == 0 {
				fmt.Print(".")
				continue
			}
			fmt.Print(string(BOARD[i][k]))
		}
		fmt.Println()
	}
}
