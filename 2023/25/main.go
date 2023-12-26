package main

import (
	_ "embed"
	"fmt"
	"math/rand"
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

type Edge struct {
	src, dest string
}

type Component struct {
	parent string
	rank   int
}

// A utility function to find set of an element i
// (uses path compression technique)
func find(components map[string]Component, i string) string {
	// find root and make root as parent of i
	// (path compression)
	if components[i].parent != i {
		components[i] = Component{find(components, components[i].parent), components[i].rank}
	}

	return components[i].parent
}

// A function that does union of two sets of x and y
// (uses union by rank)
func Union(components map[string]Component, x, y string) {
	xroot := find(components, x)
	yroot := find(components, y)

	// Attach smaller rank tree under root of high
	// rank tree (Union by Rank)
	if components[xroot].rank < components[yroot].rank {
		components[xroot] = Component{yroot, components[xroot].rank}
	} else if components[xroot].rank > components[yroot].rank {
		components[yroot] = Component{xroot, components[yroot].rank}
	} else {
		// If ranks are same, then make one as root and
		// increment its rank by one
		components[yroot] = Component{xroot, components[yroot].rank}
		components[xroot] = Component{components[xroot].parent, components[xroot].rank + 1}
	}
}

func main() {
	lines := strings.Split(input, "\n")
	edges := []Edge{}
	vertices := make(map[string]struct{})
	components := map[string]Component{}
	for _, line := range lines {
		name, connected, _ := strings.Cut(line, ": ")
		destinations := strings.Split(connected, " ")
		vertices[name] = struct{}{}
		for _, dest := range destinations {
			edges = append(edges, Edge{name, dest})
			vertices[dest] = struct{}{}
		}
	}
	for {
		// repeat Karger's algorithm until we find exactly 3 cuts
		for v := range vertices {
			// Create V subsets with single element
			components[v] = Component{v, 0}
		}
		verticesNum := len(vertices)
		for verticesNum > 2 {
			// Karger's algorithm - contracting random edges until there are 2 vertices left
			random := rand.Intn(len(edges))
			edge := edges[random]

			componentSrc := find(components, edge.src)
			componentDest := find(components, edge.dest)

			if componentSrc == componentDest {
				// If two corners belong to same subset,
				// then ignore this edge
				continue
			} else {
				// Else contract the edge (or combine the
				// corners of edge into one vertex)
				Union(components, componentSrc, componentDest)
				verticesNum--
			}
		}
		edgesCut := 0
		for _, edge := range edges {
			//count number of edges to cut - if an edge cross components then
			subset1 := find(components, edge.src)
			subset2 := find(components, edge.dest)
			if subset1 != subset2 {
				edgesCut++
			}
		}
		if edgesCut == 3 {
			break
		}
	}
	componentSize := make(map[string]int)
	for v := range vertices {
		component := find(components, v)
		componentSize[component]++
	}
	product := 1
	for _, v := range componentSize {
		product *= v
	}
	fmt.Println(product)
}
