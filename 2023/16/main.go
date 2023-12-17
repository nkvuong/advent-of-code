package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

var board [][]byte

type coord struct {
	row int
	col int
}

type beam struct {
	pos coord
	dir coord
}

func (c coord) isValid() bool {
	return c.row >= 0 && c.col >= 0 && c.row < len(board) && c.col < len(board[0])
}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func shine(toDo []beam) int {
	visited := make(map[beam]bool, len(input))
	for len(toDo) > 0 {
		curr := toDo[len(toDo)-1]
		toDo = toDo[:len(toDo)-1]
		for {
			if _, ok := visited[curr]; ok {
				// already visited
				break
			}
			visited[curr] = true

			pos := coord{curr.pos.row + curr.dir.row, curr.pos.col + curr.dir.col}
			if !pos.isValid() {
				// out of bound
				break
			}

			dir := curr.dir
			switch board[pos.row][pos.col] {
			case '\\':
				dir = coord{curr.dir.col, curr.dir.row}
			case '/':
				dir = coord{-curr.dir.col, -curr.dir.row}
			case '|':
				toDo = append(toDo, beam{pos, coord{1, 0}})
				dir = coord{-1, 0}
			case '-':
				toDo = append(toDo, beam{pos, coord{0, 1}})
				dir = coord{0, -1}
			}
			curr = beam{pos, dir}
		}
	}
	energised := make(map[coord]bool, len(input))

	for key := range visited {
		energised[key.pos] = true
	}
	return len(energised) - 1
}

func main() {
	lines := strings.Split(input, "\n")
	board = make([][]byte, 0)
	for _, line := range lines {
		board = append(board, []byte(line))
	}
	fmt.Println("Total energised cells is", shine(
		[]beam{
			{
				coord{0, -1}, coord{0, 1},
			},
		},
	))
	maxCells := 0
	for i := 0; i < len(board); i++ {
		maxCells = max(maxCells, shine(
			[]beam{
				{
					coord{i, -1}, coord{0, 1},
				},
			},
		))
		maxCells = max(maxCells, shine(
			[]beam{
				{
					coord{i, len(board)}, coord{0, -1},
				},
			},
		))
	}
	for i := 0; i < len(board[0]); i++ {
		maxCells = max(maxCells, shine(
			[]beam{
				{
					coord{-1, i}, coord{1, 0},
				},
			},
		))
		maxCells = max(maxCells, shine(
			[]beam{
				{
					coord{len(board), i}, coord{-1, 0},
				},
			},
		))
	}
	fmt.Printf("Max configuration has %d energised cells", maxCells)
}
