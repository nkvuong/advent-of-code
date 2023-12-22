package main

import (
	_ "embed"
	"fmt"
	"sort"
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

type Coord struct {
	x int
	y int
	z int
}

func strToCoord(in string) Coord {
	split := strings.Split(in, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	z, _ := strconv.Atoi(split[2])
	return Coord{x, y, z}
}

type Brick struct {
	posStart Coord
	posEnd   Coord
}

var space [11][11][501]int

func getZ(brick Brick) int {
	var z int
	for z = brick.posStart.z; z >= 1; z-- {
		for x := brick.posStart.x; x <= brick.posEnd.x; x++ {
			for y := brick.posStart.y; y <= brick.posEnd.y; y++ {
				if space[x][y][z] != 0 {
					// we found the lowest we can go
					return z + 1
				}
			}
		}
	}
	return z + 1
}

func updateSpace(i, z1, z2 int, brick Brick) {
	for x := brick.posStart.x; x <= brick.posEnd.x; x++ {
		for y := brick.posStart.y; y <= brick.posEnd.y; y++ {
			for z := z1; z <= z2; z++ {
				space[x][y][z] = i + 1
			}
		}
	}
}

func main() {

	var bricks []Brick
	supportBy, supportList := make(map[int]map[int]struct{}), make(map[int]map[int]struct{})
	space = [11][11][501]int{}
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		start, end, _ := strings.Cut(line, "~")
		b := Brick{strToCoord(start), strToCoord(end)}
		bricks = append(bricks, b)
		supportBy[i+1] = make(map[int]struct{})
		supportList[i+1] = make(map[int]struct{})
	}
	// sort the bricks by z position & height
	sort.SliceStable(bricks, func(i, j int) bool {
		b1, b2 := bricks[i], bricks[j]
		if b1.posStart.z == b1.posEnd.z && b2.posStart.z == b2.posEnd.z {
			return b1.posStart.z < b2.posEnd.z
		} else {
			return min(b1.posStart.z, b1.posEnd.z) < min(b2.posStart.z, b2.posEnd.z)
		}
	})

	// settle the bricks
	for i, brick := range bricks {
		z1, z2 := getZ(brick), getZ(brick)
		if brick.posStart.z != brick.posEnd.z {
			// vertical brick
			z2 = z1 + brick.posEnd.z - brick.posStart.z
		}
		updateSpace(i, z1, z2, brick)
	}

	// part 1
	for z := 1; z < 500; z++ {
		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				if space[x][y][z] != 0 && space[x][y][z+1] != 0 && space[x][y][z+1] != space[x][y][z] {
					// if the brick in space[x][y][z] is different to space[x][y][z+1]
					// this means brick in space[x][y][z] supports space[x][y][z+1]
					supportBy[space[x][y][z+1]][space[x][y][z]] = struct{}{}
					supportList[space[x][y][z]][space[x][y][z+1]] = struct{}{}
				}
			}
		}
	}

	singleSupporters := make(map[int]struct{})

	for _, support := range supportBy {
		if len(support) == 1 {
			// only supported by a single brick, cannot disintegrate
			for val := range support {
				singleSupporters[val] = struct{}{}
			}
		}
	}
	fmt.Println(len(bricks) - len(singleSupporters))
	// part 2
	total := 0
	for brick := range singleSupporters {
		// only consider bricks that will collapse others
		fallen := make(map[int]struct{})
		fallen[brick] = struct{}{}
		queue := make([]int, 1)
		queue[0] = brick

		for len(queue) != 0 {
			newQueue := make([]int, 0)
			for _, current := range queue {
				supportedByCurrent, ok := supportList[current]
				if !ok {
					continue
				}
				for supp := range supportedByCurrent {
					// check each brick supported by supp
					haveAllFallen := true
					for b := range supportBy[supp] {
						if _, ok := fallen[b]; !ok {
							// a different support brick has not fallen
							haveAllFallen = false
						}
					}
					if _, ok := fallen[supp]; haveAllFallen && !ok {
						// all bricks that support supp have fallen, add to the queue
						fallen[supp] = struct{}{}
						newQueue = append(newQueue, supp)
					}
				}
			}
			queue = newQueue
		}
		total += len(fallen) - 1
	}

	fmt.Println(total)
}
