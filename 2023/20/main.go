package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

type module struct {
	t      rune
	dest   []string
	memory map[string]int
}

type pulse struct {
	dest  string
	power int
}

var modules = make(map[string]module)
var cycle = make(map[string]int)

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

	lines := strings.Split(input, "\n")
	var goal string
	for _, line := range lines {
		pre, post, _ := strings.Cut(line, " -> ")
		modules[pre[1:]] = module{
			t:      rune(pre[0]),
			dest:   strings.Split(post, ", "),
			memory: make(map[string]int),
		}
		// log the previous module before rx
		// assumption - there is only one such module before rx
		if post == "rx" {
			goal = pre[1:]
		}
	}
	for _, line := range lines {
		pre, post, _ := strings.Cut(line, " -> ")
		dest := strings.Split(post, ", ")
		// set the memory of all conjunction module to low pulse
		for _, d := range dest {
			if modules[d].t == '&' {
				modules[d].memory[pre[1:]] = 0
			}
		}
	}
	high, low := 0, 0
	for i := 1; i <= 100000; i++ {
		if i == 1001 {
			fmt.Println(low * high)
		}
		queue := []pulse{{dest: "roadcaster", power: 0}}
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			if curr.power == 0 {
				low++
			} else {
				high++
			}
			if curr.dest == "output" {
				// already at output
				continue
			}
			power := 0
			module := modules[curr.dest]
			if module.t == '%' {
				// flip flop module turns on/off when low pulse received
				if curr.power == 0 {

					if len(module.memory) == 0 {
						power = 1
						module.memory["on"] = 1
					} else {
						delete(module.memory, "on")
						power = 0
					}
					for _, out := range module.dest {
						queue = append(queue, pulse{out, power})
						if modules[out].t == '&' {
							// set memory of the destination module
							modules[out].memory[curr.dest] = power
							if power == 1 && cycle[curr.dest] == 0 {
								// assumption - there is a cycle when a high pulse is emitted
								cycle[curr.dest] = i
							}
						}
					}
				}
			} else {
				// conjunction module has memory, need to check if all input are high
				memory := 0
				for _, m := range module.memory {
					memory += m
				}
				if memory == len(module.memory) {
					//all high input, send low
					power = 0
				} else {
					// otherwise send high
					power = 1
				}
				for _, out := range module.dest {
					queue = append(queue, pulse{out, power})
					if modules[out].t == '&' {
						// set memory of the destination module
						modules[out].memory[curr.dest] = power
						if power == 1 && cycle[curr.dest] == 0 {
							// assumption - there is a cycle when a high pulse is emitted
							cycle[curr.dest] = i
						}
					}
				}
			}
		}
	}
	ans := 1
	for d := range modules[goal].memory {
		// assumption - it is a conjunction module before rx
		// all incoming pulses must be high, for a low pulse to emit
		ans = lcm(ans, cycle[d])
	}
	fmt.Println(ans)
}
