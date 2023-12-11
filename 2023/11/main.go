package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

type galaxy struct {
	i int
	j int
}

const SPACE = '.'
const GALAXY = '#'

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func absDiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func (curr galaxy) distance(other galaxy, galaxyRows, galaxyColumns map[int]bool, rate int) int {
	expansion := 0
	// need to account for all empty rows that eill expand
	for i := min(curr.i, other.i); i < max(curr.i, other.i); i++ {
		if _, ok := galaxyRows[i]; !ok {
			expansion += rate
		}
	}
	// same for the columns
	for j := min(curr.j, other.j); j < max(curr.j, other.j); j++ {
		if _, ok := galaxyColumns[j]; !ok {
			expansion += rate
		}
	}
	return absDiff(curr.i, other.i) + absDiff(curr.j, other.j) + expansion
}

func main() {

	lines := strings.Split(input, "\n")
	var galaxies []galaxy
	galaxyRows := make(map[int]bool)
	galaxyColumns := make(map[int]bool)
	for i, line := range lines {
		for j, char := range line {
			if char == GALAXY {
				galaxies = append(galaxies, galaxy{i, j})
				galaxyRows[i] = true
				galaxyColumns[j] = true
			}
		}
	}
	//part 1
	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += galaxies[i].distance(galaxies[j], galaxyRows, galaxyColumns, 1)
		}
	}

	fmt.Println("Total distance is", sum)
	//part 2
	sum = 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += galaxies[i].distance(galaxies[j], galaxyRows, galaxyColumns, 1000000-1)
		}
	}

	fmt.Println("Total distance of old universe is", sum)

}
