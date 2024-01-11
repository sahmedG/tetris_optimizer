package main

func Validatetetrimino(transformedContent [][]byte) bool {
	index := 0
	for i := 0; i < len(transformedContent); i++ {
		if len(transformedContent[i]) == 0 {
			if index != 4 {
				return false
			} else {
				index = 0
			}
		}
		for k := 0; k < len(transformedContent[i]); k++ {
			if transformedContent[i][k] == 35 {
				index++
			}
		}
	}
	return true
}
