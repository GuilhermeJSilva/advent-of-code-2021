package solutions

import (
	"fmt"
)

func updateAll(powerLevels [][]int) []Coords {
	overLimit := make([]Coords, 0)
	for row_i, row := range powerLevels {
		for col_i := range row {
			powerLevels[row_i][col_i]++
			if powerLevels[row_i][col_i] > 9 {
				overLimit = append(overLimit, Coords{row_i, col_i})
			}
		}
	}
	return overLimit
}

var adjacentRow [8]int = [8]int{1, 1, 1, -1, -1, -1, 0, 0}
var adjacentCol [8]int = [8]int{1, -1, 0, 1, -1, 0, 1, -1}

func updateDay(powerLevels [][]int) int {
	toFlash := updateAll(powerLevels)
	flashes := 0

	for len(toFlash) != 0 {
		coords := toFlash[len(toFlash)-1]
		toFlash = toFlash[:len(toFlash)-1]
		if powerLevels[coords.x][coords.y] != 0 {
			flashes++
			powerLevels[coords.x][coords.y] = 0
			for i := 0; i < len(adjacentCol); i++ {
				row := coords.x + adjacentRow[i]
				col := coords.y + adjacentCol[i]
				if row >= 0 && row < len(powerLevels) && col >= 0 && col < len(powerLevels) && powerLevels[row][col] != 0 {
					powerLevels[row][col]++
					if powerLevels[row][col] > 9 {
						toFlash = append(toFlash, Coords{row, col})
					}
				}
			}
		}
	}
	return flashes
}

func updateDays(powerLevels [][]int, days int) int {
	flashes := 0
	for i := 0; i < days; i++ {
		flashes += updateDay(powerLevels)
	}
	return flashes
}

func firstSimul(powerLevels [][]int) int {
	for i := 1; ; i++ {
		flashes := updateDay(powerLevels)
		if flashes == 100 {
			return i
		}
	}
}

func copyMatrix(src [][]int) [][]int {
	dst := make([][]int, len(src))

	for i := range dst {
		dst[i] = make([]int, len(src[i]))
		for j := range dst[i] {
			dst[i][j] = src[i][j]
		}
	}
	return dst
}

func SolveDay11() {
	powerLevels := readHeightmap()
	var powerLevels2 [][]int = copyMatrix(powerLevels)
	fmt.Println("Part 1:", updateDays(powerLevels, 100))
	fmt.Println("Part 2:", firstSimul(powerLevels2))
}
