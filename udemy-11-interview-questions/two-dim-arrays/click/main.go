package click

import (
	"fmt"
	"strconv"
)

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

func clickBFS(game gameMap, i, j int) gameMap {
	if game[i][j] != 0 {
		return game
	}
	queue := []string{pos(i, j)}
	return visit(game, queue)
}

func visit(game gameMap, queue []string) gameMap {
	if len(queue) == 0 {
		return game
	}
	rows := len(game)
	columns := len(game[0])
	i, j := parse(queue[0])
	for row := i - 1; row <= i+1; row++ {
		for column := j - 1; column <= j+1; column++ {
			if row < 0 || row >= rows || column < 0 || column >= columns || game[row][column] != 0 {
				continue
			} else {
				game[row][column] = -2
				queue = append(queue, pos(row, column))
			}
		}
	}
	return visit(game, queue[1:])
}
func pos(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func parse(s string) (int, int) {
	i, _ := strconv.Atoi(s[:1])
	j, _ := strconv.Atoi(s[2:])
	return i, j
}
