package solutions

import (
	"fmt"

	"github.com/GuilhermeJSilva/advent-of-code-2021/solutions/day24"
)

type ValueIndex struct {
	value, index int
}

func maximizeDigit(diff int) (int, int) {
	if diff > 0 {
		return 9, 9 - diff
	}
	return 9 + diff, 9
}
func minimizeDigit(diff int) (int, int) {
	if diff > 0 {
		return 1 + diff, 1
	}
	return 1, 1 - diff
}

func calcDigits(commands []day24.Command, digit_opt func(int) (int, int)) {
	monad := [14]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	stack := make([]ValueIndex, 0)

	for i := 0; i < 14; i++ {
		if commands[i*18+4].Other == 1 {
			value := commands[i*18+15].Other
			stack = append(stack, ValueIndex{value, i})
		} else {
			value := commands[i*18+5].Other
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			diff := top.value + value
			monad[i], monad[top.index] = digit_opt(diff)
		}
	}
	for _, v := range monad {
		fmt.Print(v)
	}
	fmt.Println()

}

func SolveDay24() {

	commands := day24.ReadALU()
	calcDigits(commands, maximizeDigit)
	calcDigits(commands, minimizeDigit)

}
