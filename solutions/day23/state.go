package day23

import (
	"fmt"
	"math"
)

type AmphiBoard [][]int
type AmphiPosition struct {
	x int
	y int
}
type AmphiState struct {
	board        AmphiBoard
	spent_energy int
	amphipos     [4][]AmphiPosition
}

func CreateState(board AmphiBoard, energy int) AmphiState {
	state := AmphiState{}
	state.board = board
	state.spent_energy = energy

	for i, line := range board {
		for j, value := range line {
			if value == 0 {
				continue
			}
			state.amphipos[value-1] = append(state.amphipos[value-1], AmphiPosition{i, j})
		}
	}
	return state
}

func (board AmphiBoard) copy() AmphiBoard {
	new_board := make(AmphiBoard, len(board))
	for i := range new_board {
		new_board[i] = make([]int, len(board[i]))
		copy(new_board[i], board[i])
	}
	return new_board
}

func copyPositions(positions [4][]AmphiPosition) [4][]AmphiPosition {
	new_positions := [4][]AmphiPosition{}
	for i := range positions {
		new_positions[i] = make([]AmphiPosition, len(positions[i]))
		copy(new_positions[i], positions[i])
	}

	return new_positions
}

func (state AmphiState) Print() {
	fmt.Println("Energy: ", state.spent_energy)
	fmt.Println("Positions: ", state.amphipos)
	state.board.Print()
}

func (board AmphiBoard) Print() {
	for _, line := range board {
		for _, value := range line {
			if value == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string('A' + value - 1))
			}
		}
		fmt.Println()
	}
}

func energy(start AmphiPosition, end AmphiPosition, t int) int {
	distance := 0
	if start.x > end.x {
		distance += start.x - end.x
	} else {
		distance += end.x - start.x
	}
	if start.y > end.y {
		distance += start.y - end.y
	} else {
		distance += end.y - start.y
	}
	return distance * int(math.Pow10(t-1))
}

func (state AmphiState) isFinal() bool {
	for i, by_type := range state.amphipos {
		for _, position := range by_type {
			if position.x == 0 || position.y != 2*(i+1) {
				return false
			}
		}
	}
	return true
}

func (state AmphiState) move(start AmphiPosition, end AmphiPosition) *AmphiState {
	amphi_t := state.board[start.x][start.y]
	new_state := &AmphiState{}
	new_state.board = state.board.copy()
	new_state.amphipos = copyPositions(state.amphipos)
	new_state.spent_energy = state.spent_energy + energy(start, end, amphi_t)
	for pos_idx, pos := range state.amphipos[amphi_t-1] {
		if pos == start {
			new_state.amphipos[amphi_t-1][pos_idx] = end
		}
	}
	new_state.board[start.x][start.y] = 0
	new_state.board[end.x][end.y] = amphi_t
	return new_state
}

func (state AmphiState) hallwayMoves(start AmphiPosition) []*AmphiState {
	next_states := make([]*AmphiState, 0)

	for i := start.y + 1; i < 11; i++ {
		if state.board[0][i] != 0 {
			break
		}
		if i > 8 || i%2 == 1 {
			next_states = append(next_states, state.move(start, AmphiPosition{0, i}))
		}
	}
	for i := start.y - 1; i >= 0; i-- {
		if state.board[0][i] != 0 {
			break
		}
		if i < 2 || i%2 == 1 {
			next_states = append(next_states, state.move(start, AmphiPosition{0, i}))
		}
	}

	return next_states
}

func (state AmphiState) isNonIncorrect(amphi_t int) bool {
	for _, line := range state.board[1:] {
		if line[2*amphi_t] != 0 && line[2*amphi_t] != amphi_t {
			return false
		}
	}
	return true
}

func (state AmphiState) countCorrect(amphi_t int) int {
	correct := 0
	for i := len(state.board) - 1; i > 0; i-- {
		if state.board[i][2*amphi_t] != amphi_t {
			break
		}
		correct++
	}
	return correct
}

func (state AmphiState) isFrontClear(position AmphiPosition) bool {
	for i := 1; i < position.x; i++ {
		if state.board[i][position.y] != 0 {
			return false
		}
	}

	return true
}

func (state AmphiState) isHorizontalClear(start, end int) bool {
	step := 1
	if start > end {
		step = -1
	}
	for y := start; y != end; y += step {
		if state.board[0][y] != 0 {
			return false
		}
	}
	return true

}

func (state AmphiState) isVerticalClear(start, end, y int) bool {
	step := 1
	if start > end {
		step = -1
	}
	for x := start; x != end; x += step {
		if state.board[x][y] != 0 {
			return false
		}
	}
	return true

}

func (state AmphiState) isPathClear(start, end AmphiPosition) bool {
	step_x, step_y := 1, 1
	if start.y > end.y {
		step_y = -1
	}
	if start.x > end.x {
		step_x = -1
	}
	if start.x == 0 {
		return state.isHorizontalClear(start.y+step_y, end.y) && state.isVerticalClear(start.x, end.x, end.y)
	}
	return state.isVerticalClear(start.x+step_x, end.x, start.y) && state.isHorizontalClear(start.y, end.y)

}

func (state AmphiState) allMovesSingle(start AmphiPosition) []*AmphiState {
	amphi_t := state.board[start.x][start.y]
	last_room := len(state.board) - 1
	target_room := 2 * amphi_t
	correct := state.countCorrect(amphi_t)
	last_correct := last_room - correct + 1
	if start.x == 0 { // In hallway
		if state.isNonIncorrect(amphi_t) && state.isPathClear(start, AmphiPosition{last_correct - 1, target_room}) { // Can move into room
			return []*AmphiState{state.move(start, AmphiPosition{last_correct - 1, target_room})}
		}
	} else { // In room
		if start.y == target_room && state.isNonIncorrect(amphi_t) { // Correct Room
			if start.x >= last_correct { // In the back
				return []*AmphiState{}
			} else if state.isNonIncorrect(amphi_t) { // In the front with back empty
				return []*AmphiState{state.move(start, AmphiPosition{last_correct - 1, start.y})}
			}
		} else { // Wrong room
			if state.isFrontClear(start) { // Nothing in front
				return state.hallwayMoves(start)
			}
		}
	}
	return []*AmphiState{}
}

func (state AmphiState) AllMoves() []*AmphiState {
	moves := make([]*AmphiState, 0)

	for _, by_type := range state.amphipos {
		for _, position := range by_type {
			moves = append(moves, state.allMovesSingle(position)...)
		}
	}
	return moves
}
