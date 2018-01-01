package click

import "fmt"

type gameMap [][]int

func click(game gameMap, i, j int) gameMap {
	if game[i][j] == 0 {
		visited := make(map[string]struct{})
		return clickWithMemory(game, visited, i, j)
	}
	return game
}

func clickWithMemory(game gameMap, visited map[string]struct{}, i, j int) gameMap {

	rows := len(game)
	columns := len(game[0])
	if _, ok := visited[pos(i, j)]; ok || game[i][j] != 0 {
		return game
	}
	game[i][j] = -2
	visited[pos(i, j)] = struct{}{}

	for row := i - 1; row <= i+1; row++ {
		for column := j - 1; column <= j+1; column++ {
			if row < 0 || row >= rows || column < 0 || column >= columns {
				continue
			} else {
				game = clickWithMemory(game, visited, row, column)
			}
		}
	}
	return game
}

func pos(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
