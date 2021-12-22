package solutions

import "fmt"

func readPositions() []int {
	positions := make([]int, 0)
	for {
		var player, pos int
		_, err := fmt.Scanf("Player %d starting position: %d\n", &player, &pos)

		if err != nil {
			break
		}
		positions = append(positions, pos-1)
	}
	return positions
}

type DDie struct {
	next_value int
	max        int
	dies_cast  int
}

func (d *DDie) next() int {
	prev := d.next_value
	d.next_value = (d.next_value % d.max) + 1
	d.dies_cast++
	return prev
}

func playDeterministic(positions []int) int {
	scores := make([]int, len(positions))
	fmt.Println(positions)
	die := DDie{1, 100, 0}

	for i := 0; ; i++ {
		player := 0
		if i%2 == 1 {
			player = 1
		}

		positions[player] += die.next() + die.next() + die.next()
		positions[player] %= 10
		scores[player] += positions[player] + 1

		if scores[player] >= 1000 {
			other := 1 - player
			fmt.Println(player, other, scores[other], die.dies_cast)
			return scores[other] * die.dies_cast
		}
	}
}

type Results struct {
	wins     uint
	non_wins uint
}

type DiceResult struct {
	value     int
	frequency uint
}

var dice_res [7]DiceResult = [7]DiceResult{{3, 1}, {4, 3}, {5, 6}, {6, 7}, {7, 6}, {8, 3}, {9, 1}}

func calcResults(position int, remaining int, ways uint, steps int, results []Results) {
	if steps < len(results) {
		for _, dice := range dice_res {
			new_position := (position + dice.value) % 10
			new_remaining := remaining - (new_position + 1)
			new_ways := ways * dice.frequency
			if new_remaining <= 0 {
				results[steps].wins += new_ways
			} else {
				results[steps].non_wins += new_ways
				calcResults(new_position, new_remaining, new_ways, steps+1, results)
			}
		}
	}
}

func getResults(position int, victory int) []Results {
	results := make([]Results, victory)
	calcResults(position, victory, 1, 0, results)
	return results
}

func playDedic(positions []int) uint {
	victory := 21
	player1 := getResults(positions[0], victory)
	player2 := getResults(positions[1], victory)

	p1_total := uint(0)
	p2_total := uint(0)
	for i := 1; i < len(player1); i++ {
		p1_total += player1[i].wins * player2[i-1].non_wins
		p2_total += player2[i].wins * player1[i].non_wins
	}

	if p1_total > p2_total {
		return p1_total
	}
	return p2_total
}

func SolveDay21() {
	postions := readPositions()
	p2_positions := make([]int, 2)
	copy(p2_positions, postions)
	fmt.Println(playDeterministic(postions))
	fmt.Println(p2_positions)
	fmt.Println(playDedic(p2_positions))
}
