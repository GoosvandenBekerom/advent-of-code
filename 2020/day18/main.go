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
	println()
	println(part2())
}

func part1() (sum int) {
	for expression := range input() {
		expression = parse1(expression)
		sum += eval1(expression)
	}
	return
}

func parse1(expression string) string {
	for {
		i1 := strings.LastIndex(expression, "(")
		if i1 < 0 {
			break
		}
		i2 := strings.Index(expression[i1:], ")") + i1
		if i2 < 0 {
			panic("found '(' without ')'")
		}
		sub := expression[i1+1 : i2]
		evaluated := eval1(sub)
		expression = expression[:i2] + expression[i2+1:]
		expression = expression[:i1] + strconv.Itoa(evaluated) + expression[i2:]
	}
	return expression
}

func eval1(expression string) int {
	solve := func(a, b int, method rune) int {
		switch method {
		case '+':
			return a + b
		case '*':
			return a * b
		default:
			panic("unsupported method: " + string(method))
		}
	}

	matcher := regexp.MustCompile(`(\d+) ([+*]) (\d+)`)

	for {
		// 0 = start; 			1 = end; 			2 = start a; 3 = end a
		// 4 = start method; 	5 = end method; 	6 = start b; 7 = enb b
		matches := matcher.FindStringSubmatchIndex(expression)
		if len(matches) < 4 {
			break
		}
		a, err := strconv.Atoi(expression[matches[2]:matches[3]])
		check(err)
		b, err := strconv.Atoi(expression[matches[6]:matches[7]])
		check(err)
		evaluated := solve(a, b, []rune(expression[matches[4]:matches[5]])[0])
		expression = expression[:matches[0]] + strconv.Itoa(evaluated) + expression[matches[1]:]
		println("next " + expression)
	}

	solution, err := strconv.Atoi(expression)
	check(err)
	return solution
}

func part2() (sum int) {
	for expression := range input() {
		expression = parse2(expression)
		sum += eval2(expression)
	}
	return
}

func parse2(expression string) string {
	for {
		i1 := strings.LastIndex(expression, "(")
		if i1 < 0 {
			break
		}
		i2 := strings.Index(expression[i1:], ")") + i1
		if i2 < 0 {
			panic("found '(' without ')'")
		}
		sub := expression[i1+1 : i2]
		evaluated := eval2(sub)
		expression = expression[:i2] + expression[i2+1:]
		expression = expression[:i1] + strconv.Itoa(evaluated) + expression[i2:]
	}
	return expression
}

func eval2(expression string) int {
	solve := func(a, b int, method rune) int {
		switch method {
		case '+':
			return a + b
		case '*':
			return a * b
		default:
			panic("unsupported method: " + string(method))
		}
	}

	eval := func(m *regexp.Regexp, method rune) {
		for {
			// 0 = start; 1 = end; 2 = start a; 3 = end a; 4 = start b; 5 = end b;
			matches := m.FindStringSubmatchIndex(expression)
			if matches == nil {
				break
			}
			a, err := strconv.Atoi(expression[matches[2]:matches[3]])
			check(err)
			b, err := strconv.Atoi(expression[matches[4]:matches[5]])
			check(err)
			evaluated := solve(a, b, method)
			expression = expression[:matches[0]] + strconv.Itoa(evaluated) + expression[matches[1]:]
			println("next " + expression)
		}
	}

	eval(regexp.MustCompile(`(\d+) \+ (\d+)`), '+')
	eval(regexp.MustCompile(`(\d+) \* (\d+)`), '*')

	solution, err := strconv.Atoi(expression)
	check(err)
	return solution
}

// ----------------------------------------
// utils
// ----------------------------------------

func input() (raw chan string) {
	box := packr.New("day18", "./2020/day18")
	s, err := box.FindString("input")
	check(err)
	raw = make(chan string, 1)
	go func() {
		defer close(raw)
		for _, line := range strings.Split(strings.TrimSuffix(s, "\n"), "\n") {
			raw <- line
		}
	}()
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
