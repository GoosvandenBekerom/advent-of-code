package main

import (
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	passports := input()
	count1 := part1(passports)
	fmt.Printf("part 1: %d\n", count1)
	count2 := part2(passports)
	fmt.Printf("part 2: %d\n", count2)
}

func part1(passports []passport) (count int) {
	for _, passport := range passports {
		if passport.isComplete() {
			count++
		}
	}
	return
}

func part2(passports []passport) (count int) {
	for _, passport := range passports {
		if passport.isValid() {
			count++
		}
	}
	return
}

func (p passport) isComplete() bool {
	for key := range checks {
		_, ok := p[key]
		if !ok {
			return false
		}
	}
	return true
}

type validation func(string) bool

var checks = map[string]validation{
	"byr": byr,
	"iyr": iyr,
	"eyr": eyr,
	"hgt": hgt,
	"hcl": hcl,
	"ecl": ecl,
	"pid": pid,
}

func (p passport) isValid() bool {
	if !p.isComplete() {
		return false
	}
	for key, check := range checks {
		if !check(p[key]) {
			return false
		}
	}
	return true
}

func year(year string, low, high int) bool {
	if len(year) != 4 {
		return false
	}
	y, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	if y < low || y > high {
		return false
	}
	return true
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func byr(byr string) bool {
	return year(byr, 1920, 2002)
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func iyr(iyr string) bool {
	return year(iyr, 2010, 2020)
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func eyr(eyr string) bool {
	return year(eyr, 2020, 2030)
}

// hgt (Height) - a number followed by either cm or in:
// 	- If cm, the number must be at least 150 and at most 193.
// 	- If in, the number must be at least 59 and at most 76.
func hgt(hgt string) bool {
	check := func(i, low, high int) bool {
		if i < low || i > high {
			return false
		}
		return true
	}

	n, err := strconv.Atoi(hgt[:len(hgt)-2])
	if err != nil {
		return false
	}
	m := hgt[len(hgt)-2:]

	if m == "cm" {
		return check(n, 150, 193)
	}
	if m == "in" {
		return check(n, 59, 76)
	}
	return false
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func hcl(hcl string) bool {
	matched, err := regexp.Match(`^#[0-9a-f]{6}$`, []byte(hcl))
	if err != nil {
		return false
	}
	return matched
}

// ecl (Eye Color) 	- exactly one of: .
func ecl(ecl string) bool {
	matched, err := regexp.Match(`^(amb|blu|brn|gry|grn|hzl|oth)$`, []byte(ecl))
	if err != nil {
		return false
	}
	return matched
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func pid(pid string) bool {
	matched, err := regexp.Match(`^\d{9}$`, []byte(pid))
	if err != nil {
		return false
	}
	return matched
}

// ----------------------------------------
// utils
// ----------------------------------------

type passport map[string]string

func input() (passports []passport) {
	box := packr.New("day04", "./2020/day04")
	s, err := box.FindString("input")
	check(err)
	for _, in := range strings.Split(strings.TrimSuffix(s, "\n"), "\n\n") {
		pp := parse(in)
		passports = append(passports, pp)
	}
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parse(in string) (out map[string]string) {
	fields := strings.FieldsFunc(in, func(r rune) bool {
		return r == ' ' || r == '\n'
	})
	out = make(map[string]string)
	for _, field := range fields {
		kv := strings.Split(field, ":")
		out[kv[0]] = kv[1]
	}
	return
}
