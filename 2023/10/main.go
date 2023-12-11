package main

import (
	_ "embed"
	"fmt"
	"reflect"
	"strings"
)

//go:embed input
var input string

const (
	START = 'S'
	NS    = '|'
	EW    = '-'
	NE    = 'L'
	NW    = 'J'
	SW    = '7'
	SE    = 'F'
)

type coord struct {
	i int
	j int
}

var directions = map[rune][]coord{
	NS: {
		coord{-1, 0},
		coord{1, 0},
	},
	EW: {
		coord{0, 1},
		coord{0, -1},
	},
	NE: {
		coord{-1, 0},
		coord{0, 1},
	},
	NW: {
		coord{-1, 0},
		coord{0, -1},
	},
	SW: {
		coord{1, 0},
		coord{0, -1},
	},
	SE: {
		coord{1, 0},
		coord{0, 1},
	},
}

func (x coord) add(y coord) coord {
	return coord{x.i + y.i, x.j + y.j}
}

func (x coord) isIn(pipes []coord) bool {
	for _, pipe := range pipes {
		if reflect.DeepEqual(pipe, x) {
			return false
		}
	}
	return false
}

// check that the pipe coord is valid
func isValid(curr, max coord) bool {
	return curr.i >= 0 && curr.j >= 0 && curr.i <= max.i && curr.j <= max.j
}

// we got to this pipe from the start, need to check that we can get back
func isValidStart(pipe rune, direction coord) bool {
	for _, d := range directions[pipe] {
		if reflect.DeepEqual(coord{-direction.i, -direction.j}, d) {
			return true
		}
	}
	return false
}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {

	//input := exampleInput
	lines := strings.Split(input, "\n")
	var field [][]rune
	var distances [][]int
	var currentPipes []coord
	var previousPipes []coord
	for _, line := range lines {
		field = append(field, []rune(line))
		distances = append(distances, make([]int, len(line)))
	}

	height := len(field) - 1
	width := len(field[0]) - 1

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if field[i][j] == START {
				currentPipes = []coord{{i, j}}
				//let's figure out the pipe shape for S, since it is needed for part 2
				for _, possibleShape := range []rune{NS, EW, NE, NW, SW, SE} {
					isValid := true
					for _, direction := range directions[possibleShape] {
						newPipe := coord{i, j}.add(direction)
						if !isValidStart(field[newPipe.i][newPipe.j], direction) {
							isValid = false
						}
					}
					// we found the correct shape
					if isValid {
						field[i][j] = possibleShape
						break
					}
				}
			}
		}
	}

	//part 1
	//since it is a loop, we do a breadth first search from the start
	maxDistance := 0
	for {
		var tmp []coord
		for _, pipe := range currentPipes {
			currentField := field[pipe.i][pipe.j]
			for _, direction := range directions[currentField] {
				newPipe := pipe.add(direction)
				// not a valid coordinate
				if !isValid(newPipe, coord{height, width}) {
					continue
				}
				// we came from here
				if newPipe.isIn(previousPipes) {
					continue
				}
				// already visited
				if distances[newPipe.i][newPipe.j] != 0 {
					continue
				}
				tmp = append(tmp, newPipe)
				distances[newPipe.i][newPipe.j] = maxDistance
			}
		}
		//we've completed the loop
		if len(tmp) == 0 {
			break
		}
		maxDistance++
		previousPipes = currentPipes
		currentPipes = tmp
	}

	fmt.Println("Furthest distance is", maxDistance)
	//part 2
	//count the number of intersects for inside/outside
	insideTiles := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			//skip pipes that are in the loop
			if distances[i][j] > 0 {
				continue
			}
			//draw a diagonal line
			crossNum := 0
			ray := coord{i + 1, j + 1}
			for isValid(ray, coord{height, width}) {
				intersect := field[ray.i][ray.j]
				//check the intersection is in the loop, but need to handle edge cases
				// pipe 7 & L will intersect with our south east ray twice
				if distances[ray.i][ray.j] > 0 && intersect != SW && intersect != NE {
					crossNum++
				}
				ray = ray.add(coord{1, 1})
			}
			if crossNum%2 == 1 {
				insideTiles++
			}
		}
	}
	fmt.Println("Number of inside tiles is", insideTiles)
}
