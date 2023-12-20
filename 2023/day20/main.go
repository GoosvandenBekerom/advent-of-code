package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/GoosvandenBekerom/advent-of-code/math"
	"golang.org/x/exp/maps"

	"github.com/GoosvandenBekerom/advent-of-code/data"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

type moduleType uint8

const (
	broadcaster moduleType = iota
	flipflop
	conjunction
)

type module struct {
	name       string
	t          moduleType
	on         bool            // only applicable for flipflop modules
	highPulse  bool            // false is low pulse
	history    map[string]bool // name of downstream module -> last pulse (false = low, true = high)
	downstream []string
}

func parse(lines []string) map[string]*module {
	modules := make(map[string]*module)
	for _, line := range lines {
		typeAndName, rawModules, _ := strings.Cut(line, " -> ")
		m := module{t: broadcaster, downstream: strings.Split(rawModules, ", ")}
		name := typeAndName
		switch typeAndName[0] {
		case '%':
			name = typeAndName[1:]
			m.t = flipflop
		case '&':
			name = typeAndName[1:]
			m.t = conjunction
			m.history = make(map[string]bool)
		}
		m.name = name
		modules[name] = &m
	}
	// initialize history maps
	for name, m := range modules {
		for _, ds := range m.downstream {
			dsm, exists := modules[ds]
			if !exists {
				continue
			}
			if dsm.history != nil {
				dsm.history[name] = false
			}
		}
	}
	return modules
}

func pushButton(modules map[string]*module, moduleToFind string) (lowSent, highSent int, found bool) {
	lowSent = 1
	start := modules["broadcaster"]
	q := data.NewQueue[*module]()
	q.Enqueue(start)
	for !q.Empty() {
		current := q.Dequeue()
		if current.highPulse {
			if current.name == moduleToFind {
				return 0, 0, true
			}
			highSent += len(current.downstream)
		} else {
			lowSent += len(current.downstream)
		}
		for _, dsm := range current.downstream {
			next, exists := modules[dsm]
			if !exists {
				break
			}
			switch next.t {
			case flipflop:
				if !current.highPulse {
					next.on = !next.on
					next.highPulse = next.on
					q.Enqueue(next)
				}
			case conjunction:
				next.history[current.name] = current.highPulse
				nextPulse := false
				for _, lastPulse := range next.history {
					if !lastPulse {
						nextPulse = true
						break
					}
				}
				next.highPulse = nextPulse
				q.Enqueue(next)
			}
		}
	}
	return lowSent, highSent, false
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	modules := parse(lines)
	var lowCount, highCount int
	for i := 0; i < 1000; i++ {
		low, high, _ := pushButton(modules, "")
		lowCount += low
		highCount += high
	}
	return lowCount * highCount
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	modulesToFind := map[string]int{
		"lk": -1,
		"fn": -1,
		"fh": -1,
		"hh": -1,
	}
	for needle := range modulesToFind {
		modules := parse(lines)
		count := 0
		for {
			count++
			_, _, done := pushButton(modules, needle)
			if done {
				modulesToFind[needle] = count
				break
			}
		}
	}

	fmt.Printf("%v\n", modulesToFind)
	counts := maps.Values(modulesToFind)
	return math.LeastCommonMultiple(counts[0], counts[1], counts[2], counts[3])
}
