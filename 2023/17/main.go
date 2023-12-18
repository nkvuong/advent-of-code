package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

type coord struct {
	row int
	col int
}

type state struct {
	pos coord
	dir coord
}

type city struct {
	state state
	heat  int
}

type heapQueue []city

func (q heapQueue) Len() int           { return len(q) }
func (q heapQueue) Less(i, j int) bool { return q[i].heat < q[j].heat }
func (q heapQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *heapQueue) Push(x any)        { *q = append(*q, x.(city)) }
func (q *heapQueue) Pop() (x any)      { x, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]; return x }

var grid map[coord]int
var end coord

func run(min, max int) int {
	queue, seen := heapQueue{}, map[state]bool{}
	heap.Push(&queue, city{state{coord{0, 0}, coord{1, 0}}, 0})
	heap.Push(&queue, city{state{coord{0, 0}, coord{0, 1}}, 0})

	for {
		//get the last visited location with lowest heat
		item := heap.Pop(&queue).(city)
		curr, heat := item.state, item.heat

		if curr.pos == end {
			//we've reached the end
			return heat
		}
		if _, ok := seen[curr]; ok {
			//already visited
			continue
		}
		seen[curr] = true

		for _, d := range []coord{
			{curr.dir.col, curr.dir.row},
			{-curr.dir.col, -curr.dir.row},
		} {
			for i := min; i <= max; i++ {
				// move forward up to 3 steps, then turn right and left
				n := coord{curr.pos.row + d.row*i, curr.pos.col + d.col*i}
				if _, ok := grid[n]; ok {
					h := 0
					for j := 1; j <= i; j++ {
						h += grid[coord{curr.pos.row + d.row*j, curr.pos.col + d.col*j}]
					}
					heap.Push(&queue, city{state{n, d}, heat + h})
				}
			}
		}
	}
}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {

	lines := strings.Split(input, "\n")
	grid = make(map[coord]int)
	for row, line := range lines {
		for col, r := range line {
			grid[coord{row, col}] = int(r - '0')
		}
	}
	end = coord{len(lines) - 1, len(lines[0]) - 1}

	fmt.Println("Minimal heat loss is", run(1, 3))
	fmt.Println("Minimal heat loss for ultra crucible is", run(4, 10))
}
