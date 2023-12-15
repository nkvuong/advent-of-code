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

func isPattern(pattern []string, num int) bool {
	curr := num + 1
	reflection := num
	for {
		if curr >= len(pattern) {
			return true
		}
		if reflection < 0 {
			return true
		}
		if pattern[curr] != pattern[reflection] {
			return false
		}
		curr++
		reflection--
	}
}

func findReflection(pattern string, ignoreReflection int) int {
	// check for horizontal reflection line
	horizontal := strings.Split(pattern, "\n")
	for i := 0; i < len(horizontal)-1; i++ {
		if isPattern(horizontal, i) && (i+1)*100 != ignoreReflection {
			return (i + 1) * 100
		}
	}
	// construct the vertical lines
	var vertical []string
	for i := 0; i < len(horizontal[0]); i++ {
		var sb strings.Builder
		for j := 0; j < len(horizontal); j++ {
			sb.Write([]byte{horizontal[j][i]})
		}
		vertical = append(vertical, sb.String())
	}
	// check for vertical reflection line
	for i := 0; i < len(vertical)-1; i++ {
		if isPattern(vertical, i) && (i+1) != ignoreReflection {
			return i + 1
		}
	}
	return -1
}

func main() {
	patterns := strings.Split(input, "\n\n")
	reflectionSum := 0
	for _, pattern := range patterns {
		reflectionSum += findReflection(pattern, 0)
	}
	fmt.Println("Reflection sum is", reflectionSum)
	reflectionSum = 0
	for num, pattern := range patterns {
		oldReflection := findReflection(pattern, 0)
		for i := 0; i < len(pattern); i++ {
			if pattern[i] == '\n' {
				continue
			}
			replacement := "."
			if pattern[i] == '.' {
				replacement = "#"
			}
			newPattern := pattern[:i] + replacement + pattern[i+1:]
			newReflection := findReflection(newPattern, oldReflection)
			if newReflection != -1 {
				fmt.Println(num, newReflection)
				reflectionSum += newReflection
				break
			}
		}
	}
	fmt.Println("Reflection sum with smudge is", reflectionSum)
}
