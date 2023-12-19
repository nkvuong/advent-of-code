package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input
var input string

var workflows = make(map[string][]string)

type part struct {
	x int
	m int
	a int
	s int
}

func atoi(in string) int {
	out, _ := strconv.Atoi(in)
	return out
}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func apply(rule string, item part) string {
	if !strings.Contains(rule, ":") {
		return rule
	}
	check, out, _ := strings.Cut(rule, ":")
	var value int
	switch check[0] {
	case 'x':
		value = item.x
	case 'm':
		value = item.m
	case 'a':
		value = item.a
	case 's':
		value = item.s
	}
	gate := atoi(check[2:])
	switch check[1] {
	case '<':
		if value < gate {
			return out
		}
	case '>':
		if value > gate {
			return out
		}
	}
	return ""
}

func count(workflow string, rangeValues map[rune][2]int) int {
	if workflow == "R" {
		return 0
	} else if workflow == "A" {
		product := 1
		for _, v := range rangeValues {
			product *= (v[1] - v[0] + 1)
		}
		return product
	}

	total := 0
	for _, rule := range workflows[workflow] {
		if !strings.Contains(rule, ":") {
			// there is no check
			total += count(rule, rangeValues)
			continue
		}
		check, out, _ := strings.Cut(rule, ":")
		category := rune(check[0])
		gate := atoi(check[2:])
		v := rangeValues[category]
		var trueRange, falseRange [2]int
		// calculate the range of values where check is valid/invalid
		if check[1] == '<' {
			trueRange = [2]int{v[0], gate - 1}
			falseRange = [2]int{gate, v[1]}
		} else {
			trueRange = [2]int{gate + 1, v[1]}
			falseRange = [2]int{v[0], gate}
		}
		// for true range, create a clone and keep counting
		// but start at the rule output
		if trueRange[0] <= trueRange[1] {
			cloneRange := make(map[rune][2]int, 4)
			for k, v := range rangeValues {
				cloneRange[k] = v
			}
			cloneRange[category] = trueRange
			total += count(out, cloneRange)
		}

		if falseRange[0] > falseRange[1] {
			// no range to process
			break
		}

		// for false range, keep processing rest of rules
		rangeValues[category] = falseRange
	}

	return total
}

func main() {

	workflowsString, parts, _ := strings.Cut(input, "\n\n")
	for _, w := range strings.Split(workflowsString, "\n") {
		name, ruleString, _ := strings.Cut(strings.TrimSuffix(w, "}"), "{")
		rules := strings.Split(ruleString, ",")
		workflows[name] = rules
	}
	sum := 0
	for _, p := range strings.Split(parts, "\n") {
		r := regexp.MustCompile(`-?\d+`)
		matches := r.FindAllString(p, -1)
		item := part{x: atoi(matches[0]), m: atoi(matches[1]), a: atoi(matches[2]), s: atoi(matches[3])}
		currentRule := "in"
		for {
			rules := workflows[currentRule]
			for _, rule := range rules {
				currentRule = apply(rule, item)
				if currentRule != "" {
					break
				}
			}
			if slices.Contains([]string{"R", "A"}, currentRule) {
				break
			}
		}
		if currentRule == "A" {
			sum += item.x + item.m + item.a + item.s
		}
	}
	fmt.Println(sum)

	fmt.Println(count("in", map[rune][2]int{
		'x': {1, 4000},
		'm': {1, 4000},
		'a': {1, 4000},
		's': {1, 4000},
	}))
}
