package main

import (
	_ "embed"
	"fmt"
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

type card struct {
	number  int
	winners map[int]bool
	numbers map[int]bool
}

func (c card) getPoints() (int, int) {
	points := 1
	matches := 0
	for number := range c.numbers {
		if c.winners[number] {
			matches++
			points *= 2
		}
	}
	return points / 2, matches
}

func parseNumbers(input string) map[int]bool {
	output := make(map[int]bool)
	r := regexp.MustCompile(`\d+`)
	matches := r.FindAllString(input, -1)
	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		output[num] = true
	}
	return output
}

func buildCard(input string) card {
	cards := strings.Split(input, ":")
	number, _ := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(cards[0], "Card ", ""), " ", ""))
	content := strings.Split(cards[1], "|")
	return card{
		number:  number,
		winners: parseNumbers(content[0]),
		numbers: parseNumbers(content[1]),
	}
}

func main() {
	sum := 0
	tally := make(map[int]int)
	cardsStr := strings.Split(input, "\n")
	numCards := strings.Count(input, "\n") + 1
	for _, cardStr := range cardsStr {
		c := buildCard(cardStr)
		// get the number of points and matches of a card
		points, matches := c.getPoints()
		sum += points
		tally[c.number] += 1
		// add new cards to the tally
		for i := c.number + 1; i <= min(c.number+matches, numCards); i++ {
			tally[i] += tally[c.number]
		}
	}
	fmt.Println("Sum of scratchcards points is", sum)
	sum = 0
	for _, num := range tally {
		sum += num
	}
	fmt.Println("Total number of scratchcards", sum)
}
