package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			Contents = append(Contents, scanner.Bytes())
		}

		if !Validatetetrimino(Contents) {
			fmt.Println("ERROR")
			return
		}

		tetrominoesList := Lookformatch(Contents)
		if tetrominoesList == nil {
			fmt.Println("ERROR")
			return
		}

		minSize := InitialBoard(tetrominoesList)
		CreateBoard(minSize)
		Initialpositions(0, tetrominoesList, minSize)

	} else {
		fmt.Println("wrong number of arguments!\nexample: ./tetris input.txt")
	}
}
