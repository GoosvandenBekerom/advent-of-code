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

// ----------------------------------------
// solution
// ----------------------------------------

func part1() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	var fish []int
	for _, raw := range strings.Split(input, ",") {
		fish = append(fish, toInt(raw))
	}

	days := 80
	for i := 0; i < days; i++ {
		for i, f := range fish {
			n := f - 1

			if n == -1 {
				n = 6
				fish = append(fish, 8)
			}

			fish[i] = n
		}

	}

	fmt.Printf("solution: %d\n", len(fish))
}

func part2() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	fish := make(map[int]int)
	for _, raw := range strings.Split(input, ",") {
		fish[toInt(raw)]++
	}

	for day := 0; day < 256; day++ {
		var prev int
		for i := 8; i >= 0; i-- {
			current := fish[i]
			fish[i] = prev
			prev = current
		}
		fish[6] += prev
		fish[8] += prev
	}

	var sum int
	for _, n := range fish {
		sum += n
	}

	fmt.Printf("solution: %d\n", sum)
}

// ----------------------------------------
// utils
// ----------------------------------------

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func toInt(s string) int {
	v, err := strconv.Atoi(s)
	check(err)
	return v
}
