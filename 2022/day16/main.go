package main

import (
	_ "embed"
	"fmt"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
	"regexp"
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

	valves := parseInput(lines)
	current := "AA"
	pressureRelease := 0
	for remainingMinutes := 30; remainingMinutes > 0; remainingMinutes-- {
		printState(current, remainingMinutes, valves, pressureRelease)
		v := valves[current]
		if v.flowRate > 0 && !v.open {
			pressureRelease += remainingMinutes * v.flowRate
			v.open = true
			continue
		}
		highestCode, highestFlow := "", 0
		for _, tunnel := range v.tunnels {
			next := valves[tunnel]
			if !next.open && next.flowRate > highestFlow {
				highestFlow = next.flowRate
				highestCode = tunnel
			}
		}
		// todo highestcode can be empty, I think i need a graph to go deeper or something
		current = highestCode
	}
	return pressureRelease
}

func printState(current string, remaining int, valves valveMap, release int) {
	var opened []string
	for code, v := range valves {
		if v.open {
			opened = append(opened, code)
		}
	}
	fmt.Printf("position %s -> %d minutes remaining\nvalves open: %v\nreleasing %d\n\n", current, remaining, opened, release)
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	return -1
}

type valveMap map[string]*valve
type valve struct {
	flowRate int
	tunnels  []string
	open     bool
}

var regex = regexp.MustCompile(`Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z]{2}(, [A-Z]{2})*)`)

func parseInput(lines []string) valveMap {
	valves := make(valveMap)
	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)
		valves[matches[1]] = &valve{
			flowRate: utils.ToInt(matches[2]),
			tunnels:  strings.Split(matches[3], ", "),
		}
	}
	return valves
}
