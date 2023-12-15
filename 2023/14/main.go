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

func main() {

	load := 0
	height := strings.Count(input, "\n") + 1
	roundCount := make(map[int]int)
	col := 0
	row := 0
	for _, rock := range input {
		switch {
		case
			rock == '\n':
			{
				col = 0
				row++
				continue
			}
		case rock == 'O':
			{
				roundCount[col]++
				load += height - roundCount[col] + 1
			}
		case rock == '#':
			{
				roundCount[col] = row + 1
			}
		}
		col++
	}
	fmt.Println("Total load is", load)

}
