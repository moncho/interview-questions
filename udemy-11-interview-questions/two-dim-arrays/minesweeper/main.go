package minesweeper

type gameMap [][]int

func mineSweeper(bombs [][]int, rows, columns int) gameMap {
	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, columns)
	}
	for _, bomb := range bombs {
		result[bomb[0]][bomb[1]] = -1
		updateSurroundingTiles(result, bomb[0], bomb[1])
	}

	return result

}

func updateSurroundingTiles(m gameMap, row, column int) {
	var currentRow int
	var currentCol int

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			currentRow = row + i
			currentCol = column + j
			if currentRow < 0 || currentRow >= len(m) ||
				currentCol < 0 || currentCol >= len(m[currentRow]) ||
				m[currentRow][currentCol] == -1 {
				continue
			} else {
				m[currentRow][currentCol]++
			}
		}
	}
}
