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

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func stringToIntSlice(in string) (out []int) {
	r := regexp.MustCompile(`\d+`)
	matches := r.FindAllString(in, -1)
	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		out = append(out, num)
	}
	return out
}

func getArrangements(springMap string, countDamaged []int) int {
	arrangements := 0
	// state is a tuple of 4 values
	currentStates := map[[4]int]int{{0, 0, 0, 0}: 1}
	newStates := map[[4]int]int{}
	for len(currentStates) > 0 {
		for state, num := range currentStates {
			mapCount, damagedCount, damagedNum, expectWorking := state[0], state[1], state[2], state[3]
			if mapCount == len(springMap) {
				// if we are at the end of the map
				if damagedCount == len(countDamaged) {
					arrangements += num
				}
				continue
			}
			switch {
			case (springMap[mapCount] == '#' || springMap[mapCount] == '?') && damagedCount < len(countDamaged) && expectWorking == 0:
				// we are still looking for damaged springs
				if springMap[mapCount] == '?' && damagedNum == 0 {
					// we are not in a run of damaged springs, so ? can be working
					newStates[[4]int{mapCount + 1, damagedCount, damagedNum, expectWorking}] += num
				}
				damagedNum++
				if damagedNum == countDamaged[damagedCount] {
					// we've found the full section of damaged springs
					damagedCount++
					damagedNum = 0
					expectWorking = 1
					// we only want a working spring next
				}
				newStates[[4]int{mapCount + 1, damagedCount, damagedNum, expectWorking}] += num
			case (springMap[mapCount] == '.' || springMap[mapCount] == '?') && damagedNum == 0:
				// we are not in a section of damaged springs
				expectWorking = 0
				newStates[[4]int{mapCount + 1, damagedCount, damagedNum, expectWorking}] += num
			}
		}
		currentStates, newStates = newStates, currentStates
		newStates = map[[4]int]int{}
	}
	return arrangements
}

func main() {
	springs := strings.Split(input, "\n")
	sumArrangements := 0
	sumArrangementsUnfolded := 0
	for _, spring := range springs {
		springMap, countString, _ := strings.Cut(spring, " ")
		var springMapUnfolded, countStringUnfolded string
		for i := 0; i < 5; i++ {
			springMapUnfolded, countStringUnfolded = springMapUnfolded+springMap+"?", countStringUnfolded+countString+","
		}
		springMapUnfolded = strings.TrimSuffix(springMapUnfolded, "?")
		countDamanged := stringToIntSlice(countString)
		countDamangedUnfolded := stringToIntSlice(countStringUnfolded)
		sumArrangements += getArrangements(springMap, countDamanged)
		sumArrangementsUnfolded += getArrangements(springMapUnfolded, countDamangedUnfolded)
	}
	fmt.Println("Total sum of arrangements is", sumArrangements)
	fmt.Println("Total sum of arrangements unfolded is", sumArrangementsUnfolded)
}
