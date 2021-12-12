package solutions

import (
	"fmt"
	"strings"
	"unicode"
)

type Edge struct {
	src string
	dst string
}

func readEdges() []Edge {
	edges := make([]Edge, 0)
	for {
		var str_edge string
		_, err := fmt.Scanln(&str_edge)
		if err != nil {
			break
		}
		strs := strings.Split(str_edge, "-")
		edges = append(edges, Edge{strs[0], strs[1]})
	}
	return edges
}

func buildGraph(edges []Edge) map[string][]string {
	graph := make(map[string][]string, 0)
	for _, edge := range edges {
		_, found := graph[edge.src]
		if !found {
			graph[edge.src] = make([]string, 0)

		}
		graph[edge.src] = append(graph[edge.src], edge.dst)
	}
	for _, edge := range edges {
		_, found := graph[edge.dst]
		if !found {
			graph[edge.dst] = make([]string, 0)

		}
		graph[edge.dst] = append(graph[edge.dst], edge.src)
	}
	return graph
}

func upperLoop(path []string, next string) bool {
	for i := len(path) - 1; i >= 0; i-- {
		stop := path[i]
		if unicode.IsLower(rune(stop[0])) {
			return false
		}
		if next == stop {
			return true
		}
	}
	return false
}

func lowerPresent(path []string, next string) bool {
	for _, p := range path {
		if p == next {
			return true
		}
	}
	return false
}

func isUpper(char byte) bool {
	return unicode.IsUpper(rune(char))
}

func isLower(char byte) bool {
	return unicode.IsLower(rune(char))
}

// DFS: search last big caves for next, search all small caves for next
func countPaths(graph map[string][]string, path []string) int {
	last := path[len(path)-1]
	if last == "end" {
		return 1
	}

	count := 0
	for _, next := range graph[last] {
		if (isUpper(next[0]) && !upperLoop(path, next)) || (isLower(next[0]) && !lowerPresent(path, next)) {
			count += countPaths(graph, append(path, next))
		}
	}
	return count
}

func lowerOneTwice(path []string, next string) bool {
	counter := make(map[string]int, 0)
	for _, stop := range path {
		if isLower(stop[0]) {
			counter[stop]++
			if counter[stop] > 1 {
				return lowerPresent(path, next)
			}
		}
	}
	return false

}

func countPathsTwice(graph map[string][]string, path []string) int {
	last := path[len(path)-1]
	if last == "end" {
		// fmt.Println(path)
		return 1
	}

	count := 0
	for _, next := range graph[last] {
		if next != "start" && ((isUpper(next[0]) && !upperLoop(path, next)) || (isLower(next[0]) && !lowerOneTwice(path, next))) {
			count += countPathsTwice(graph, append(path, next))
		}
	}
	return count
}

func SolveDay12() {
	edges := readEdges()
	graph := buildGraph(edges)
	fmt.Println(graph)
	fmt.Println("Part 1:", countPaths(graph, []string{"start"}))
	fmt.Println("Part 2:", countPathsTwice(graph, []string{"start"}))

}
