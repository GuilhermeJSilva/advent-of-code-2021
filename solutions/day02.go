package solutions

import (
	"fmt"
)

type Movement struct {
	direction string
	units     int
}

func readMovementList() []Movement {
	movementList := []Movement{}
	var units int = 0
	var direction string = ""
	for {
		n_read, _ := fmt.Scanf("%s %d", &direction, &units)
		if n_read != 2 {
			break
		} else {
			movementList = append(movementList, Movement{direction, units})
		}
	}
	return movementList
}

func finalPositions(movements []Movement) int {
	var vertical int = 0
	var horizontal int = 0
	for _, mov := range movements {
		if mov.direction == "up" {
			vertical -= mov.units
		} else if mov.direction == "down" {
			vertical += mov.units
		} else {
			horizontal += mov.units
		}
	}
	fmt.Printf("%v %v\n", vertical, horizontal)
	return vertical * horizontal
}

func finalPositionsWithAim(movements []Movement) int {
	var vertical int = 0
	var horizontal int = 0
	var aim int = 0
	for _, mov := range movements {
		if mov.direction == "up" {
			aim -= mov.units
		} else if mov.direction == "down" {
			aim += mov.units
		} else {
			horizontal += mov.units
			vertical += mov.units * aim
		}
	}
	fmt.Printf("%v %v\n", vertical, horizontal)
	return vertical * horizontal
}

func SolveDay2() {
	var movementList []Movement = readMovementList()
	fmt.Printf("Part 1: %v\n", finalPositions(movementList))
	fmt.Printf("Part 2: %v\n", finalPositionsWithAim(movementList))
}
