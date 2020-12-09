package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	fmt.Printf("part 1: %d\n", part1())
	fmt.Printf("part 2: %d\n", part2())
}

func part1() int {
	program := input()
	acc, _ := process(program)
	return acc
}

func part2() int {
	program := input()
	for i := 1; i < len(program)+1; i++ {
		mutated := mutate(program, i)
		acc, err := process(mutated)
		if err == nil {
			return acc
		}
	}
	return 0
}

func mutate(program program, nthOp int) (mutated program) {
	mutated = cp(program)
	count := 0
	for line, op := range program {
		if op.optype == JMP {
			count++
			if count == nthOp {
				mutated[line] = operation{
					optype: NOP,
					param:  op.param,
				}
				return
			}
		}
		if op.optype == NOP {
			count++
			if count == nthOp {
				mutated[line] = operation{
					optype: JMP,
					param:  op.param,
				}
				return
			}
		}
	}
	panic("no operation found to mutate")
}

func process(program program) (acc int, err error) {
	history := make(map[int]bool)
	current := 0
	for {
		if _, dupe := history[current]; dupe {
			return acc, errors.New("infinite loop detected")
		}
		if current >= len(program) {
			return
		}
		history[current] = true
		op := program[current]
		switch op.optype {
		case ACC:
			acc += op.param
		case JMP:
			current += op.param
			continue
		case NOP:
			break
		default:
			log.Fatalf("program reached unknown optype %v", op.optype)
		}
		current++
	}
	return
}

// ----------------------------------------
// utils
// ----------------------------------------

type optype int

const (
	ACC optype = iota
	JMP optype = iota
	NOP optype = iota
)

// operation is a combination of a operation type and its single parameter
type operation struct {
	optype optype
	param  int
}

// program has its line number as its key and its operation as its value
type program map[int]operation

func input() program {
	box := packr.New("day08", "./2020/day08")
	s, err := box.FindString("input")
	check(err)
	program := make(map[int]operation)
	for i, in := range strings.Split(strings.TrimSuffix(s, "\n"), "\n") {
		parts := strings.Split(in, " ")
		program[i] = operation{
			optype: optypeFromString(parts[0]),
			param:  toSignedInt(parts[1]),
		}
	}
	return program
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func optypeFromString(raw string) (t optype) {
	switch raw {
	case "acc":
		t = ACC
	case "jmp":
		t = JMP
	case "nop":
		t = NOP
	default:
		log.Fatalf("got unknown operation type %s", raw)
	}
	return
}

func toSignedInt(raw string) int {
	positive := raw[0] == '+'
	val, err := strconv.Atoi(raw[1:])
	check(err)
	if positive {
		return val
	}
	return -val
}

func cp(original map[int]operation) (copy map[int]operation) {
	copy = make(map[int]operation)
	for key, value := range original {
		copy[key] = value
	}
	return
}
