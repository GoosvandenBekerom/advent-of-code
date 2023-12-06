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

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	return runGames(
		utils.Map(strings.Fields(strings.Split(lines[0], ":")[1]), func(s string) int { return utils.ToInt(s) }),
		utils.Map(strings.Fields(strings.Split(lines[1], ":")[1]), func(s string) int { return utils.ToInt(s) }),
	)
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	return runGames(
		[]int{utils.ToInt(strings.Join(strings.Fields(strings.Split(lines[0], ":")[1]), ""))},
		[]int{utils.ToInt(strings.Join(strings.Fields(strings.Split(lines[1], ":")[1]), ""))},
	)
}

func runGames(times []int, distances []int) int {
	shortestTimePressed := make([]int, len(times))
	for index, time := range times {
		for i := 1; i <= time/2+1; i++ {
			distance := i * (time - i)
			if distance > distances[index] {
				shortestTimePressed[index] = i
				break
			}
		}
	}

	sum := 1
	for index, time := range times {
		s := shortestTimePressed[index]
		sum *= time - s - s + 1
	}

	return sum
}
