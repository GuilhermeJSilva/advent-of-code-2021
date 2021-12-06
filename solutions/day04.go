package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	board       [5][5]int
	found       [][]bool
	row_missing []int
	col_missing []int
}

func readIntLine() []int {
	var line string
	pulledNumbers := []int{}
	fmt.Scanln(&line)
	strNumbers := strings.Split(line, ",")
	for _, strNumber := range strNumbers {
		value, _ := strconv.Atoi(strNumber)
		pulledNumbers = append(pulledNumbers, value)
	}
	return pulledNumbers
}

func readBoard() (Board, error) {
	board := Board{}

	fmt.Scanln()
	for i := 0; i < 5; i++ {
		board.row_missing = append(board.row_missing, 5)
		board.col_missing = append(board.col_missing, 5)
		board.found = append(board.found, make([]bool, 5))
		nRead, err := fmt.Scanf("%d %d %d %d %d", &board.board[i][0], &board.board[i][1], &board.board[i][2], &board.board[i][3], &board.board[i][4])
		if nRead != 5 {
			return board, err
		}
	}
	return board, nil
}

func readBoards() []Board {
	boards := []Board{}
	for {
		board, err := readBoard()
		if err == nil {
			boards = append(boards, board)
		} else {
			break
		}
	}
	return boards
}

func findInBoard(board *Board, number int) (bool, int, int) {
	for row, rowArr := range board.board {
		for col, value := range rowArr {
			if value == number {
				return true, row, col
			}
		}
	}
	return false, -1, -1

}

func updateBoard(board *Board, number int) bool {
	found, row, col := findInBoard(board, number)
	if found {
		board.found[row][col] = true
		board.col_missing[col]--
		board.row_missing[row]--
		if board.col_missing[col] == 0 || board.row_missing[row] == 0 {
			return true
		}
	}
	return false
}

func sumNotFound(board Board) int {
	sum := int(0)
	for row, rowArr := range board.board {
		for col, value := range rowArr {
			if !board.found[row][col] {
				sum += value
			}
		}
	}
	return sum
}

func findWinner(boards []Board, numbers []int) int {
	for _, number := range numbers {
		for i, board := range boards {
			end := updateBoard(&board, number)
			if end {
				fmt.Println(i, number, board)
				return sumNotFound(board) * number
			}
		}
	}

	return 0
}

func findLoser(boards []Board, numbers []int) int {
	winners := make([]bool, len(boards))
	last_winner := -1
	n_players := len(boards)
	for _, number := range numbers {
		for i, board := range boards {
			if !winners[i] {
				end := updateBoard(&board, number)
				if end {
					n_players--
					winners[i] = true
					last_winner = i
				}
			}

			if n_players == 0 {
				fmt.Println(i, number, board)
				return sumNotFound(boards[last_winner]) * number
			}
		}
	}
	return 0
}

func SolveDay4() {
	pulledNumbers := readIntLine()
	boards := readBoards()
	// fmt.Printf("Part 1: %v\n", findWinner(boards, pulledNumbers))
	fmt.Printf("Part 2: %v\n", findLoser(boards, pulledNumbers))

}
