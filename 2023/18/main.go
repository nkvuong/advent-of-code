package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

type coord struct {
	row int
	col int
}

func (cur coord) add(that coord) coord {
	return coord{cur.row + that.row, cur.col + that.col}
}

type instruction struct {
	dir    string
	num    int
	colour string
}

var directions = map[string]coord{
	"D": {1, 0},
	"L": {0, -1},
	"U": {-1, 0},
	"R": {0, 1},
}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func shoelace(polygon []coord) int {
	sum := 0
	p0 := polygon[len(polygon)-1]
	for _, p1 := range polygon {
		sum += p0.col*p1.row - p0.row*p1.col
		p0 = p1
	}
	return sum / 2
}

func main() {

	lines := strings.Split(input, "\n")
	instructions := []instruction{}

	for _, line := range lines {
		splits := strings.Split(line, " ")
		dir, colour := splits[0], strings.TrimSuffix(splits[2], ")")
		num, _ := strconv.Atoi(splits[1])
		instructions = append(instructions, instruction{dir, num, colour})
	}
	//part 1
	pos := coord{0, 0}
	polygon := []coord{{0, 0}}
	for _, ins := range instructions {
		for i := 1; i <= ins.num; i++ {
			pos = pos.add(directions[ins.dir])
			polygon = append(polygon, pos)
		}
	}
	// Pick's theorem
	sum := shoelace(polygon) + len(polygon)/2 + 1

	fmt.Println(sum)

	//part 2
	intToDir := []string{"R", "D", "L", "U"}
	pos = coord{0, 0}
	polygon = []coord{{0, 0}}
	for _, ins := range instructions {
		colour := ins.colour
		direction := intToDir[int(colour[len(colour)-1]-'0')]
		num, _ := strconv.ParseInt(colour[2:7], 16, 0)
		for i := 1; i <= int(num); i++ {
			pos = pos.add(directions[direction])
			polygon = append(polygon, pos)
		}
	}
	sum = shoelace(polygon) + len(polygon)/2 + 1

	fmt.Println(sum)
}
