package solutions

import (
	"fmt"
)

type Line struct {
	start_x int
	start_y int
	end_x   int
	end_y   int
}

func maxCoords(lines []Line) (int, int) {
	var max_x, max_y int = 0, 0
	for _, line := range lines {
		if line.start_x > max_x {
			max_x = line.start_x
		}
		if line.start_y > max_y {
			max_y = line.start_y
		}
		if line.end_x > max_x {
			max_x = line.end_y
		}
		if line.end_y > max_y {
			max_y = line.end_y
		}
	}
	return max_x, max_y
}

func genBoard(lines []Line) [][]int {
	max_x, max_y := maxCoords(lines)
	board := make([][]int, 0)
	for i := 0; i < max_x+1; i++ {
		board = append(board, make([]int, max_y+1))
	}
	fmt.Println(max_x, max_y, len(board), len(board[0]))
	return board
}

func readLines() []Line {
	lines := make([]Line, 0)
	for {
		line := Line{}
		n_read, err := fmt.Scanf("%d,%d -> %d,%d", &line.start_x, &line.start_y, &line.end_x, &line.end_y)
		if n_read == 4 {
			lines = append(lines, line)
		} else {
			fmt.Println(n_read, err)
			break
		}
	}
	return lines
}

func step(start int, end int) int {
	if start > end {
		return -1
	} else if start < end {
		return 1
	}
	return 0
}

func fillLine(line Line, board [][]int) int {
	step_x := step(line.start_x, line.end_x)
	step_y := step(line.start_y, line.end_y)
	madeDangerous := 0

	for x, y := line.start_x, line.start_y; x != line.end_x || y != line.end_y; {
		board[x][y]++
		if board[x][y] == 2 {
			madeDangerous++
		}

		x += step_x
		y += step_y
	}

	board[line.end_x][line.end_y]++
	if board[line.end_x][line.end_y] == 2 {
		madeDangerous++
	}
	return madeDangerous
}

func dangerousZones(lines []Line, excludeDiagonal bool) int {
	board := genBoard(lines)
	danger := 0
	for _, line := range lines {
		if !excludeDiagonal || (line.start_x == line.end_x || line.start_y == line.end_y) {
			danger += fillLine(line, board)
		}
	}
	return danger
}

func SolveDay5() {
	lines := readLines()
	fmt.Println("Part 1: ", dangerousZones(lines, true))
	fmt.Println("Part 2: ", dangerousZones(lines, false))
}
