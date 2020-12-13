package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/deanveloper/modmath/v1/bigmod"
	"github.com/gobuffalo/packr/v2"
)

func main() {
	println(part1())
	println(part2().String())
}

func part1() int {
	earliest, busses := input(true)
	fmt.Printf("earliest time to leave %d, busses: %v \n", earliest, busses)

	shortestWait := max(busses)
	shortestBus := shortestWait
	for _, bus := range busses {
		n := earliest + bus/2
		closest := n - (n % bus)
		if closest < earliest {
			closest += bus
		}
		remainder := closest - earliest
		if remainder < shortestWait {
			shortestWait = remainder
			shortestBus = bus
		}
	}

	return shortestWait * shortestBus
}

func part2() *big.Int {
	_, busses := input(false)

	var results []bigmod.CrtEntry
	for off, bus := range busses {
		if bus < 0 {
			continue
		}
		results = append(results, bigmod.CrtEntry{A: big.NewInt(int64(bus - off)), N: big.NewInt(int64(bus))})
	}

	return bigmod.SolveCrtMany(results)
}

// ----------------------------------------
// utils
// ----------------------------------------

func input(ignoreX bool) (earliest int, busses []int) {
	box := packr.New("day13", "./2020/day13")
	s, err := box.FindString("input")
	check(err)
	lines := strings.Split(strings.TrimSuffix(s, "\n"), "\n")
	earliest, err = strconv.Atoi(lines[0])
	check(err)
	for _, raw := range strings.Split(lines[1], ",") {
		if raw == "x" {
			if ignoreX {
				continue
			}
			raw = "-1"
		}
		bus, err := strconv.Atoi(raw)
		check(err)
		busses = append(busses, bus)
	}
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func max(ints []int) (max int) {
	for _, i := range ints {
		if i > max {
			max = i
		}
	}
	return
}
