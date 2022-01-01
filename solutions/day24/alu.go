package day24

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Registers [4]int
type Inputs []int

type ALU struct {
	Reg         Registers
	inp         Inputs
	inp_pointer int
}

func CreateAlu(inputs Inputs) ALU {
	alu := ALU{}
	alu.inp = inputs
	return alu
}

type Command struct {
	Name        string
	Target      int
	Other       int
	Is_register bool
}

func (alu *ALU) execute(cmd Command) bool {
	other_val := cmd.Other
	if cmd.Is_register {
		other_val = alu.Reg[cmd.Other]
	}
	switch cmd.Name {
	case "inp":
		alu.Reg[cmd.Target] = alu.inp[alu.inp_pointer]
		alu.inp_pointer++

	case "add":
		alu.Reg[cmd.Target] += other_val

	case "mul":
		alu.Reg[cmd.Target] *= other_val

	case "div":
		if other_val == 0 {
			return false
		}
		alu.Reg[cmd.Target] /= other_val

	case "mod":
		if alu.Reg[cmd.Target] < 0 || other_val <= 0 {
			return false
		}
		alu.Reg[cmd.Target] %= other_val

	case "eql":
		if alu.Reg[cmd.Target] == other_val {
			alu.Reg[cmd.Target] = 1
		} else {
			alu.Reg[cmd.Target] = 0
		}
	}

	return true
}

func (alu *ALU) Program(cmds []Command) bool {
	for _, cmd := range cmds {
		if !alu.execute(cmd) {
			return false
		}
	}
	return true
}

func ReadALU() []Command {
	commands := make([]Command, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		if len(parts) == 2 {
			commands = append(commands, Command{parts[0], int(parts[1][0] - 'w'), 0, false})
		} else if len(parts) == 3 {
			value, err := strconv.ParseInt(parts[2], 10, 64)
			if err == nil {
				commands = append(commands, Command{parts[0], int(parts[1][0] - 'w'), int(value), false})
			} else {
				commands = append(commands, Command{parts[0], int(parts[1][0] - 'w'), int(parts[2][0] - 'w'), true})
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return commands
}
