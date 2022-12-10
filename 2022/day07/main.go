package main

import (
	_ "embed"
	"fmt"
	"github.com/GoosvandenBekerom/advent-of-code/data"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
	"sort"
	"strings"
)

//go:embed input
var input string

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	sizes := calculateDirectorySizes(input)

	var sum int
	for _, size := range sizes {
		if size <= 100000 {
			sum += size
		}
	}

	return sum
}

func part2(input string) int {
	fmt.Println("___________________________________________")
	fmt.Println("part 2:")

	sizes := calculateDirectorySizes(input)

	orderedSizes := make([]int, len(sizes))
	var i int
	for _, size := range sizes {
		orderedSizes[i] = size
		i++
	}

	sort.Slice(orderedSizes, func(i, j int) bool { return orderedSizes[i] > orderedSizes[j] })

	required := 30_000_000 - (70_000_000 - sizes["/"])
	previous := 70_000_000
	for _, size := range orderedSizes {
		if size < required {
			return previous
		}
		previous = size
	}

	return -1
}

func calculateDirectorySizes(input string) map[string]int {
	stack := make(data.Stack[string], 0)
	var listing bool
	sizes := make(map[string]int)

	for _, line := range strings.Split(input, "\n") {
		if line[0] == '$' {
			listing = false
			switch line[2:4] {
			case "cd":
				arg := line[5:]
				switch arg {
				case "/":
					stack = data.Stack[string]{"/"}
					continue
				case "..":
					stack.Pop()
					continue
				default:
					stack.Push(arg)
					break
				}
			case "ls":
				listing = true
				continue
			}
		}

		if listing {
			if line[:4] == "dir " {
				continue
			}

			rawSize, _, _ := strings.Cut(line, " ")

			size := utils.ToInt(rawSize)
			var path string
			for _, dir := range stack {
				path += dir
				sizes[path] += size
			}
		}
	}

	return sizes
}
