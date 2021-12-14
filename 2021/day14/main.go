package main

import (
	_ "embed"
	"math"
	"regexp"

	"bytes"

	"fmt"
	"strings"
)

//go:embed input
var input string

func main() {
	part1()
	part2()
}

// ----------------------------------------
// solution
// ----------------------------------------

type pair struct{ left, right byte }
type ruleMap map[pair]byte

func parseInput() ([]byte, ruleMap) {
	parts := strings.Split(normalizeNewlines(input), "\n\n")

	rules := make(ruleMap)
	pattern := regexp.MustCompile(`([A-Z]{2}) -> ([A-Z])`)
	for _, match := range pattern.FindAllStringSubmatch(parts[1], -1) {
		rules[pair{match[1][0], match[1][1]}] = match[2][0]
	}
	return []byte(parts[0]), rules
}

func solve(template []byte, rules ruleMap, iterations int) int {
	pairs := make(map[pair]int)
	for i := 0; i < len(template)-1; i++ {
		pairs[pair{template[i], template[i+1]}]++
	}

	for i := 0; i < iterations; i++ {
		temp := make(map[pair]int)
		for p, num := range pairs {
			next := rules[p]
			temp[pair{p.left, next}] += num
			temp[pair{next, p.right}] += num
		}

		pairs = temp
	}

	counter := make(map[byte]int)
	for p, num := range pairs {
		counter[p.left] += num
	}
	counter[template[len(template)-1]]++

	min := math.MaxInt
	max := 0
	for _, v := range counter {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return max - min
}

func part1() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	template, rules := parseInput()
	println(solve(template, rules, 10))
}

func part2() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	template, rules := parseInput()
	println(solve(template, rules, 40))
}

// ----------------------------------------
// utils
// ----------------------------------------

func normalizeNewlines(s string) string {
	d := []byte(s)
	// replace CR LF \r\n (windows) with LF \n (unix)
	d = bytes.Replace(d, []byte{13, 10}, []byte{10}, -1)
	// replace CF \r (mac) with LF \n (unix)
	d = bytes.Replace(d, []byte{13}, []byte{10}, -1)
	return string(d)
}
