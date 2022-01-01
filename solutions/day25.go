package solutions

import (
	"fmt"
)

type State [][]byte

func readInitialState() State {
	state := make([][]byte, 0)
	var line string
	for {
		_, err := fmt.Scanln(&line)
		if err != nil {
			break
		}
		state = append(state, []byte(line))

	}
	return state
}

func (state *State) step() bool {
	moved := false
	size_col := len(*state)
	size_line := len((*state)[0])

	for i := 0; i < size_col; i++ {
		for j := 0; j < size_line; j++ {
			if (*state)[i][j] == '>' && (*state)[i][(j+1)%size_line] == '.' {
				(*state)[i][(j+1)%size_line] = '>'
				(*state)[i][j] = 'X'
				j++
				moved = true
			}
		}
	}
	for i := 0; i < size_col; i++ {
		for j := 0; j < size_line; j++ {
			if (*state)[i][j] == 'X' {
				(*state)[i][j] = '.'

			}
		}
	}

	for j := 0; j < size_line; j++ {
		for i := 0; i < size_col; i++ {
			if (*state)[i][j] == 'v' && (*state)[(i+1)%size_col][j] == '.' {
				(*state)[(i+1)%size_col][j] = 'v'
				(*state)[i][j] = 'X'
				i++
				moved = true
			}
		}
	}
	for i := 0; i < size_col; i++ {
		for j := 0; j < size_line; j++ {
			if (*state)[i][j] == 'X' {
				(*state)[i][j] = '.'

			}
		}
	}

	return moved
}

func (state *State) untilNoMotion() int {
	steps := 0
	for state.step() {
		steps++
		// fmt.Println("Step", steps)
		// for _, line := range *state {
		// 	fmt.Println(string(line))
		// }
		// fmt.Println("*****")
	}
	return steps + 1
}

func SolveDay25() {
	state := readInitialState()
	fmt.Println(state.untilNoMotion())
}
