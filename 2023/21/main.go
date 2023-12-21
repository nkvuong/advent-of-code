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

var gardenMap = make([][]byte, 0)
var height, width int

func visit(start [2]int, steps int, infinite bool) int {
	var neighbours = [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	in := map[[2]int]struct{}{start: {}}
	for i := 1; i <= steps; i++ {
		out := make(map[[2]int]struct{})
		for plot := range in {
			for _, neighbour := range neighbours {
				new := [2]int{plot[0] + neighbour[0], plot[1] + neighbour[1]}
				var garden byte
				if infinite {
					// for infinite garden, we don't check for boundary
					garden = gardenMap[(new[0]+steps*height)%height][(new[1]+steps*width)%width]
				} else {
					if new[0] < 0 || new[1] < 0 || new[0] >= height || new[1] >= width {
						//out of bound
						continue
					}
					garden = gardenMap[new[0]][new[1]]
				}
				if garden != '#' {
					out[new] = struct{}{}
				}
			}
		}
		in = out
	}
	return len(in)
}

func main() {
	lines := strings.Split(input, "\n")
	var start [2]int
	for i, line := range lines {
		gardenMap = append(gardenMap, []byte(line))
		if strings.Contains(line, "S") {
			start = [2]int{i, strings.Index(line, "S")}
		}
	}
	height, width = len(lines), len(lines[0])
	fmt.Println(visit(start, 64, false))
	// part 2
	// notice that 26501365 = 202300 * 131 + 65
	// number of steps formed a diamond, growing larger
	// it will be a quadratic formula of the form a + b*x + c*x^2
	// we'll calculate it for 0, 1, 2 to calculate the coefficients
	// p(0) = a, p(1) = a+b+c, p(2) = a+2b+4c
	// so b = (4*p(1) - p(2)-3a)/2 and c = (p(2)-2p(1)+a)/2
	half := width / 2
	var polynomial []int
	for _, step := range []int{half, half + width, half + 2*width} {
		polynomial = append(polynomial, visit(start, step, true))
	}
	a := polynomial[0]
	b := (4*polynomial[1] - polynomial[2] - 3*a) / 2
	c := (polynomial[2] - 2*polynomial[1] + a) / 2
	n := (26501365 - half) / width
	fmt.Println(a + b*n + c*n*n)
}
