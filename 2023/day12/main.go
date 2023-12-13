package main

import (
	"bytes"
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

type configuration struct {
	springs string
	groups  string
}

func (p configuration) Unfold() configuration {
	return configuration{
		springs: strings.Join([]string{p.springs, p.springs, p.springs, p.springs, p.springs}, "?"),
		groups:  strings.Join([]string{p.groups, p.groups, p.groups, p.groups, p.groups}, ","),
	}
}

func parse(lines []string) (c []configuration) {
	for _, line := range lines {
		parts := strings.Fields(line)
		c = append(c, configuration{
			springs: parts[0],
			groups:  parts[1],
		})
	}
	return c
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	var sum int
	for _, conf := range parse(lines) {
		sum += findArrangements([]byte(conf.springs+"."), conf.groups)
	}
	return sum
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	var sum int
	for _, conf := range parse(lines) {
		conf = conf.Unfold()
		sum += findArrangements([]byte(conf.springs+"."), conf.groups)
	}
	return sum
}

var solved = make(map[string]int, 4096)

func findArrangements(springs []byte, groups string) int {
	key := fmt.Sprintf("%s-%s", springs, groups)
	if solution, ok := solved[key]; ok {
		return solution
	}
	actual := solveRecursive(springs, groups)
	solved[key] = actual
	return actual
}

func solveRecursive(line []byte, groups string) int {
	if len(line) == 0 {
		if len(groups) == 0 {
			return 1
		}
		return 0
	}

	if len(groups) == 0 {
		// ensure all remaining records are not '#'
		if bytes.ContainsRune(line, '#') {
			return 0
		}
		return 1
	}

	if line[0] == '.' {
		return findArrangements(line[1:], groups)
	}

	// Try to consume the next group of #'s
	if line[0] == '#' {
		groupsSplit := strings.SplitN(groups, ",", 2)
		group, _ := strconv.Atoi(groupsSplit[0])
		for i := 0; i < group; i++ {
			if line[i] == '.' {
				return 0
			}
		}
		if line[group] == '#' {
			return 0
		}

		if len(groupsSplit) == 1 {
			// This ensures that the final group is an empty string
			groupsSplit = append(groupsSplit, "")
		}

		return findArrangements(line[group+1:], groupsSplit[1])
	}

	// Next char is '?'. Try both '.' and '#' for matches.
	return findArrangements(append([]byte{'#'}, line[1:]...), groups) +
		findArrangements(append([]byte{'.'}, line[1:]...), groups)
}
