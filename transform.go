package main

import "bytes"

func SwitchTetromino(content []byte) [][]byte{
	sep := []byte{13,10}
	return bytes.Split(content, sep)
}
