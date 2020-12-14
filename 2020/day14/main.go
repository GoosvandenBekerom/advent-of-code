package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	println(part1())
	println(part2())
}

/*
rules:
If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
Otherwise, the seat's state does not change.
*/
func part1() uint64 {
	var orMask, andMask uint64
	mem := make(map[int]uint64)
	for line := range input() {
		if line[:4] == "mask" {
			raw := strings.Split(line, " = ")[1]
			orMask = toUint36(strings.ReplaceAll(raw, "X", "0"))
			andMask = toUint36(strings.ReplaceAll(raw, "X", "1"))
			continue
		}
		address, value := parseMem(line)
		mem[address] = (value | orMask) & andMask
	}
	var sum uint64
	for _, v := range mem {
		sum += v
	}
	return sum
}

func part2() uint64 {
	mem := make(map[uint64]uint64)
	var mask string
	for line := range input() {
		if line[:4] == "mask" {
			mask = strings.Split(line, " = ")[1]
			continue
		}
		address, value := parseMem(line)
		for _, a := range allAddresses(uint64(address), mask) {
			mem[a] = value
		}
	}
	var sum uint64
	for _, v := range mem {
		sum += v
	}
	return sum
}

func toUint36(raw string) uint64 {
	val, err := strconv.ParseUint(raw, 2, 36)
	check(err)
	return val
}

func parseMem(line string) (int, uint64) {
	pattern := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
	matches := pattern.FindAllStringSubmatch(line, -1)
	address, err := strconv.Atoi(matches[0][1])
	check(err)
	value, err := strconv.ParseUint(matches[0][2], 10, 36)
	check(err)
	return address, value
}

func allAddresses(address uint64, mask string) (addresses []uint64) {
	var masks []uint64
	orMask := toUint36(strings.ReplaceAll(mask, "X", "0"))
	str := strings.Map(func(r rune) rune {
		switch r {
		case '1':
			return '0'
		case 'X':
			return '1'
		default:
			return r
		}
	}, mask)
	andNotMask := toUint36(str)
	address = (address | orMask) &^ andNotMask
	allMasks(mask, &masks)
	for _, m := range masks {
		addresses = append(addresses, address|m)
	}
	return addresses
}

func allMasks(raw string, masks *[]uint64) {
	if strings.IndexByte(raw, 'X') < 0 {
		*masks = append(*masks, toUint36(raw))
	} else {
		allMasks(strings.Replace(raw, "X", "0", 1), masks)
		allMasks(strings.Replace(raw, "X", "1", 1), masks)
	}
}

// ----------------------------------------
// utils
// ----------------------------------------

func input() (lines chan string) {
	box := packr.New("day14", "./2020/day14")
	s, err := box.FindString("input")
	check(err)
	lines = make(chan string, 1)
	go func() {
		defer close(lines)
		for _, line := range strings.Split(strings.TrimSuffix(s, "\n"), "\n") {
			lines <- line
		}
	}()
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
