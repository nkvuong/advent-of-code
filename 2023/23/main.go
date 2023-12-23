package main

import (
	_ "embed"
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
)

//go:embed input
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

var hikingMap [][]byte
var height, width int

func dfs(plot [2]int, edges map[[2]int]map[[2]int]int) int {

	visited, maxLength := make(map[[2]int]bool), 0
	stack := [][3]int{{plot[0], plot[1], 0}}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		currentPlot, length := [2]int{current[0], current[1]}, current[2]
		if current[2] == -1 {
			// backtrack
			visited[currentPlot] = false
			continue
		}
		if visited[currentPlot] {
			// already visited
			continue
		}
		visited[currentPlot] = true
		// remind ourselves to backtrack
		stack = append(stack, [3]int{currentPlot[0], currentPlot[1], -1})
		for neighbour, value := range edges[currentPlot] {
			// need to visit all the neighbours
			stack = append(stack, [3]int{neighbour[0], neighbour[1], length + value})
		}
		if current[0] == height-1 {
			// we are at our destination
			maxLength = max(maxLength, length)
		}
	}
	return maxLength
}

func main() {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		hikingMap = append(hikingMap, []byte(line))
	}
	height, width = len(hikingMap), len(hikingMap[0])
	edges := make(map[[2]int]map[[2]int]int)

	// part 1
	for i := 0; i < len(hikingMap); i++ {
		for j := 0; j < len(hikingMap[i]); j++ {
			plot := hikingMap[i][j]
			edges[[2]int{i, j}] = make(map[[2]int]int)
			neighbours := [][2]int{
				{0, 1},
				{0, -1},
				{1, 0},
				{-1, 0},
			}
			switch plot {
			case '^':
				neighbours = [][2]int{{-1, 0}}
			case '>':
				neighbours = [][2]int{{0, 1}}
			case '<':
				neighbours = [][2]int{{0, -1}}
			case 'v':
				neighbours = [][2]int{{1, 0}}
			}
			for _, neighbour := range neighbours {
				new := [2]int{i + neighbour[0], j + neighbour[1]}
				if new[0] < 0 || new[1] < 0 || new[0] >= height || new[1] >= width {
					//out of bound
					continue
				}
				if hikingMap[new[0]][new[1]] != '#' {
					// add to edges if new tile is not forest
					edges[[2]int{i, j}][new] = 1
				}
			}
		}
	}
	for i := 0; i < len(hikingMap[0]); i++ {
		if hikingMap[0][i] == '.' {
			fmt.Println("Longest hike is", dfs([2]int{0, i}, edges))
		}
	}

	// part 2

	edges = make(map[[2]int]map[[2]int]int)
	for i := 0; i < len(hikingMap); i++ {
		for j := 0; j < len(hikingMap[i]); j++ {
			if hikingMap[i][j] == '#' {
				// forest tile
				continue
			}
			edges[[2]int{i, j}] = make(map[[2]int]int)
			neighbours := [][2]int{
				{0, 1},
				{0, -1},
				{1, 0},
				{-1, 0},
			}
			for _, neighbour := range neighbours {
				new := [2]int{i + neighbour[0], j + neighbour[1]}
				if new[0] < 0 || new[1] < 0 || new[0] >= height || new[1] >= width {
					//out of bound
					continue
				}
				if hikingMap[new[0]][new[1]] != '#' {
					// add to edges if new tile is not forest
					edges[[2]int{i, j}][new] = 1
				}
			}
		}
	}
	for {
		contracted := false
		for node, edge := range edges {
			if len(edge) == 2 {
				// perform edge contraction for any node with just 2 edges
				// a-b-c -> a-c
				nodes := maps.Keys(edge)
				n1, n2 := nodes[0], nodes[1]
				edges[n1][n2] = edges[n1][node] + edges[node][n2]
				edges[n2][n1] = edges[n2][node] + edges[node][n1]
				delete(edges[n1], node)
				delete(edges[n2], node)
				delete(edges, node)
				contracted = true
				break
			}
		}
		if !contracted {
			break
		}
	}
	for i := 0; i < len(hikingMap[0]); i++ {
		if hikingMap[0][i] == '.' {
			fmt.Println("Longest dry hike is", dfs([2]int{0, i}, edges))
		}
	}

}
