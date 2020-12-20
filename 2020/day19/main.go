package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	//println(part1())
	println(part2())
}

var rules []*rule

func part1() (sum int) {
	var messages <-chan string
	rules, messages = input()

	patterns := make(map[int]string)

	for {
		for i, rule := range rules {
			if rule.char > 0 {
				patterns[i] = string(rule.char)
				continue
			}
			temp := ""
			for _, pos := range rule.sequence1 {
				if _, ok := patterns[pos]; !ok {
					temp = ""
					break
				}
				temp += patterns[pos]
			}
			if temp != "" && len(rule.sequence2) == 0 {
				patterns[i] = temp
				continue
			}
			if temp == "" {
				continue
			}
			temp = "(" + temp + "|"
			for _, pos := range rule.sequence2 {
				if _, ok := patterns[pos]; !ok {
					temp = ""
					break
				}
				temp += patterns[pos]
			}
			if temp != "" {
				temp = temp + ")"
				patterns[i] = temp
			}
		}

		if len(patterns) == len(rules) {
			break
		}
	}

	matcher := regexp.MustCompile("^" + patterns[0] + "$")

	for msg := range messages {
		if matcher.Match([]byte(msg)) {
			sum++
		}
	}
	return
}

func part2() (sum int) {
	rules, messages := input2()

	patterns := make(map[int]string)

	for {
		for i, rule := range rules {
			if rule.char > 0 {
				patterns[i] = string(rule.char)
				continue
			}
			temp1 := ""
			for _, pos := range rule.sequence1 {
				if _, ok := rules[pos]; !ok {
					continue
				}
				if _, ok := patterns[pos]; !ok {
					temp1 = ""
					break
				}
				temp1 += patterns[pos]
			}
			if temp1 != "" && len(rule.sequence2) == 0 {
				patterns[i] = temp1
				continue
			}
			if temp1 == "" {
				continue
			}
			if i == 8 {
				patterns[i] = "(" + temp1 + ")+"
				continue
			}
			if i == 11 {
				p42 := patterns[42]
				p31 := patterns[31]

				patterns[i] = fmt.Sprintf("%[1]s(%[1]s(%[1]s(%[1]s%[2]s)?%[2]s)?%[2]s)?%[2]s", p42, p31)
				continue
			}
			temp2 := "(" + temp1 + "|"
			for _, pos := range rule.sequence2 {
				if _, ok := rules[pos]; !ok {
					continue
				}
				if _, ok := patterns[pos]; !ok {
					temp2 = ""
					break
				}
				temp2 += patterns[pos]
			}
			if temp2 != "" {
				temp2 = temp2 + ")"
				patterns[i] = temp2
			}
		}

		if len(patterns) == len(rules) {
			break
		}
	}

	matcher := regexp.MustCompile("^" + patterns[0] + "$")

	for msg := range messages {
		if matcher.Match([]byte(msg)) {
			println(msg)
			sum++
		}
	}
	return
}

// ----------------------------------------
// domain
// ----------------------------------------

// 3 types
// - ordered pointers to rules (just digits separated by space)
// - 1 of 2 ordered pointers to rules (2 groups of digits separated by |)
// - either "a" or "b" (including ")
type rule struct {
	char                 byte
	sequence1, sequence2 []int
}

// ----------------------------------------
// utils
// ----------------------------------------

func input() (rules []*rule, messages chan string) {
	box := packr.New("day19", "./2020/day19")
	s, err := box.FindString("input")
	check(err)
	// parts[0] raw rules
	// parts[1] raw messages
	parts := strings.Split(strings.TrimSuffix(s, "\n"), "\n\n")
	matcher := regexp.MustCompile(`(\d+): ([a-b0-9 "]+)\|?([0-9 ]+)?`)
	allMatches := matcher.FindAllStringSubmatch(parts[0], -1)
	rules = make([]*rule, len(allMatches))
	for _, matches := range allMatches {
		rule := &rule{}
		if matches[2][0] == '"' {
			rule.char = matches[2][1]
		} else {
			for _, raw := range strings.Split(strings.Trim(matches[2], " "), " ") {
				rule.sequence1 = append(rule.sequence1, toInt(raw))
			}
		}
		if len(matches[3]) > 0 {
			for _, raw := range strings.Split(strings.Trim(matches[3], " "), " ") {
				rule.sequence2 = append(rule.sequence2, toInt(raw))
			}
		}
		rules[toInt(matches[1])] = rule
	}

	messages = make(chan string, 1)
	go func() {
		defer close(messages)
		for _, message := range strings.Split(parts[1], "\n") {
			messages <- message
		}
	}()
	return
}

func input2() (rules map[int]*rule, messages chan string) {
	box := packr.New("day19", "./2020/day19")
	s, err := box.FindString("input")
	check(err)
	// parts[0] raw rules
	// parts[1] raw messages
	parts := strings.Split(strings.TrimSuffix(s, "\n"), "\n\n")
	matcher := regexp.MustCompile(`(\d+): ([a-b0-9 "]+)\|?([0-9 ]+)?`)
	allMatches := matcher.FindAllStringSubmatch(parts[0], -1)
	rules = make(map[int]*rule)
	for _, matches := range allMatches {
		// replace rules
		// 8: 42 | 42 8
		// 11: 42 31 | 42 11 31
		pos := toInt(matches[1])
		if pos == 8 {
			matches[2] = "42"
			matches[3] = "42 8"
		}
		if pos == 11 {
			matches[2] = "42 31"
			matches[3] = "42 11 31"
		}
		rule := &rule{}
		if matches[2][0] == '"' {
			rule.char = matches[2][1]
		} else {
			for _, raw := range strings.Split(strings.Trim(matches[2], " "), " ") {
				rule.sequence1 = append(rule.sequence1, toInt(raw))
			}
		}
		if len(matches[3]) > 0 {
			for _, raw := range strings.Split(strings.Trim(matches[3], " "), " ") {
				rule.sequence2 = append(rule.sequence2, toInt(raw))
			}
		}
		rules[toInt(matches[1])] = rule
	}

	messages = make(chan string, 1)
	go func() {
		defer close(messages)
		for _, message := range strings.Split(parts[1], "\n") {
			messages <- message
		}
	}()
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
