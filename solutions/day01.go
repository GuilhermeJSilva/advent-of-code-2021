package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readIntList() []int {
	intList := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parsedInt, error := strconv.ParseInt(scanner.Text(), 10, 64)
		if error == nil {
			intList = append(intList, int(parsedInt))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return intList
}

func countIncreases(values []int) int {
	var previous int = int((^uint(0)) >> 1)
	var increases int = 0
	for _, num := range values {
		if previous < num {
			increases = increases + 1
		}
		previous = num
	}
	return increases
}

func countTrioIncreases(values []int) int {
	var previous int = int((^uint(0)) >> 1)
	var increases int = 0
	for i := 2; i < len(values); i++ {
		sum := values[i] + values[i-1] + values[i-2]
		if previous < sum {
			increases = increases + 1
		}
		previous = sum

	}
	return increases
}

func SolveDay1() {
	var intList []int = readIntList()
	fmt.Printf("Part 1: %v\n", countIncreases(intList))
	fmt.Printf("Part 2: %v\n", countTrioIncreases(intList))
}
