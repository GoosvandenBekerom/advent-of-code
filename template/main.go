package main

import (
	_ "embed"

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
	fmt.Println("part 1:")
	fmt.Println("___________________________________________")
	for _, line := range strings.Split(input, "\n") {
		println(line)
	}
}

func part2() {
	fmt.Println("part 2:")
	fmt.Println("___________________________________________")
}

// ----------------------------------------
// utils
// ----------------------------------------

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
