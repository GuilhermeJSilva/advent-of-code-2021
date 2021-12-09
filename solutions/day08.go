package solutions

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Entry struct {
	signals []string
	output  []string
}

func readEntries() []Entry {
	entries := make([]Entry, 0)
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		line := reader.Text()
		parts := strings.Split(line, "|")
		entry := Entry{strings.Fields(parts[0]), strings.Fields(parts[1])}
		entries = append(entries, entry)
	}

	return entries
}

func findUniques(entries []Entry) int {
	uniques := 0

	for _, entry := range entries {
		for _, output := range entry.output {
			size := len(output)
			if size == 2 || size == 4 || size == 3 || size == 7 {
				uniques++
			}
		}
	}
	return uniques
}

func contains(lhs string, rhs string) int {
	matches := len(rhs)
	for _, character := range rhs {
		if !strings.Contains(lhs, string(character)) {
			matches--
		}
	}
	return matches
}

func decodeSignals(signals []string) []func(string) bool {
	sort.Slice(signals, func(i, j int) bool {
		return len(signals[i]) < len(signals[j])
	})
	decoder := make([]func(string) bool, 10)
	decoder[0] = func(s string) bool {
		return len(s) == 6 && contains(s, signals[0]) == 2 && contains(s, signals[2]) == 3
	}
	decoder[1] = func(s string) bool { return len(s) == 2 }
	decoder[2] = func(s string) bool { return len(s) == 5 && contains(s, signals[2]) == 2 }
	decoder[3] = func(s string) bool { return len(s) == 5 && contains(s, signals[0]) == 2 }
	decoder[4] = func(s string) bool { return len(s) == 4 }
	decoder[5] = func(s string) bool { return len(s) == 5 }
	decoder[6] = func(s string) bool { return len(s) == 6 && contains(s, signals[0]) == 1 }
	decoder[7] = func(s string) bool { return len(s) == 3 }
	decoder[8] = func(s string) bool { return len(s) == 7 }
	decoder[9] = func(s string) bool { return len(s) == 6 }

	return decoder
}

func decodeOutput(entry Entry) int {
	value := 0
	decoders := decodeSignals(entry.signals)
	for _, output := range entry.output {
		for i, decoder := range decoders {
			if decoder(output) {
				value += i
				break
			}
		}
		value *= 10
	}
	return value / 10
}

func sumOutputs(entries []Entry) int {
	sum := 0
	for _, entry := range entries {
		sum += decodeOutput(entry)
	}
	return sum
}

func SolveDay8() {
	entries := readEntries()
	fmt.Println("Part 1: ", findUniques(entries))
	fmt.Println("Part 2: ", sumOutputs(entries))
}
