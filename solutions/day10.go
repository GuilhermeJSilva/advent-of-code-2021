package solutions

import "fmt"

func readStrings() []string {
	lines := make([]string, 0)
	for {
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil {
			break
		}
		lines = append(lines, line)
	}
	return lines
}

func firstError(str string) int {
	stack := make([]byte, 0)

	for _, char := range str {
		if char == '}' {
			if stack[len(stack)-1] != '{' {
				return 1197
			}
			stack = stack[:len(stack)-1]
		} else if char == '>' {
			if stack[len(stack)-1] != '<' {
				return 25137
			}
			stack = stack[:len(stack)-1]

		} else if char == ']' {
			if stack[len(stack)-1] != '[' {
				return 57
			}
			stack = stack[:len(stack)-1]

		} else if char == ')' {
			if stack[len(stack)-1] != '(' {
				return 3
			}
			stack = stack[:len(stack)-1]

		} else {
			stack = append(stack, byte(char))
		}
	}

	return 0
}

func sumFirstErrors(strs []string) int {
	sum := 0
	for _, str := range strs {
		err := firstError(str)
		sum += err
	}
	return sum
}

var matches [4]string = [4]string{"()", "[]", "{}", "<>"}

func sequenceScore(seq []byte) int {
	score := 0

	for i := len(seq) - 1; i >= 0; i-- {
		score *= 5
		for j, match := range matches {
			if seq[i] == match[0] {
				score += j + 1
			}
		}
	}

	return score
}

func completionScore(str string) int {
	stack := make([]byte, 0)

	for _, char := range str {
		matched := false
		if len(stack) != 0 {
			for _, match := range matches {
				if byte(char) == match[1] {
					if stack[len(stack)-1] != match[0] {
						return 0
					}
					stack = stack[:len(stack)-1]
					matched = true
					break
				}
			}
		}

		if !matched {
			stack = append(stack, byte(char))
		}
	}

	return sequenceScore(stack)
}

func medianCompletionScore(strs []string) int {
	scores := make([]int, 0)

	for _, str := range strs {
		score := completionScore(str)
		if score != 0 {
			scores = append(scores, score)
		}
	}
	return median(scores)
}

func SolveDay10() {
	lines := readStrings()
	fmt.Println("Part 1:", sumFirstErrors(lines))
	fmt.Println("Part 2:", medianCompletionScore(lines))
}
