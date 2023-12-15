package main

import (
	_ "embed"
	"fmt"
	"strings"

	"golang.org/x/exp/slices"

	"golang.org/x/exp/maps"

	"github.com/GoosvandenBekerom/advent-of-code/utils"
)

//go:embed input
var input string

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(csv string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	var sum int
	for _, part := range strings.Split(csv, ",") {
		sum += HASH(part)
	}
	return sum
}

func part2(csv string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	type valueAndOrder struct{ value, order int }
	boxes := make(map[int]map[string]valueAndOrder)
	var box int
	var order int
	for _, part := range strings.Split(csv, ",") {
		if strings.ContainsRune(part, '=') {
			label := part[:len(part)-2]
			box = HASH(label)
			if _, exists := boxes[box]; !exists {
				boxes[box] = make(map[string]valueAndOrder)
			}
			nextOrder := order
			if previous, exists := boxes[box][label]; exists {
				nextOrder = previous.order
			}
			boxes[box][label] = valueAndOrder{
				value: utils.ToInt(string(part[len(part)-1])),
				order: nextOrder,
			}
			order++
		}
		if strings.ContainsRune(part, '-') {
			label := part[:len(part)-1]
			box = HASH(label)
			delete(boxes[box], label)
		}
	}

	var sum int
	for i := 0; i < 256; i++ {
		lenses := maps.Values(boxes[i])
		slices.SortFunc(lenses, func(a, b valueAndOrder) int {
			return a.order - b.order
		})
		for lensI, lens := range lenses {
			sum += (i + 1) * (lensI + 1) * lens.value
		}
	}
	return sum
}

func HASH(s string) int {
	var hash int
	for _, char := range s {
		hash += int(char)
		hash *= 17
		hash = hash % 256
	}
	return hash
}
