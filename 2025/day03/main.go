package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

var num = map[byte]int{
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func getHighest(values string) (value string, highestIdx int) {
	highest := 0
	for i, c := range []byte(values) {
		if num[c] > highest {
			highest = num[c]
			highestIdx = i
		}
	}

	return strconv.Itoa(highest), highestIdx
}

func getHighestNumberIn(value string, length int) int {
	var answer string
	var fromIndex = 0
	for i := 1; i <= length; i++ {
		highest, highestIdx := getHighest(value[fromIndex : len(value)-(length-i)])
		fromIndex += highestIdx + 1
		answer += highest
	}

	val, _ := strconv.Atoi(answer)
	return val
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	answer := 0
	for _, line := range lines {
		number := getHighestNumberIn(line, 2)
		answer += number
	}

	return answer
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	answer := 0
	for _, line := range lines {
		number := getHighestNumberIn(line, 12)
		answer += number
	}

	return answer
}
