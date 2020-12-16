package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	//println(part1())
	println(part2())
}

func part1() int {
	props, _, nearby := input()
	var invalid []int
	for _, t := range nearby {
		for _, value := range t {
			valid := false
			for _, prop := range props {
				if (value >= prop.low1 && value <= prop.high1) ||
					(value >= prop.low2 && value <= prop.high2) {
					valid = true
					break
				}
			}
			if !valid {
				invalid = append(invalid, value)
			}
		}
	}
	sum := 0
	for _, val := range invalid {
		sum += val
	}
	return sum
}

func part2() int {
	props, own, nearby := input()

	validTickets := getValidTickets(props, nearby)
	allPositions := getPositionOptions(props, validTickets)
	filtered := filterPositionOptions(allPositions)

	var departurePositions []int
	for pos, properties := range filtered {
		fmt.Printf("position %d can contain %v\n", pos, properties)
		for name, origin := range properties {
			if strings.HasPrefix(name, "departure") {
				departurePositions = append(departurePositions, origin)
				break
			}
		}
	}
	fmt.Printf("departure related positions %v\n", departurePositions)
	solution := 1
	for _, p := range departurePositions {
		solution *= own[p]
	}
	return solution
}

func valid(value int, prop property) bool {
	return (value >= prop.low1 && value <= prop.high1) || (value >= prop.low2 && value <= prop.high2)
}

func getValidTickets(props []property, in []ticket) (out []ticket) {
	for _, prop := range props {
		for _, t := range in {
			ticketValid := true
			for _, value := range t {
				if !valid(value, prop) {
					ticketValid = false
				}
			}
			if ticketValid {
				out = append(out, t)
			}
		}
	}
	return
}

func getPositionOptions(props []property, tickets []ticket) []map[string]bool {
	positions := make([]map[string]bool, len(props))

	for i := 0; i < len(positions); i++ {
		positions[i] = make(map[string]bool)
		for _, p := range props {
			positions[i][p.name] = true
			for _, t := range tickets {
				positions[i][p.name] = positions[i][p.name] && valid(t[i], p)
			}
			if !positions[i][p.name] {
				delete(positions[i], p.name)
			}
		}
	}
	return positions
}

type byOptionLen []map[string]int

func (a byOptionLen) Len() int           { return len(a) }
func (a byOptionLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byOptionLen) Less(i, j int) bool { return len(a[i]) < len(a[j]) }

func filterPositionOptions(positions []map[string]bool) []map[string]int {
	withOriginPosition := make([]map[string]int, len(positions))
	for i, options := range positions {
		withOriginPosition[i] = make(map[string]int)
		for option := range options {
			withOriginPosition[i][option] = i
		}
	}
	sort.Sort(byOptionLen(withOriginPosition))
	for p1, options := range withOriginPosition {
		for option := range options {
			for p2, options2 := range withOriginPosition {
				if p1 == p2 {
					continue
				}
				delete(options2, option)
			}
		}
	}
	return withOriginPosition
}

// ----------------------------------------
// domain
// ----------------------------------------

type property struct {
	name                     string
	low1, high1, low2, high2 int
}

type ticket []int

// ----------------------------------------
// utils
// ----------------------------------------

func input() (props []property, own ticket, nearby []ticket) {
	box := packr.New("day16", "./2020/day16")
	s, err := box.FindString("input")
	check(err)
	groups := strings.Split(strings.TrimSuffix(s, "\n"), "\n\n")
	//groups[0] = properties
	propMatcher := regexp.MustCompile(`(.+)[!:] (\d+)-(\d+) or (\d+)-(\d+)`)
	for _, matches := range propMatcher.FindAllStringSubmatch(groups[0], -1) {
		props = append(props, property{
			name:  matches[1],
			low1:  toInt(matches[2]),
			high1: toInt(matches[3]),
			low2:  toInt(matches[4]),
			high2: toInt(matches[5]),
		})
	}
	//groups[1] = own ticket (second line of group)
	ownLine := strings.Split(groups[1], "\n")[1]
	for _, raw := range strings.Split(ownLine, ",") {
		own = append(own, toInt(raw))
	}
	//groups[2] = nearby tickets (starts at second line of group)
	for i, line := range strings.Split(groups[2], "\n")[1:] {
		rawNums := strings.Split(line, ",")
		nearby = append(nearby, make(ticket, len(rawNums)))
		for j, raw := range rawNums {
			nearby[i][j] = toInt(raw)
		}
	}
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}
