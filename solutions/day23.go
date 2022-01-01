package solutions

import (
	"fmt"

	"github.com/GuilhermeJSilva/advent-of-code-2021/solutions/day23"
)

func readStartState() [][]int {
	state := make([][]int, 0)
	var line string
	for {
		_, err := fmt.Scanln(&line)
		if err != nil {
			break
		}
		state = append(state, make([]int, 11))
		next := 2
		for _, val := range line {
			if val >= 'A' && val <= 'D' {
				state[len(state)-1][next] = int(val - 'A' + 1)
				next += 2
			}
		}
	}
	return state[1 : len(state)-1]
}

func SolveDay23() {

	board := readStartState()

	amphiState := day23.CreateState(board, 0)
	final := day23.LowestEnergy(amphiState)
	if final != nil {
		final.State.Print()
	}
}
