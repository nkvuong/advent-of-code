package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

type lens struct {
	label    string
	strength int
}

func hash(input string) (output int) {
	for _, r := range input {
		output += int(r)
		output *= 17
		output = output % 256
	}
	return output
}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	sum := 0
	for _, seq := range strings.Split(input, ",") {
		sum += hash(seq)
	}
	fmt.Println("Total hash is ", sum)

	boxes := make(map[int][]*lens)
	for _, seq := range strings.Split(input, ",") {
		if strings.Contains(seq, "-") {
			for num, box := range boxes {
				for i := 0; i < len(box); i++ {
					if box[i].label == strings.TrimSuffix(seq, "-") {
						newBox := append(box[:i], box[i+1:]...)
						boxes[num] = newBox
						break
					}
				}
			}
		}
		if strings.Contains(seq, "=") {
			found := false
			label, s, _ := strings.Cut(seq, "=")
			strength, _ := strconv.Atoi(s)
			l := lens{label, strength}
			num := hash(l.label)
			for i := 0; i < len(boxes[num]); i++ {
				if boxes[num][i].label == l.label {
					copy(boxes[num][i:], []*lens{&l})
					found = true
				}
			}
			if !found {
				boxes[num] = append(boxes[num], &l)
			}
		}
	}
	sum = 0
	for num, lens := range boxes {
		for i := 0; i < len(lens); i++ {
			sum += (num + 1) * (i + 1) * lens[i].strength
		}
	}
	fmt.Println("Total focus power is", sum)
}
