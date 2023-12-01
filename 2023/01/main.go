package main

import (
	_ "embed"
	"fmt"
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

func main() {
	lines := strings.Split(input, "\n")
	//part 1
	sum := 0
	for _, line := range lines {
		//add digits to a separate slice
		var numbers []int
		for _, r := range line {
			if unicode.IsDigit(r) {
				numbers = append(numbers, int(r-'0'))
			}
		}
		if len(numbers) > 0 {
			sum += numbers[0]*10 + numbers[len(numbers)-1]
		}
	}
	fmt.Println("Part 1: Sum of all of the calibration values is", sum)
	digits := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	sum = 0
	for _, line := range lines {
		var numbers []int
		//this slice tracks location of written digits
		locations := make([]int, len(line))

		// we only care about first & last instance of the digit
		for num, digit := range digits {
			first := strings.Index(line, digit)
			if first > -1 {
				locations[first] = num + 1
			}
			last := strings.LastIndex(line, digit)
			if last > -1 {
				locations[last] = num + 1
			}
		}
		for i, r := range line {
			//either it is a written digit
			if locations[i] > 0 {
				numbers = append(numbers, locations[i])
			}
			//or a digit
			if unicode.IsDigit(r) {
				numbers = append(numbers, int(r-'0'))
			}
		}
		sum += numbers[0]*10 + numbers[len(numbers)-1]
	}
	fmt.Println("Part 2: Sum of all of the calibration values is", sum)
}
