package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

type coord struct {
	x int
	y int
}

func main() {
	schematic := strings.Replace(input, "\n", "", -1)
	height := strings.Count(input, "\n") + 1
	length := len(schematic) / height
	//part 1
	sum := 0
	//find all numbers
	r := regexp.MustCompile(`\d+`)
	matches := r.FindAllStringIndex(schematic, -1)
	var partNumbers [][]int
	for _, match := range matches {
		locationStart := match[0]
		locationEnd := match[1]
		if isPartNumber(schematic, serialise(locationStart, length), locationEnd-locationStart, serialise(len(schematic)-1, length)) {
			num, _ := strconv.Atoi(schematic[locationStart:locationEnd])
			partNumbers = append(partNumbers, match)
			sum += num
		}
	}

	fmt.Println("Sum of all part numbers is", sum)
	//part 2
	sum = 0
	// find all * characters
	for location, r := range schematic {
		if r == '*' {
			// get all part neighbours
			partNeighbours := getPartNeighbours(schematic, serialise(location, length), serialise(len(schematic)-1, length), partNumbers)
			if len(partNeighbours) == 2 {
				left, _ := strconv.Atoi(schematic[partNeighbours[0][0]:partNeighbours[0][1]])
				right, _ := strconv.Atoi(schematic[partNeighbours[1][0]:partNeighbours[1][1]])
				sum += left * right
			}
		}
	}
	fmt.Println("Sum of all gear ratios is", sum)
}

func isPartNumber(schematic string, cur coord, length int, max coord) bool {
	var neighbours []coord
	// add the above & below neighbours
	for k := cur.x - 1; k <= cur.x+length; k++ {
		neighbours = append(neighbours, coord{k, cur.y + 1}, coord{k, cur.y - 1})
	}
	//add the left & right neighbours
	neighbours = append(neighbours, coord{cur.x - 1, cur.y}, coord{cur.x + length, cur.y})
	//check if each neighbour for valid coords and if they contain a symbol
	for _, neighbour := range neighbours {
		if isValid(neighbour, max) && isSymbol(schematic, neighbour.deserialise(max.x)) {
			return true
		}
	}
	return false
}

func (c coord) deserialise(length int) int {
	return c.y*(length+1) + c.x
}

func serialise(location int, length int) coord {
	return coord{
		x: location % length,
		y: location / length,
	}
}

// get all part numbers that are neighbours to a coordinate
func getPartNeighbours(schematic string, cur coord, max coord, parts [][]int) [][]int {
	// list all neighbours
	var neighbours []coord
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				if isValid(coord{cur.x + i, cur.y + j}, max) {
					neighbours = append(neighbours, coord{cur.x + i, cur.y + j})
				}
			}
		}
	}
	var partNeighbours [][]int
	for _, part := range parts {
		//check if a neighbour is in the part by serialising its coordinate
		for _, neighbour := range neighbours {
			if neighbour.deserialise(max.x) >= part[0] && neighbour.deserialise(max.x) < part[1] {
				partNeighbours = append(partNeighbours, part)
				break
			}
		}
	}
	return partNeighbours
}

// check if a location in the schematic is a symbol (not a digit or .)
func isSymbol(schematic string, x int) bool {
	r := schematic[x]
	if unicode.IsDigit(rune(r)) || r == '.' {
		return false
	}
	return true
}

// check if a coordinate is valid
func isValid(cur coord, max coord) bool {
	return cur.x >= 0 && cur.y >= 0 && cur.x <= max.x && cur.y <= max.y
}
