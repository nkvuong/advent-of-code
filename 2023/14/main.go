package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func moveNorth(board [][]byte) {
	for i := 1; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			for k := i - 1; k >= 0; k-- {
				if board[k][j] == '.' && board[k+1][j] == 'O' {
					board[k][j] = 'O'
					board[k+1][j] = '.'
				}
			}
		}
	}
}

func calc(board [][]byte) (sum int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				sum += len(board) - i
			}
		}
	}
	return sum
}

func rotate(board [][]byte) [][]byte {
	rows, cols := len(board), len(board[0])

	rotated := make([][]byte, cols)
	for i := 0; i < cols; i++ {
		rotated[i] = make([]byte, rows)
		for j := 0; j < rows; j++ {
			rotated[i][j] = board[rows-1-j][i]
		}
	}
	return rotated
}

func key(board [][]byte) (out string) {
	for _, line := range board {
		out = out + string(line)
	}
	return out
}

func deepCopy(board [][]byte) (out [][]byte) {
	out = make([][]byte, len(board))
	for i, row := range board {
		out[i] = append([]byte(nil), row...)
	}
	return out
}

func main() {

	lines := strings.Split(input, "\n")
	board := make([][]byte, 0)
	for _, line := range lines {
		board = append(board, []byte(line))
	}

	// part 1
	copyBoard := deepCopy(board)
	moveNorth(copyBoard)
	fmt.Println("Total load is", calc(copyBoard))

	// part 2
	const B = 1_000_000_000
	cache := map[string]int{}
	result := map[int][][]byte{}
	current, previous := 0, 0
	for ; ; current++ {
		if cache[key(board)] > 0 {
			if previous == 0 {
				// first time repeat, reset the cache
				previous = current
				cache = map[string]int{}
			} else {
				// second time repeat, stop
				break
			}
		} else {
			cache[key(board)] = current
			result[current] = deepCopy(board)
		}

		// North
		moveNorth(board)
		//West = rotate + move north
		board = rotate(board)
		moveNorth(board)
		board = rotate(rotate(rotate(board)))
		//South = rotate x2 + move north
		board = rotate(rotate(board))
		moveNorth(board)
		board = rotate(rotate(board))
		// East = rotate x3 + move north
		board = rotate(rotate(rotate(board)))
		moveNorth(board)
		board = rotate(board)
	}

	period := current - previous - 1
	fmt.Println("Total load after a billion rotations", calc(result[previous+(B-previous)%period]))
}
