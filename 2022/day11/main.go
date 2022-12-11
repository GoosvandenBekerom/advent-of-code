package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

type monkey struct {
	items               []int
	calculateWorryLevel func(int) int
	divider             int
	findMonkeyToThrowTo func(int, int) int
}

func input() []monkey {
	return []monkey{
		{ // Monkey 0
			items:               []int{98, 70, 75, 80, 84, 89, 55, 98},
			calculateWorryLevel: func(old int) int { return old * 2 },
			divider:             11,
			findMonkeyToThrowTo: func(new int, divider int) int {
				if new%divider == 0 {
					return 1
				}
				return 4
			},
		},
		{ // Monkey 1
			items:               []int{59},
			calculateWorryLevel: func(old int) int { return old * old },
			divider:             19,
			findMonkeyToThrowTo: func(new int, divider int) int {
				if new%divider == 0 {
					return 7
				}
				return 3
			},
		},
		{ // Monkey 2
			items:               []int{77, 95, 54, 65, 89},
			calculateWorryLevel: func(old int) int { return old + 6 },
			divider:             7,
			findMonkeyToThrowTo: func(new int, divider int) int {
				if new%divider == 0 {
					return 0
				}
				return 5
			},
		},
		{ // Monkey 3
			items:               []int{71, 64, 75},
			calculateWorryLevel: func(old int) int { return old + 2 },
			divider:             17,
			findMonkeyToThrowTo: func(new int, divider int) int {
				if new%divider == 0 {
					return 6
				}
				return 2
			},
		},
		{ // Monkey 4
			items:               []int{74, 55, 87, 98},
			calculateWorryLevel: func(old int) int { return old * 11 },
			divider:             3,
			findMonkeyToThrowTo: func(new int, divider int) int {
				if new%divider == 0 {
					return 1
				}
				return 7
			},
		},
		{ // Monkey 5
			items:               []int{90, 98, 85, 52, 91, 60},
			calculateWorryLevel: func(old int) int { return old + 7 },
			divider:             5,
			findMonkeyToThrowTo: func(new int, divider int) int {
				if new%divider == 0 {
					return 0
				}
				return 4
			},
		},
		{ // Monkey 6
			items:               []int{99, 51},
			calculateWorryLevel: func(old int) int { return old + 1 },
			divider:             13,
			findMonkeyToThrowTo: func(new int, divider int) int {
				if new%divider == 0 {
					return 5
				}
				return 2
			},
		},
		{ // Monkey 7
			items:               []int{98, 94, 59, 76, 51, 65, 75},
			calculateWorryLevel: func(old int) int { return old + 5 },
			divider:             2,
			findMonkeyToThrowTo: func(new int, divider int) int {
				if new%divider == 0 {
					return 3
				}
				return 6
			},
		},
	}
}

func main() {
	fmt.Println(part1(input(), false))
	fmt.Println(part2(input(), false))
}

func printState(round int, monkeys []monkey) {
	fmt.Printf("after round %d:\n", round)
	for i, m := range monkeys {
		items := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(m.items)), ", "), "[]")
		fmt.Printf("Monkey %d: %s\n", i, items)
	}
	println()
}

func play(monkeys []monkey, rounds int, lcm int, verbose bool) int {
	inspections := make([]int, len(monkeys))
	for round := 1; round < rounds+1; round++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				newWorryLevel := m.calculateWorryLevel(item)
				if lcm > 0 {
					newWorryLevel = newWorryLevel % lcm
				} else {
					newWorryLevel = newWorryLevel / 3
				}
				nextMonkey := m.findMonkeyToThrowTo(newWorryLevel, m.divider)
				monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, newWorryLevel)
				inspections[i]++
			}
			monkeys[i].items = m.items[:0]
		}
		if verbose {
			printState(round, monkeys)
		}
	}
	if verbose {
		for i, amount := range inspections {
			fmt.Printf("Monkey %d inspected items %d times\n", i, amount)
		}
	}
	sort.Ints(inspections)
	return inspections[len(inspections)-2] * inspections[len(inspections)-1]
}

func part1(monkeys []monkey, verbose bool) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	return play(monkeys, 20, -1, verbose)
}

func part2(monkeys []monkey, verbose bool) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	// https://en.wikipedia.org/wiki/Least_common_multiple
	lcm := 1
	for _, m := range monkeys {
		lcm *= m.divider
	}

	return play(monkeys, 10000, lcm, verbose)
}
