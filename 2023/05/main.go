package main

import (
	_ "embed"
	"fmt"
	"math"
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

type converter struct {
	sourceStart      int
	sourceEnd        int
	destinationStart int
}

type seedRange struct {
	start int
	end   int
}

type plantMap struct {
	converters []converter
}

func (c converter) getDestination(source int) int {
	return source - c.sourceStart + c.destinationStart
}

func newMap(input string) plantMap {
	splits := strings.Split(input, "\n")
	var converters []converter
	for _, split := range splits[1:] {
		r := regexp.MustCompile(`\d+`)
		nums := r.FindAllString(split, -1)
		destinationStart, _ := strconv.Atoi(nums[0])
		sourceStart, _ := strconv.Atoi(nums[1])
		length, _ := strconv.Atoi(nums[2])
		converters = append(converters, converter{
			sourceStart:      sourceStart,
			destinationStart: destinationStart,
			sourceEnd:        sourceStart + length,
		})
	}
	return plantMap{converters: converters}
}

func (m plantMap) transform(source int) int {
	for _, convert := range m.converters {
		if source >= convert.sourceStart && source < convert.sourceEnd {
			return convert.getDestination(source)
		}
	}
	return source
}

func (m plantMap) transformRange(source []seedRange) []seedRange {
	var destination []seedRange
	for _, converter := range m.converters {
		var unchanged []seedRange
		for _, seed := range source {
			//ignore invalid range
			if seed.end > seed.start {
				//overlapped range goes to destination
				//we will ignore invalid range in our final calculation
				destination = append(destination, seedRange{
					start: converter.getDestination(max(seed.start, converter.sourceStart)),
					end:   converter.getDestination(min(seed.end, converter.sourceEnd)),
				})
				//append left & right nonoverlap
				unchanged = append(unchanged, seedRange{
					start: max(converter.sourceEnd, seed.start),
					end:   seed.end,
				}, seedRange{
					start: seed.start,
					end:   min(converter.sourceStart, seed.end),
				})
			}
		}
		source = unchanged
	}
	return append(destination, source...)
}

func main() {

	splits := strings.Split(input, "\n\n")
	seedsString := splits[0]
	r := regexp.MustCompile(`\d+`)
	seedStrings := r.FindAllString(seedsString, -1)
	var maps []plantMap
	for _, s := range splits[1:] {
		maps = append(maps, newMap(s))
	}
	var seeds []int
	for _, seed := range seedStrings {
		num, _ := strconv.Atoi(seed)
		seeds = append(seeds, num)
	}
	//part 1
	min := math.MaxInt64
	for _, seed := range seeds {
		for _, m := range maps {
			seed = m.transform(seed)
		}
		if seed < min {
			min = seed
		}
	}
	fmt.Println("Lowest location number is", min)

	//part 2
	min = math.MaxInt64
	for i := 0; i < len(seeds); i += 2 {
		destination := []seedRange{{start: seeds[i], end: seeds[i] + seeds[i+1]}}
		for _, m := range maps {
			destination = m.transformRange(destination)
		}
		for _, r := range destination {
			if r.start < min && r.end > r.start {
				min = r.start
			}
		}
	}
	fmt.Println("Lowest location number is", min)
}
