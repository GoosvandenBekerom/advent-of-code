package main

import (
	_ "embed"
	"fmt"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
	"regexp"
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

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	re := regexp.MustCompile("[0-9]+")
	var sum int
	for _, line := range lines {
		digits := re.FindAllString(line, -1)
		lastDigit := digits[len(digits)-1]
		answer, _ := strconv.Atoi(string(digits[0][0]) + string(lastDigit[len(lastDigit)-1]))
		sum += answer
	}

	return sum
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	re1 := regexp.MustCompile("([0-9]+|one|two|three|four|five|six|seven|eight|nine)")
	re2 := regexp.MustCompile("([0-9]+|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)")
	var sum int
	for _, line := range lines {
		first := mapDigit(re1.FindString(line))
		second := mapDigit(re2.FindString(utils.Reverse(line)))
		sum += utils.ToInt(first + second)
	}

	return sum
}

func mapDigit(digit string) string {
	switch digit {
	case "one", "eno":
		return "1"
	case "two", "owt":
		return "2"
	case "three", "eerht":
		return "3"
	case "four", "ruof":
		return "4"
	case "five", "evif":
		return "5"
	case "six", "xis":
		return "6"
	case "seven", "neves":
		return "7"
	case "eight", "thgie":
		return "8"
	case "nine", "enin":
		return "9"
	default:
		return string(digit[0])
	}
}
