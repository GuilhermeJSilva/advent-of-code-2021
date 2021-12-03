package solutions

import (
	"fmt"
	"strconv"
)

func readDiagnosis() []string {
	var diagnosis []string = []string{}
	var bin_text string = ""
	for {
		n_read, _ := fmt.Scanf("%s", &bin_text)
		if n_read != 1 {
			break
		} else {
			diagnosis = append(diagnosis, bin_text)
		}
	}
	return diagnosis
}

func countOnes(diagnosis []string) []int {
	var length_binary int = len(diagnosis[0])
	var number_ones []int = make([]int, length_binary)

	for _, binary := range diagnosis {
		for index, value := range binary {
			if value == '1' {
				number_ones[index]++
			}
		}
	}
	return number_ones
}

func calculateGamma(diagnosis []string) uint {
	var number_ones []int = countOnes(diagnosis)

	var gamma uint = 0
	for _, n_ones := range number_ones {
		if n_ones > len(diagnosis)/2 {
			gamma = gamma | 1
		}
		gamma = gamma << 1
	}
	return gamma >> 1
}

func powerConsumption(diagnosis []string) uint {
	var length_binary int = len(diagnosis[0])
	var mask uint = ^((^uint(0)) << uint(length_binary))
	var gamma uint = calculateGamma(diagnosis)
	return gamma * ((^gamma) & mask)
}

func filterDiagnosis(diagnosis []string, pos int, target byte) []string {
	var filtered_diagnosis []string = []string{}
	for _, binary := range diagnosis {
		if binary[pos] == target {
			filtered_diagnosis = append(filtered_diagnosis, binary)
		}
	}
	return filtered_diagnosis
}

func lastBinaryStanding(diagnosis []string, max bool) uint {
	var length_binary int = len(diagnosis[0])
	for i := 0; i < length_binary && len(diagnosis) > 1; i++ {
		var number_ones []int = countOnes(diagnosis)
		if (max && float64(number_ones[i]) >= float64(len(diagnosis))/2.0) || (!max && float64(number_ones[i]) < float64(len(diagnosis))/2) {
			diagnosis = filterDiagnosis(diagnosis, i, '1')
		} else {
			diagnosis = filterDiagnosis(diagnosis, i, '0')
		}
	}
	oxygen, _ := strconv.ParseUint(diagnosis[0], 2, 64)
	return uint(oxygen)
}

func calculateLifeSupport(diagnosis []string) uint {
	var oxygen_rating uint = lastBinaryStanding(diagnosis, true)
	var co2_rating uint = lastBinaryStanding(diagnosis, false)
	fmt.Printf("%v %v\n", oxygen_rating, co2_rating)
	return oxygen_rating * co2_rating
}

func SolveDay3() {
	var diagnosis []string = readDiagnosis()
	fmt.Printf("Part 1: %v\n", powerConsumption(diagnosis))
	fmt.Printf("Part 2: %v\n", calculateLifeSupport(diagnosis))
}
