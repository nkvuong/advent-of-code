package main

import (
	_ "embed"
	"fmt"
	"strconv"
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

func main() {
	games := strings.Split(input, "\n")
	//part 1
	limit := map[string](int){
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sum := 0
	for _, game := range games {
		game, sets, _ := strings.Cut(game, ": ")
		id, _ := strconv.Atoi(strings.ReplaceAll(game, "Game ", ""))
		valid := true
		for _, set := range strings.Split(sets, "; ") {
			valid = valid && IsValid(set, limit)
		}
		if valid {
			sum += id
		}
	}
	fmt.Println("Sum of possible IDs is", sum)
	//part 2
	sum = 0
	for _, game := range games {
		_, sets, _ := strings.Cut(game, ": ")
		limit = map[string](int){
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, set := range strings.Split(sets, "; ") {
			limit = GetMin(set, limit)
		}
		sum += limit["red"] * limit["green"] * limit["blue"]
	}
	fmt.Println("Power of these sets is", sum)
}

// check if a set is valid given a limit
func IsValid(set string, limit map[string](int)) bool {
	balls := strings.Split(set, ", ")
	for _, ball := range balls {
		n, colour, _ := strings.Cut(ball, " ")
		num, _ := strconv.Atoi(n)
		if num > limit[colour] {
			return false
		}
	}
	return true
}

// calculate the new minimum number of balls required
func GetMin(game string, limit map[string](int)) map[string](int) {
	balls := strings.Split(game, ", ")
	for _, ball := range balls {
		n, colour, _ := strings.Cut(ball, " ")
		num, _ := strconv.Atoi(n)
		if num > limit[colour] {
			limit[colour] = num
		}
	}
	return limit
}
