package main

import (
	_ "embed"
	"fmt"
	"strings"
	"sync/atomic"
	"unicode"
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

type cave struct {
	name  string
	isBig bool
}

func part1() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	destinations := make(map[string][]*cave)

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "-")
		left := &cave{name: parts[0], isBig: isUppercase(parts[0])}
		right := &cave{name: parts[1], isBig: isUppercase(parts[1])}

		destinations[left.name] = append(destinations[left.name], right)
		if left.name == "start" {
			continue
		}
		destinations[right.name] = append(destinations[right.name], left)
	}

	delete(destinations, "end")

	visited := make(map[string]struct{})
	var paths int32
	findPaths("start", destinations, visited, &paths)
	fmt.Printf("solution: %d\n", paths)
}

func findPaths(cave string, destinations map[string][]*cave, visited map[string]struct{}, paths *int32) {
	visited[cave] = struct{}{}

	if cave == "end" {
		atomic.AddInt32(paths, 1)
		return
	}

	for _, destination := range destinations[cave] {
		_, found := visited[destination.name]
		if found && !destination.isBig {
			continue
		}
		findPaths(destination.name, destinations, copyMap(visited), paths)
	}
}

func part2() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	destinations := make(map[string][]*cave)

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "-")
		left := &cave{name: parts[0], isBig: isUppercase(parts[0])}
		right := &cave{name: parts[1], isBig: isUppercase(parts[1])}

		destinations[left.name] = append(destinations[left.name], right)
		destinations[right.name] = append(destinations[right.name], left)
	}

	delete(destinations, "end")

	var paths int32
	findPaths2("start", destinations, make(map[string]struct{}), &paths, false)
	fmt.Printf("solution: %d\n", paths)
}

func findPaths2(cave string, destinations map[string][]*cave, visited map[string]struct{}, paths *int32, twice bool) {
	visited[cave] = struct{}{}

	if cave == "end" {
		atomic.AddInt32(paths, 1)
		return
	}

	for _, destination := range destinations[cave] {
		_, found := visited[destination.name]
		if !found || destination.isBig {
			findPaths2(destination.name, destinations, copyMap(visited), paths, twice)
			continue
		}
		if !twice && cave != "start" {
			findPaths2(destination.name, destinations, copyMap(visited), paths, true)
		}
	}
}

// ----------------------------------------
// utils
// ----------------------------------------

func isUppercase(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func copyMap(in map[string]struct{}) map[string]struct{} {
	out := make(map[string]struct{})
	for k, v := range in {
		out[k] = v
	}
	return out
}
