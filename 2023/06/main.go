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

type race struct {
	time   int
	record int
}

func (r race) getRecords() int {
	newRecord := 0
	for i := 1; i < r.time; i++ {
		if i*(r.time-i) > r.record {
			newRecord++
		}
	}
	return newRecord
}

func stringToIntSlice(in string) []int {
	var out []int
	r := regexp.MustCompile(`\d+`)
	matches := r.FindAllString(in, -1)
	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		out = append(out, num)
	}
	return out
}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {

	splits := strings.Split(input, "\n")
	timeString, recordString := splits[0], splits[1]

	//part 1
	times := stringToIntSlice(timeString)
	records := stringToIntSlice(recordString)

	product := 1
	for i := 0; i < len(times); i++ {
		product = product * race{time: times[i], record: records[i]}.getRecords()
	}
	fmt.Println("Product of all numbers is", product)

	//part 2
	time, _ := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(timeString, " ", ""), "Time:", ""))
	record, _ := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(recordString, " ", ""), "Distance:", ""))
	fmt.Println("For the long race", race{time: time, record: record}.getRecords())
}
