package solutions

import (
	"container/heap"
	"fmt"
	"math"
)

type AmphiBoard [3][11]int
type AmphiPosition struct {
	x int
	y int
}
type AmphiState struct {
	board        AmphiBoard
	spent_energy int
	amphipos     [4][2]AmphiPosition
}

func createState(board AmphiBoard, energy int) AmphiState {
	next := [4]int{0, 0, 0, 0}
	state := AmphiState{}
	state.board = board
	state.spent_energy = energy

	for i, line := range board {
		for j, value := range line {
			if value == 0 {
				continue
			}
			state.amphipos[value-1][next[value-1]] = AmphiPosition{i, j}
			next[value-1]++
		}
	}
	return state

}
func (state AmphiState) print() {
	fmt.Println("Energy: ", state.spent_energy)
	fmt.Println("Positions: ", state.amphipos)
	state.board.print()
}

func (board AmphiBoard) print() {
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

func (state AmphiState) move(start AmphiPosition, end AmphiPosition) AmphiState {
	amphi_t := state.board[start.x][start.y]
	new_state := state
	new_state.spent_energy += energy(start, end, amphi_t)
	if new_state.amphipos[amphi_t-1][0] == start {
		new_state.amphipos[amphi_t-1][0] = end
	} else {
		new_state.amphipos[amphi_t-1][1] = end
	}
	new_state.board[start.x][start.y] = 0
	new_state.board[end.x][end.y] = amphi_t
	return new_state
}

func (state AmphiState) hallwayMoves(start AmphiPosition) []AmphiState {
	next_states := make([]AmphiState, 0)

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

func (state AmphiState) allMovesSingle(start AmphiPosition) []AmphiState {
	amphi_t := state.board[start.x][start.y]
	if start.x == 0 { // In hallway
		if state.board[1][2*amphi_t] == 0 { // Front of target room is empty
			if state.board[2][2*amphi_t] == amphi_t { // The pair is in the back
				return []AmphiState{state.move(start, AmphiPosition{1, 2 * amphi_t})}
			} else if state.board[2][2*amphi_t] == 0 { // The room is empty
				return []AmphiState{state.move(start, AmphiPosition{2, 2 * amphi_t})}
			} else {
				return []AmphiState{}
				// return state.hallwayMoves(start)
			}
		} else {
			return []AmphiState{}
			// return state.hallwayMoves(start)
		}
	} else { // In room
		if start.y == 2*amphi_t { // Correct Room
			if start.x == 2 { // In the back
				return []AmphiState{}
			} else if state.board[2][start.y] == 0 { // In the front with back empty
				return []AmphiState{state.move(start, AmphiPosition{2, start.y})}
			} else if state.board[2][start.y] == amphi_t { // In the front with correct back
				return []AmphiState{}
			} else {
				return state.hallwayMoves(start)
			}
		} else { // Wrong room
			if start.x == 1 { // Front of the room
				return state.hallwayMoves(start)
			} else if state.board[1][start.y] == 0 { // Nothing in front
				return state.hallwayMoves(start)
			} else { // Blocked in
				return []AmphiState{}
			}
		}
	}
}

func (state AmphiState) allMoves() []AmphiState {
	moves := make([]AmphiState, 0)

	for _, by_type := range state.amphipos {
		for _, position := range by_type {
			moves = append(moves, state.allMovesSingle(position)...)
		}
	}
	return moves
}

type AmphiStateIndex struct {
	state AmphiState
	index int
}

type AmphiPQ []*AmphiStateIndex

func (pq AmphiPQ) Len() int { return len(pq) }

func (pq AmphiPQ) Less(i, j int) bool {
	return pq[i].state.spent_energy < pq[j].state.spent_energy
}

func (pq AmphiPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *AmphiPQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*AmphiStateIndex)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *AmphiPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
func lowestEnergy(start AmphiState) int {
	pq := make(AmphiPQ, 0)
	states := map[AmphiState]int{}
	states[start] = 1
	heap.Push(&pq, &AmphiStateIndex{state: start})
	i := 0
	for pq.Len() > 0 {
		i++
		item := heap.Pop(&pq).(*AmphiStateIndex)
		if item.state.isFinal() {
			item.state.print()
			return item.state.spent_energy
		}

		for _, new_state := range item.state.allMoves() {
			new_item := AmphiStateIndex{
				state: new_state,
			}
			if _, ok := states[new_state]; !ok {
				heap.Push(&pq, &new_item)
				states[new_state] = 1
			}
		}
	}
	return -1
}

func SolveDay23() {

	// amphiBoard := AmphiBoard{
	// 	[11]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	[11]int{0, 0, 2, 0, 3, 0, 2, 0, 4, 0, 0},
	// 	[11]int{0, 0, 1, 0, 4, 0, 3, 0, 1, 0, 0},
	// }
	amphiBoard := AmphiBoard{
		[11]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[11]int{0, 0, 4, 0, 2, 0, 4, 0, 2, 0, 0},
		[11]int{0, 0, 3, 0, 1, 0, 1, 0, 3, 0, 0},
	}

	amphiState := createState(amphiBoard, 0)
	fmt.Println(lowestEnergy(amphiState) + 2) // There is a but somewhere
}
