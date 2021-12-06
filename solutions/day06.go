package solutions

import (
	"fmt"
)

func reproduceFish(state []int, days int) int {
	fish_counts := make([]int, 9)
	for _, s := range state {
		fish_counts[s]++
	}
	total := len(state)
	for i := 0; i < days; i++ {
		to_add := fish_counts[0]
		for j := 1; j < 9; j++ {
			fish_counts[j-1] = fish_counts[j]
		}
		fish_counts[6] += to_add
		fish_counts[8] = to_add
		total += to_add

	}
	return total
}

func SolveDay6() {
	state := readIntLine()
	fmt.Println("Part 1: ", reproduceFish(state, 80))
	fmt.Println("Part 2: ", reproduceFish(state, 256))
}
