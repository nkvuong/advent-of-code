package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
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

func parseHistory(history string) []int {
	var out []int
	r := regexp.MustCompile(`-?\d+`)
	matches := r.FindAllString(history, -1)
	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		out = append(out, num)
	}
	return out
}

func sum(in []int, abs bool) int {
	total := 0
	for _, num := range in {
		if abs {
			total += int(math.Abs(float64(num)))
		} else {
			total += num
		}
	}
	return total
}

func findNext(history []int) (int, int) {
	var diffRight []int
	var diffLeft []int
	for sum(history, true) != 0 {
		tmp := []int{}
		diffRight = append(diffRight, history[len(history)-1])
		diffLeft = append(diffLeft, history[0])
		for i := 0; i < len(history)-1; i++ {
			tmp = append(tmp, history[i+1]-history[i])
		}
		history = tmp
	}
	// need to alternate the signs for the odd previous values, as they are subtracted
	for i := 1; i < len(diffLeft); i += 2 {
		diffLeft[i] = diffLeft[i] * (-1)
	}
	return sum(diffLeft, false), sum(diffRight, false)
}

func main() {

	histories := strings.Split(input, "\n")
	totalNext := 0
	totalPrev := 0
	for _, history := range histories {
		prev, next := findNext(parseHistory(history))
		totalNext += next
		totalPrev += prev
	}
	fmt.Println("Total sum of next histories", totalNext)
	fmt.Println("Total sum of previous histories", totalPrev)
}
