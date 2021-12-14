package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func readCoords() []Coords {
	coords := make([]Coords, 0)
	for {
		var x, y int
		_, err := fmt.Scanf("%d,%d", &x, &y)
		if err != nil {
			break
		}
		coords = append(coords, Coords{x, y})

	}
	return coords
}

type Fold struct {
	pos        int
	horizontal bool
}

func readFolds() []Fold {
	folds := make([]Fold, 0)
	for {
		var fold string
		_, err := fmt.Scanf("fold along %s", &fold)
		parts := strings.Split(fold, "=")
		if err != nil {
			break
		}
		position, _ := strconv.Atoi(parts[1])
		folds = append(folds, Fold{position, parts[0] == "y"})
	}
	return folds
}

func calculateFold(coords []Coords, fold Fold) {
	if fold.horizontal {
		for i := range coords {
			if coords[i].y > fold.pos {
				coords[i].y = 2*fold.pos - coords[i].y
			}
		}
	} else {
		for i := range coords {
			if coords[i].x > fold.pos {
				coords[i].x = 2*fold.pos - coords[i].x
			}
		}

	}
}

func calculateAllFolds(coords []Coords, folds []Fold) {
	for _, fold := range folds {
		calculateFold(coords, fold)
	}
}

func coordsLeftFold(coords []Coords, fold Fold) int {
	calculateFold(coords, fold)
	max_x, max_y := maxDimensions(coords)
	coord_map := make([][]int, max_y+1)
	for i := range coord_map {
		coord_map[i] = make([]int, max_x+1)
	}

	dots := len(coords)
	for _, coord := range coords {
		if coord_map[coord.y][coord.x] == 1 {
			dots--
		}
		coord_map[coord.y][coord.x] = 1
	}

	return dots
}

func maxDimensions(coords []Coords) (int, int) {
	x, y := 0, 0
	for _, coord := range coords {
		if coord.x > x {
			x = coord.x
		}
		if coord.y > y {
			y = coord.y
		}
	}

	return x, y
}

func buildMapFormCoords(coords []Coords) [][]int {
	max_x, max_y := maxDimensions(coords)
	coord_map := make([][]int, max_y+1)
	for i := range coord_map {
		coord_map[i] = make([]int, max_x+1)
	}

	for _, coord := range coords {
		coord_map[coord.y][coord.x] = 1
	}

	return coord_map

}

func PrintCoords(coords []Coords) {
	for _, row := range buildMapFormCoords(coords) {
		for _, value := range row {
			if value == 0 {
				fmt.Print("  ")
			} else {
				fmt.Print("# ")
			}
		}
		fmt.Println()
	}
}

func SolveDay13() {
	coords := readCoords()
	folds := readFolds()
	fmt.Println("Part 1: ", coordsLeftFold(coords, folds[0]))
	calculateAllFolds(coords, folds[1:])
	PrintCoords(coords)

}
