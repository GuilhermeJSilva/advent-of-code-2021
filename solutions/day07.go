package solutions

import (
	"fmt"
	"math"
	"sort"
)

func median(values []int) int {
	sort.Ints(values)
	middle := len(values) / 2
	if len(values)%2 == 1 {
		return values[middle]
	}
	return (values[middle-1] + values[middle]) / 2

}

func intAverage(values []int) (int, int) {
	var sum float64 = 0
	for _, value := range values {
		sum += float64(value)
	}
	avg := sum / float64(len(values))

	return int(math.Floor(avg)), int(math.Ceil(avg))
}

func distanceTo(values []int, target int) int {
	distance := 0
	for _, value := range values {
		if value > target {
			distance += value - target
		} else {
			distance += target - value
		}
	}
	return distance
}

func increasingDistance(value int) int {
	return (value * (value + 1)) / 2
}

func factorialDistanceTo(values []int, target int) int {
	distance := 0
	for _, value := range values {
		if value > target {
			distance += increasingDistance(value - target)
		} else {
			distance += increasingDistance(target - value)
		}
	}
	return distance

}

func SolveDay7() {
	distances := readIntLine()
	fmt.Println("Part 1:", distanceTo(distances, median(distances)))
	min_avg, max_avg := intAverage(distances)
	fmt.Println("Part 2:", factorialDistanceTo(distances, min_avg))
	fmt.Println("Part 2:", factorialDistanceTo(distances, max_avg))

}
