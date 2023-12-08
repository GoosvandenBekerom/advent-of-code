package main

import (
	_ "embed"
	"fmt"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
	"golang.org/x/exp/slices"
	"strings"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	return play(lines, false, map[int32]int{
		'A': 1,
		'K': 2,
		'Q': 3,
		'J': 4,
		'T': 5,
		'9': 6,
		'8': 7,
		'7': 8,
		'6': 9,
		'5': 10,
		'4': 11,
		'3': 12,
		'2': 13,
	})
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	return play(lines, true, map[int32]int{
		'A': 1,
		'K': 2,
		'Q': 3,
		'T': 5,
		'9': 6,
		'8': 7,
		'7': 8,
		'6': 9,
		'5': 10,
		'4': 11,
		'3': 12,
		'2': 13,
		'J': 14,
	})
}

type handType int

const (
	fiveOfAKind handType = iota
	fourOfAKind
	fullHouse
	threeOfAKind
	twoPair
	onePair
	highCard
)

func getHandType(hand string, jokersEnabled bool) handType {
	unique := make(map[int32]int)
	for _, card := range hand {
		unique[card]++
	}
	counts := make(map[int]int32)
	for card, count := range unique {
		counts[count] = card
	}

	amountOfJokers := 0
	if amount, hasJokers := unique['J']; jokersEnabled && hasJokers {
		amountOfJokers = amount
	}

	if _, has5 := counts[5]; has5 {
		return fiveOfAKind
	}

	if _, has4 := counts[4]; has4 {
		if amountOfJokers == 4 || amountOfJokers == 1 {
			return fiveOfAKind
		}
		return fourOfAKind
	}

	if _, has3 := counts[3]; has3 {
		if amountOfJokers == 2 {
			return fiveOfAKind
		}
		if amountOfJokers == 3 {
			if len(unique) == 2 {
				return fiveOfAKind
			}
			if len(unique) == 3 {
				return fourOfAKind
			}
		}
		if amountOfJokers == 1 {
			return fourOfAKind
		}
		if _, has2 := counts[2]; has2 {
			return fullHouse
		}
		return threeOfAKind
	}

	if _, has2 := counts[2]; has2 {
		if amountOfJokers == 3 {
			return fiveOfAKind
		}
		if amountOfJokers == 2 {
			if len(unique) == 3 {
				return fourOfAKind
			}
			return threeOfAKind
		}
		if amountOfJokers == 1 {
			if len(unique) == 3 {
				return fullHouse
			}
			return threeOfAKind
		}
		if len(unique) == 3 {
			return twoPair
		}
		return onePair
	}

	if amountOfJokers == 1 {
		return onePair
	}

	return highCard
}

func getSortHandsFunc(cardOrder map[int32]int) func(a hand, b hand) int {
	return func(a, b hand) int {
		diff := int(b.Type - a.Type)
		if diff != 0 {
			return diff
		}

		for i, card := range a.Cards {
			orderA, orderB := cardOrder[card], cardOrder[rune(b.Cards[i])]
			if orderA != orderB {
				return orderB - orderA
			}
		}

		return 0
	}
}

type hand struct {
	Cards string
	Type  handType
	Bet   int
}

func play(lines []string, jokers bool, cardOrder map[int32]int) int {
	var hands []hand

	for _, line := range lines {
		cards, bet, _ := strings.Cut(line, " ")
		hands = append(hands, hand{Cards: cards, Type: getHandType(cards, jokers), Bet: utils.ToInt(bet)})
	}

	slices.SortFunc(hands, getSortHandsFunc(cardOrder))

	var sum int
	for i, hand := range hands {
		//println(hand.Cards, hand.Bet, hand.Type)
		sum += hand.Bet * (i + 1)
	}

	return sum
}
