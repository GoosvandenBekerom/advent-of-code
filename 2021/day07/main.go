package main

import (
	_ "embed"
	"math"
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

	var positions []int

	for _, raw := range strings.Split(input, ",") {
		positions = append(positions, toInt(raw))
	}

	optimal := median(positions...)

	var fuel int
	for _, pos := range positions {
		fuel += int(math.Abs(float64(pos - optimal)))
	}

	fmt.Printf("solution: %d\n", fuel)
}

func part2() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	raw := strings.Split(input, ",")
	positions := make([]int, 0, len(raw))

	for _, r := range raw {
		positions = append(positions, toInt(r))
	}

	sort.Ints(positions)

	min := positions[0]
	max := positions[len(positions)-1]

	requiredFuel := make([]int, max-min+1)

	for i := min; i <= max; i++ {
		var fuel int
		for _, pos := range positions {
			steps := int(math.Abs(float64(pos - i)))
			step := 1
			for j := 0; j < steps; j++ {
				fuel += step
				step++
			}
		}
		requiredFuel[i] = fuel
	}

	sort.Ints(requiredFuel)

	fmt.Printf("solution: %d\n", requiredFuel[0])
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

func median(nums ...int) int {
	sort.Ints(nums)

	medianIndex := len(nums) / 2

	if len(nums)%2 == 1 {
		return nums[medianIndex]
	}

	return (nums[medianIndex-1] + nums[medianIndex]) / 2
}
