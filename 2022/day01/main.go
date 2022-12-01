package main

import (
	_ "embed"
	"sort"
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

// ----------------------------------------
// solution
// ----------------------------------------

func part1() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	var currentElf, largestElf int64
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			if currentElf > largestElf {
				largestElf = currentElf
			}
			currentElf = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		currentElf += int64(calories)
	}

	if currentElf > largestElf {
		largestElf = currentElf
	}

	fmt.Printf("most calories caried by single elf: %d\n", largestElf)
}

func part2() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	var currentElf int
	var elfs []int

	//pq := make(datastructures.PriorityQueue, 3)

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			elfs = append(elfs, currentElf)
			currentElf = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		currentElf += calories
	}

	sort.Ints(elfs)

	var sum int
	for i := len(elfs) - 1; i > len(elfs)-4; i-- {
		sum += elfs[i]
	}

	fmt.Printf("amount of calories caried by top 3 elfs: %d\n", sum)
}

// ----------------------------------------
// utils
// ----------------------------------------

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
