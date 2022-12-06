package main

import (
	_ "embed"
	"fmt"
)

//go:embed input
var input string

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(buffer string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	return findMessageMarker(buffer, 4)
}

func part2(buffer string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	return findMessageMarker(buffer, 14)
}

func findMessageMarker(buffer string, size int) int {
	var cut int
	for {
		c := make(map[byte]struct{})
		for _, b := range []byte(buffer[cut : cut+size]) {
			c[b] = struct{}{}
		}

		if len(c) == size {
			return cut + size
		}

		cut++
	}
}
