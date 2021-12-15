package solutions

import (
	"container/heap"
	"fmt"
)

func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println()
}

type PQItem struct {
	coords   Coords
	distance int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*PQItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PQItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func lowestPath(graph [][]int) int {
	pq := make(PriorityQueue, 0)
	start := &PQItem{coords: Coords{0, 0}, distance: 0}
	graph[0][0] *= -1
	heap.Push(&pq, start)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*PQItem)
		if item.coords.x == len(graph)-1 && item.coords.y == len(graph[0])-1 {
			return item.distance
		}
		for i := 0; i < 4; i++ {
			x := item.coords.x + x_offset[i]
			y := item.coords.y + y_offset[i]
			if x >= 0 && x < len(graph) && y >= 0 && y < len(graph[0]) && graph[x][y] > 0 {
				new_item := PQItem{
					coords:   Coords{x, y},
					distance: item.distance + graph[x][y],
				}

				graph[x][y] *= -1
				heap.Push(&pq, &new_item)
			}
		}
	}
	return -1
}

func multiplyBoard(graph [][]int) [][]int {
	len_x := len(graph)
	len_y := len(graph[0])

	new_graph := make([][]int, len_x*5)
	for row_i := range new_graph {
		new_graph[row_i] = make([]int, len_y*5)
		for col_i := range new_graph[row_i] {
			to_add := row_i/len_x + col_i/len_y
			new_graph[row_i][col_i] = (graph[row_i%len_x][col_i%len_y] + to_add)
			if new_graph[row_i][col_i] > 9 {
				new_graph[row_i][col_i] -= 9
			}
		}
	}
	return new_graph
}

func SolveDay15() {
	riskmap := readHeightmap()
	fmt.Println(len(riskmap), len(riskmap[0]))
	fmt.Println("Part 1:", lowestPath(copyMatrix(riskmap)))
	fmt.Println("Part 2:", lowestPath(multiplyBoard(riskmap)))
}
