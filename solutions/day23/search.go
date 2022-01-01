package day23

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

type AmphiStateKey string

func (state AmphiState) toKey() AmphiStateKey {
	lines := make([]string, 0)
	for _, board_line := range state.board {
		lines = append(lines, fmt.Sprint(board_line))
	}

	return AmphiStateKey(strings.Join(lines, "") + strconv.Itoa(state.spent_energy))
}

type AmphiStateIndex struct {
	State *AmphiState
	index int
	Prev  *AmphiStateIndex
}

type AmphiPQ []*AmphiStateIndex

func (pq AmphiPQ) Len() int { return len(pq) }

func (pq AmphiPQ) Less(i, j int) bool {
	return pq[i].State.spent_energy < pq[j].State.spent_energy
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

func LowestEnergy(start AmphiState) *AmphiStateIndex {
	pq := make(AmphiPQ, 0)
	states := map[AmphiStateKey]*AmphiState{}
	states[start.toKey()] = &start
	heap.Push(&pq, &AmphiStateIndex{State: &start, Prev: nil})
	i := 0
	for pq.Len() > 0 {
		i++
		item := heap.Pop(&pq).(*AmphiStateIndex)
		if item.State.isFinal() {
			return item
		}

		for _, new_state := range item.State.AllMoves() {
			new_item := AmphiStateIndex{
				State: new_state,
				Prev:  item,
			}
			new_key := new_state.toKey()
			if _, ok := states[new_key]; !ok {
				heap.Push(&pq, &new_item)
				states[new_key] = new_state
			}
		}
	}
	return nil
}
