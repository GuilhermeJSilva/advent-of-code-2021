package solutions

import (
	"fmt"
	"math"
)

func readInserionRules() map[string]string {
	rules := make(map[string]string, 0)
	for {
		var target, result string
		_, err := fmt.Scanf("%s -> %s", &target, &result)
		if err != nil {
			break
		}
		rules[target] = result
	}

	return rules
}

func insertionStep(start map[string]int, counter map[string]int, rules map[string]string) map[string]int {
	new := make(map[string]int)
	for k := range start {
		converted, found := rules[k]
		if found {
			counter[converted] += start[k]
			created_1 := string(k[0]) + converted
			new[created_1] += start[k]
			created_2 := converted + string(k[1])
			new[created_2] += start[k]
		}
	}
	return new

}

func insertionSteps(start string, rules map[string]string, steps int) map[string]int {
	char_counter := buildCharCounter(start)
	pair_counter := buildPairCounter(start)
	for i := 0; i < steps; i++ {
		pair_counter = insertionStep(pair_counter, char_counter, rules)
	}
	return char_counter
}

func getMaxMin(counter map[string]int) (int, int) {
	max := math.MinInt
	min := math.MaxInt
	for _, value := range counter {
		if value > max {
			max = value
		} else if value < min {
			min = value
		}
	}
	return max, min
}

func maxMinAfter(start string, rules map[string]string, steps int) int {
	counter := insertionSteps(start, rules, steps)
	max, min := getMaxMin(counter)
	return max - min
}

func buildPairCounter(str string) map[string]int {
	counter := make(map[string]int)
	for i := 0; i < len(str)-1; i++ {
		pair := string(str[i]) + string(str[i+1])
		counter[pair]++

	}
	return counter
}

func buildCharCounter(str string) map[string]int {
	char_counter := make(map[string]int)
	for _, char := range str {
		char_counter[string(char)]++
	}
	return char_counter
}

func SolveDay14() {
	var sequence, tmp string
	fmt.Scanln(&sequence)
	fmt.Scanln(&tmp)
	rules := readInserionRules()
	fmt.Println(sequence)
	fmt.Println(rules)
	fmt.Println("Part 1: ", maxMinAfter(sequence, rules, 10))
	fmt.Println("Part 2: ", maxMinAfter(sequence, rules, 40))
}
