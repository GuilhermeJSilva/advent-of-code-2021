package solutions

import "fmt"

type Entry struct {
	signals []string
	output  []string
}

func readEntries() []Entry {
	entries := make([]Entry, 0)
	for {
		var line string
		_, err := fmt.Scanln(&line)
		if err == nil {
			break
		}

	}

	return entries
}

func SolveDay8() {
	entries := readEntries()
	fmt.Println(entries[0])
}
