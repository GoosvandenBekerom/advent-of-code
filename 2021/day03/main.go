package main

import (
	_ "embed"
	"strconv"

	"fmt"
	"log"
	"strings"
)

//go:embed input
var input string

func main() {
	part1()
	part2()
}

/*
--- Day 3: Binary Diagnostic ---
The submarine has been making some odd creaking noises, so you ask it to produce a diagnostic report just in case.

The diagnostic report (your puzzle input) consists of a list of binary numbers which, when decoded properly, can tell you many useful things about the conditions of the submarine. The first parameter to check is the power consumption.

You need to use the binary numbers in the diagnostic report to generate two new binary numbers (called the gamma rate and the epsilon rate). The power consumption can then be found by multiplying the gamma rate by the epsilon rate.

Each bit in the gamma rate can be determined by finding the most common bit in the corresponding position of all numbers in the diagnostic report. For example, given the following diagnostic report:

00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
Considering only the first bit of each number, there are five 0 bits and seven 1 bits. Since the most common bit is 1, the first bit of the gamma rate is 1.

The most common second bit of the numbers in the diagnostic report is 0, so the second bit of the gamma rate is 0.

The most common value of the third, fourth, and fifth bits are 1, 1, and 0, respectively, and so the final three bits of the gamma rate are 110.

So, the gamma rate is the binary number 10110, or 22 in decimal.

The epsilon rate is calculated in a similar way; rather than use the most common bit, the least common bit from each position is used. So, the epsilon rate is 01001, or 9 in decimal. Multiplying the gamma rate (22) by the epsilon rate (9) produces the power consumption, 198.

Use the binary numbers in your diagnostic report to calculate the gamma rate and epsilon rate, then multiply them together. What is the power consumption of the submarine? (Be sure to represent your answer in decimal, not binary.)
*/

// ----------------------------------------
// solution
// ----------------------------------------

func part1() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	lines := strings.Split(input, "\n")
	length := len(lines[0])

	balances := make([]int, length)
	for _, line := range lines {
		for i, c := range line {
			if c == '0' {
				balances[i]--
			} else {
				balances[i]++
			}
		}
	}

	var gamma, epsilon string
	for _, balance := range balances {
		if balance < 0 {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	g, err := strconv.ParseUint(gamma, 2, 64)
	check(err)

	e, err := strconv.ParseUint(epsilon, 2, 64)
	check(err)

	fmt.Printf("gamma: %d\n", g)
	fmt.Printf("epsilon: %d\n", e)
	fmt.Printf("solution: %d\n", g*e)
}

func part2() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	lines := strings.Split(input, "\n")
	oxygen := getCommon(lines, true)
	co2 := getCommon(lines, false)

	o, err := strconv.ParseUint(oxygen, 2, 64)
	check(err)

	c, err := strconv.ParseUint(co2, 2, 64)
	check(err)

	fmt.Printf("oxygen: %d\n", o)
	fmt.Printf("co2: %d\n", c)
	fmt.Printf("solution: %d\n", o*c)
}

func getCommon(bitstrings []string, most bool) string {
	var result string
	for i := 0; i < len(bitstrings[0]); i++ {
		var remaining []string
		for _, line := range bitstrings {
			if strings.HasPrefix(line, result) {
				remaining = append(remaining, line)
			}
		}

		if len(remaining) == 1 {
			return remaining[0]
		}

		var balance int
		for _, line := range remaining {
			if line[i] == '0' {
				balance--
			} else {
				balance++
			}
		}

		if most {
			if balance < 0 {
				result += "0"
			} else {
				result += "1"
			}
		} else {
			if balance < 0 {
				result += "1"
			} else {
				result += "0"
			}
		}
	}
	return result
}

// ----------------------------------------
// utils
// ----------------------------------------

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
