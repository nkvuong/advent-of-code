package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

type node struct {
	left  string
	right string
}

func navigate(curr string, instruction string, maps map[string]node, finalNode string) (string, int) {
	for i, step := range instruction {
		if strings.HasSuffix(curr, finalNode) {
			return curr, int(i)
		}
		if step == 'L' {
			curr = maps[curr].left
		} else {
			curr = maps[curr].right
		}
	}
	return curr, len(instruction)
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func lcm(integers ...int) int {
	result := integers[0] * integers[1] / gcd(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {

	maps := make(map[string]node)
	instruction, nodes, _ := strings.Cut(input, "\n\n")
	var currList []string
	var stepCount []int
	for _, n := range strings.Split(nodes, "\n") {
		name, direction, _ := strings.Cut(n, " = (")
		left, right, _ := strings.Cut(strings.TrimSuffix(direction, ")"), ", ")
		maps[name] = node{left, right}
		if strings.HasSuffix(name, "A") {
			currList = append(currList, name)
			stepCount = append(stepCount, 0)
		}
	}
	//part 1
	count := 0
	curr := "AAA"
	for !strings.HasSuffix(curr, "ZZZ") {
		var steps int
		curr, steps = navigate(curr, instruction, maps, "ZZZ")
		count += steps
	}
	fmt.Printf("Take %d steps to reach ZZZ\n", count)
	//part 2
	for i, curr := range currList {
		//calculate minimum number of steps for each ghost to reach a Z
		for !strings.HasSuffix(curr, "Z") {
			var steps int
			curr, steps = navigate(curr, instruction, maps, "Z")
			stepCount[i] += steps
		}
	}
	//take the LCM of all the steps (as it should be cyclic)
	fmt.Printf("Take %d steps to reach **Z\n", lcm(stepCount...))
}
