package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/GoosvandenBekerom/advent-of-code/utils"
)

//go:embed input
var input string

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(line string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	answer := 0
	for _, r := range strings.Split(line, ",") {
		left, right, _ := strings.Cut(r, "-")
		lower := utils.ToInt(left)
		upper := utils.ToInt(right)
		for lower <= upper {
			number := strconv.Itoa(lower)
			if len(number)%2 == 0 && number[:len(number)/2] == number[len(number)/2:] {
				answer += lower
			}
			lower++
		}
	}

	return answer
}

func part2(line string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	answer := 0
	for _, r := range strings.Split(line, ",") {
		left, right, _ := strings.Cut(r, "-")
		lower := utils.ToInt(left)
		upper := utils.ToInt(right)
		for lower <= upper {
			if isRepeating(strconv.Itoa(lower)) {
				answer += lower
			}
			lower++
		}
	}

	return answer
}

func isRepeating(s string) bool {
	if len(s) == 2 {
		return s[0] == s[1]
	}

	index := strings.Index((s + s)[1:], s)
	return index != -1 && index < len(s)-1
}
