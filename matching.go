package main

func Lookformatch(transformedContent [][]byte) []Tetrominoes {
	var tetrominoToAppend Tetrominoes
	var tetrominoesList []Tetrominoes
	tetrisCounter := 1

	for a := 0; a < len(transformedContent); a++ {
		for b := 0; b < len(transformedContent[a]); b++ {
			var found bool = false
			for i := 0; i < len(Valids); i++ {
				var skip bool = false
				for k := 0; k < len(Valids[i]); k++ {
					for m := 0; m < len(Valids[i][k]); m++ {
						if a+k == tetrisCounter*5-1 || b+m >= 4 {
							skip = true
							break
						}
						if transformedContent[a+k][b+m] == Valids[i][k][m] {
							continue
						}
						skip = true
						break
					}
					if skip {
						break
					}
				}
				if !skip {
					found = true
					tetrominoToAppend.form = Valids[i]
					tetrominoesList = append(tetrominoesList, tetrominoToAppend)
					break
				}
			}
			if found {
				if a+(tetrisCounter*5-a) >= len(transformedContent) {
					a = len(transformedContent) - 1
				} else {
					a = a + (tetrisCounter*5 - a) - 1
				}
				tetrisCounter++
				break
			}

			if a == (tetrisCounter*5)-2 && b == 3 {
				return nil
			}
		}
	}

	return tetrominoesList
}
