package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"

	"github.com/GoosvandenBekerom/advent-of-code/utils"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(parse(lines)))
	fmt.Println(part2(parse(lines)))
}

type ingredientRange struct {
	lower, upper int
}
type ingredients struct {
	ranges    []ingredientRange
	available []int
}

func parse(lines []string) ingredients {
	parsingRanges := true
	var result ingredients
	for _, line := range lines {
		if line == "" {
			parsingRanges = false
			continue
		}
		if parsingRanges {
			left, right, _ := strings.Cut(line, "-")
			result.ranges = append(result.ranges, ingredientRange{utils.ToInt(left), utils.ToInt(right)})
		} else {
			result.available = append(result.available, utils.ToInt(line))
		}
	}
	return result
}

func part1(ingredients ingredients) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	fresh := 0
	for _, id := range ingredients.available {
		for _, r := range ingredients.ranges {
			if id >= r.lower && id <= r.upper {
				fresh++
				break
			}
		}
	}

	return fresh
}

func part2(ingredients ingredients) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	sort.Slice(ingredients.ranges, func(i, j int) bool {
		return ingredients.ranges[i].lower < ingredients.ranges[j].lower
	})

	answer := 0
	highestCovered := 0
	for _, r := range ingredients.ranges {
		lower := highestCovered + 1
		if r.lower > highestCovered {
			lower = r.lower
		}
		if r.upper < lower {
			continue
		}
		answer += r.upper - lower + 1
		highestCovered = r.upper
	}

	return answer
}
