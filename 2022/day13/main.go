package main

import (
	_ "embed"
	"fmt"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
	"sort"
	"strings"
)

//go:embed input
var input string

func main() {
	fmt.Println(part1(parseInput()))
	fmt.Println(part2(parseInput()))
}

func parseInput() [][2]string {
	rawPairs := strings.Split(input, "\n\n")
	pairs := make([][2]string, len(rawPairs))
	for i, pair := range rawPairs {
		first, second, _ := strings.Cut(pair, "\n")
		pairs[i] = [2]string{first, second}
	}
	return pairs
}

type packet []any

func parsePacket(raw string) (out packet, remainder string) {
	raw = raw[1:]
	var p packet
	for len(raw) > 0 {
		switch {
		case raw[0] == ']':
			return out, raw[1:]
		case raw[0] == ',':
			raw = raw[1:]
		case raw[0] == '[':
			p, raw = parsePacket(raw)
			out = append(out, p)
		case raw[0] >= '0' && raw[0] <= '9':
			var size int
			for size = 0; size < len(raw) && raw[size] >= '0' && raw[size] <= '9'; size++ {
				// count amount of digits that are part of the number
			}
			out = append(out, utils.ToInt(raw[:size]))
			raw = raw[size:]
		}
	}
	return
}

func comparePackets(packet1, packet2 packet) (diff int) {
	for i := 0; i < len(packet1) && i < len(packet2) && diff == 0; i++ {
		switch packet1[i].(type) {
		case packet:
			switch packet2[i].(type) {
			case packet:
				diff = comparePackets(packet1[i].(packet), packet2[i].(packet))
			default:
				diff = comparePackets(packet1[i].(packet), packet{packet2[i].(int)})
			}
		default:
			switch packet2[i].(type) {
			case packet:
				diff = comparePackets(packet{packet1[i].(int)}, packet2[i].(packet))
			default:
				diff = packet1[i].(int) - packet2[i].(int)
			}
		}
	}
	if diff != 0 {
		return diff
	}
	return len(packet1) - len(packet2)
}

func part1(pairs [][2]string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	sum := 0
	for i, pair := range pairs {
		packet1, _ := parsePacket(pair[0])
		packet2, _ := parsePacket(pair[1])
		if comparePackets(packet1, packet2) < 0 {
			sum += i + 1
		}
	}
	return sum
}

func part2(pairs [][2]string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	var packets []packet
	for _, pair := range pairs {
		packet1, _ := parsePacket(pair[0])
		packet2, _ := parsePacket(pair[1])
		packets = append(packets, packet1)
		packets = append(packets, packet2)
	}
	divider1, divider2 := packet{2}, packet{6}
	packets = append(packets, divider1, divider2)

	sort.Slice(packets, func(i, j int) bool {
		return comparePackets(packets[i], packets[j]) < 0
	})

	solution := 1
	for i, p := range packets {
		if comparePackets(p, divider1) == 0 || comparePackets(p, divider2) == 0 {
			solution *= i + 1
		}
	}
	return solution
}
