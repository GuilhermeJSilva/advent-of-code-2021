package solutions

import (
	"fmt"
	"sort"
)

func readHeightmapLine() ([]int, error) {
	var line string
	_, err := fmt.Scanln(&line)
	if err != nil {
		return nil, err
	}
	intLine := make([]int, len(line))
	for pos, char := range line {
		intLine[pos] = int(char) - '0'
	}
	return intLine, nil
}

func readHeightmap() [][]int {
	heightmap := make([][]int, 0)

	for {
		line, err := readHeightmapLine()
		if err == nil {
			heightmap = append(heightmap, line)
		} else {
			break
		}
	}

	return heightmap
}

func sumRiskLevel(heightmap [][]int) int {
	sum := 0

	for i, row := range heightmap {
		for j, value := range row {
			if (i <= 0 || heightmap[i-1][j] > value) && (i >= len(heightmap)-1 || heightmap[i+1][j] > value) && (j <= 0 || heightmap[i][j-1] > value) && (j >= len(row)-1 || heightmap[i][j+1] > value) {
				sum += value + 1
			}
		}
	}
	return sum
}

type Coords struct {
	x int
	y int
}

func getLowPoints(heightmap [][]int) []Coords {
	points := make([]Coords, 0)

	for i, row := range heightmap {
		for j, value := range row {
			if (i <= 0 || heightmap[i-1][j] > value) && (i >= len(heightmap)-1 || heightmap[i+1][j] > value) && (j <= 0 || heightmap[i][j-1] > value) && (j >= len(row)-1 || heightmap[i][j+1] > value) {
				points = append(points, Coords{i, j})
			}
		}
	}
	fmt.Println(points)
	return points

}

var x_offset [4]int = [4]int{1, -1, 0, 0}
var y_offset [4]int = [4]int{0, 0, 1, -1}

func flowsInto(heightmap [][]int, origin Coords, basin int, explored [][]int) bool {
	foundBasin := false

	for i := 0; i < 4; i++ {
		x := origin.x + x_offset[i]
		y := origin.y + y_offset[i]
		if (x >= 0 && x < len(heightmap)) && (y >= 0 && y < len(heightmap[0])) {
			if heightmap[x][y] < heightmap[origin.x][origin.y] {
				if explored[x][y] != basin {
					return false
				}
				foundBasin = true
			}
		}
	}

	return foundBasin
}

func appendAdjacent(to_explore []Coords, curr Coords, heightmap [][]int, explored [][]int) ([]Coords, int) {
	n_appended := 0
	for i := 0; i < 4; i++ {
		x := curr.x + x_offset[i]
		y := curr.y + y_offset[i]
		if (x >= 0 && x < len(heightmap)) && (y >= 0 && y < len(heightmap[0])) && explored[x][y] == 0 && heightmap[x][y] < 9 && flowsInto(heightmap, Coords{x, y}, explored[curr.x][curr.y], explored) {
			to_explore = append(to_explore, Coords{x, y})
			explored[x][y] = explored[curr.x][curr.y]
			n_appended++
		}
	}
	return to_explore, n_appended
}

func basinSize(heightmap [][]int) int {
	explored := make([][]int, len(heightmap))
	for i := range heightmap {
		explored[i] = make([]int, len(heightmap[0]))
	}
	basins := make([]int, 0)
	var n_appended int = 0

	for i, lowPoint := range getLowPoints(heightmap) {
		basin_size := 1
		to_explore := make([]Coords, 0)
		explored[lowPoint.x][lowPoint.y] = i + 1
		to_explore, n_appended = appendAdjacent(to_explore, lowPoint, heightmap, explored)
		basin_size += n_appended
		for len(to_explore) > 0 {
			curr := to_explore[0]
			to_explore = to_explore[1:]
			to_explore, n_appended = appendAdjacent(to_explore, curr, heightmap, explored)
			basin_size += n_appended
		}
		basins = append(basins, basin_size)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	return basins[0] * basins[1] * basins[2]

}

func SolveDay9() {
	heightmap := readHeightmap()
	fmt.Println("Part 1:", sumRiskLevel(heightmap))
	fmt.Println("Part 2:", basinSize(heightmap))
}
