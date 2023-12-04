package main

import (
	_ "embed"
	"fmt"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
	"strings"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

type number struct {
	startIndex int
	length     int
}

type pos struct {
	line, index int
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	numbersByLine := make([][]number, len(lines))
	symbols := make(map[pos]int32)
	for lineNumber, line := range lines {
		numbersByLine[lineNumber] = make([]number, 0)
		var currentNumber *number
		for i, c := range line {
			if c >= '0' && c <= '9' {
				if currentNumber == nil {
					currentNumber = &number{startIndex: i, length: 1}
				} else {
					currentNumber.length++
				}
			} else {
				if currentNumber != nil {
					numbersByLine[lineNumber] = append(numbersByLine[lineNumber], *currentNumber)
					currentNumber = nil
				}

				if c == '.' {
					continue
				}

				if strings.ContainsAny(string(c), "=-*$%/+&#@") {
					symbols[pos{line: lineNumber, index: i}] = c
				}
			}
		}
		if currentNumber != nil {
			numbersByLine[lineNumber] = append(numbersByLine[lineNumber], *currentNumber)
			currentNumber = nil
		}
	}

	var sum int

	for lineNumber, numbers := range numbersByLine {
		for _, n := range numbers {
			for column := n.startIndex - 1; column < n.startIndex+n.length+1; column++ {
				if containsAny(symbols,
					pos{lineNumber - 1, column}, // line above
					pos{lineNumber, column},     // line itself
					pos{lineNumber + 1, column}, // line below
				) {
					sum += utils.ToInt(lines[lineNumber][n.startIndex : n.startIndex+n.length])
				}
			}
		}
	}

	return sum
}

func containsAny(symbols map[pos]int32, positions ...pos) bool {
	for _, position := range positions {
		if _, found := symbols[position]; found {
			return true
		}
	}

	return false
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	numberPositions := make(map[pos]*number)
	potentialGears := make(map[pos]struct{})
	for lineNumber, line := range lines {
		var currentNumber *number
		for i, c := range line {
			if c >= '0' && c <= '9' {
				if currentNumber == nil {
					currentNumber = &number{startIndex: i, length: 1}
				} else {
					currentNumber.length++
				}
				numberPositions[pos{lineNumber, i}] = currentNumber
			} else {
				if currentNumber != nil {
					currentNumber = nil
				}

				if c == '*' {
					potentialGears[pos{line: lineNumber, index: i}] = struct{}{}
				}
			}
		}
		if currentNumber != nil {
			currentNumber = nil
		}
	}

	var sum int

	for p := range potentialGears {
		matchingNumbers := make(map[*number]int)
		// line above
		if n, ok := numberPositions[pos{p.line - 1, p.index - 1}]; ok {
			matchingNumbers[n] = p.line - 1
		}
		if n, ok := numberPositions[pos{p.line - 1, p.index}]; ok {
			matchingNumbers[n] = p.line - 1
		}
		if n, ok := numberPositions[pos{p.line - 1, p.index + 1}]; ok {
			matchingNumbers[n] = p.line - 1
		}
		// line itself
		if n, ok := numberPositions[pos{p.line, p.index - 1}]; ok {
			matchingNumbers[n] = p.line
		}
		if n, ok := numberPositions[pos{p.line, p.index}]; ok {
			matchingNumbers[n] = p.line
		}
		if n, ok := numberPositions[pos{p.line, p.index + 1}]; ok {
			matchingNumbers[n] = p.line
		}
		// line below
		if n, ok := numberPositions[pos{p.line + 1, p.index - 1}]; ok {
			matchingNumbers[n] = p.line + 1
		}
		if n, ok := numberPositions[pos{p.line + 1, p.index}]; ok {
			matchingNumbers[n] = p.line + 1
		}
		if n, ok := numberPositions[pos{p.line + 1, p.index + 1}]; ok {
			matchingNumbers[n] = p.line + 1
		}

		if len(matchingNumbers) == 2 {
			gearRatio := 1
			for n, lineNumber := range matchingNumbers {
				gearRatio *= utils.ToInt(lines[lineNumber][n.startIndex : n.startIndex+n.length])
			}
			sum += gearRatio
		}
	}

	return sum
}
