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

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIRS
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

const JOKER = 'J'

type hand struct {
	cards []rune
	bid   int
}

func (c hand) compare(cardStrength map[rune]int, other hand) bool {
	for i := 0; i < 5; i++ {
		if cardStrength[c.cards[i]] < cardStrength[other.cards[i]] {
			return true
		} else if cardStrength[c.cards[i]] > cardStrength[other.cards[i]] {
			return false
		}
	}
	return false
}

// return strength of a player's hand,
func (c hand) strength(withJoker bool) int {
	cardCount := make(map[rune]int)
	for _, c := range c.cards {
		cardCount[c] += 1
	}
	//with joker, always better to add them to the highest number of cards
	if withJoker {
		max := 0
		var maxCard rune
		for k, v := range cardCount {
			if k != JOKER && v > max {
				max = v
				maxCard = k
			}
		}
		if max > 0 {
			cardCount[maxCard] += cardCount[JOKER]
			delete(cardCount, JOKER)
		}
	}
	productCount := 1
	for _, v := range cardCount {
		productCount = productCount * v
	}
	switch productCount {
	case 5:
		return FIVE_OF_A_KIND
	case 4:
		switch len(cardCount) {
		case 2:
			return FOUR_OF_A_KIND
		case 3:
			return TWO_PAIRS
		}
	case 6:
		return FULL_HOUSE
	case 3:
		return THREE_OF_A_KIND
	case 2:
		return ONE_PAIR
	case 1:
		return HIGH_CARD
	}
	return 0
}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	splits := strings.Split(input, "\n")
	var hands []hand
	for _, split := range splits {
		left, right, _ := strings.Cut(split, " ")
		bid, _ := strconv.Atoi(right)
		cards := []rune(left)
		hands = append(hands, hand{
			cards: cards,
			bid:   bid,
		})
	}

	//part 1
	sum := 0
	cardStrength := map[rune]int{
		'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10, '9': 9,
		'8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2,
	}
	sort.Slice(hands, func(i, j int) bool {
		iStrength := hands[i].strength(false)
		jStrength := hands[j].strength(false)
		if iStrength != jStrength {
			return iStrength < jStrength
		}
		return hands[i].compare(cardStrength, hands[j])
	})
	for i, card := range hands {
		sum += (i + 1) * card.bid
	}
	fmt.Println("Total sum is", sum)

	//part 2
	sum = 0
	//make joker the weakest card
	cardStrength[JOKER] = 1
	sort.Slice(hands, func(i, j int) bool {
		iStrength := hands[i].strength(true)
		jStrength := hands[j].strength(true)
		if iStrength != jStrength {
			return iStrength < jStrength
		}
		return hands[i].compare(cardStrength, hands[j])
	})
	for i, card := range hands {
		sum += (i + 1) * card.bid
	}
	fmt.Println("Total sum with joker is", sum)
}
