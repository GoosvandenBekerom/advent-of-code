package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

const magicColor = "shiny gold"

func main() {
	fmt.Printf("part 1: %d\n", part1())
	fmt.Printf("part 2: %d\n", part2())
}

func part1() (count int) {
	winners := make(map[string]bool)
	winners[magicColor] = true

	bags := input()
	delete(bags, magicColor)
	for {
		remainder, matches := process(bags, keys(winners)...)
		if len(matches) == 0 {
			break
		}
		for _, m := range matches {
			winners[m] = true
		}
		bags = remainder
	}

	delete(winners, magicColor)
	return len(winners)
}

func process(bags bags, colors ...string) (remainder bags, matches []string) {
	remainder = make(map[string]map[string]int)
	for bag, fits := range bags {
		matched := false
		for _, color := range colors {
			if fits[color] > 0 {
				matches = append(matches, bag)
				matched = true
				break
			}
		}
		if !matched {
			remainder[bag] = fits
		}
	}
	return
}

func part2() int {
	bags := input()
	return count(bags, magicColor)
}

func count(bags bags, color string) (sum int) {
	for bag, amount := range bags[color] {
		sum += (count(bags, bag) + 1) * amount
	}
	return
}

// ----------------------------------------
// utils
// ----------------------------------------

// bag has a color as its key and its value is a map of the bags that fit inside it
// with their color as their key and the amount as their value
type bags map[string]map[string]int

func input() bags {
	box := packr.New("day07", "./2020/day07")
	s, err := box.FindString("input")
	check(err)
	bags := make(map[string]map[string]int)
	for _, in := range strings.Split(strings.TrimSuffix(s, "\n"), "\n") {
		raw := strings.Split(in, " bags contain ")
		pattern := regexp.MustCompile(`(\d+)\s([\w ]+)\s`)
		matches := pattern.FindAllStringSubmatch(raw[1], -1)

		fits := make(map[string]int)
		for i := range matches {
			n, err := strconv.Atoi(matches[i][1])
			check(err)
			fits[matches[i][2]] = n
		}
		bags[raw[0]] = fits
	}
	return bags
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func keys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
