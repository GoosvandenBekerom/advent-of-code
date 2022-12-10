// Package computer contains the computer used in day 10
// because I have a feeling it's going to be a returning thing in later days
// I figured, lets move it to its own package while I have the time.
package computer

import (
	"github.com/GoosvandenBekerom/advent-of-code/utils"
	"strings"
)

// CPU represents a computing unit with cycles and registers
type CPU struct {
	Cycle int
	X     int
}

// Tick increments the cpu cycle
func (cpu *CPU) Tick() {
	cpu.Cycle++
}

// CalculateSignalStrength calculates the signal strength of the current CPU cycle
// It does so by multiplying the current cycle with the current value of register X.
func (cpu *CPU) CalculateSignalStrength() int {
	return cpu.Cycle * cpu.X
}

// OP represents a CPU operation
// it contains its cost in cpu cycles and an execute method
type OP struct {
	CycleCost int
	execute   func(string, *CPU)
}

// OPs contains a map of supported CPU operations
var OPs = map[string]OP{
	"noop": {CycleCost: 1},
	"addx": {CycleCost: 2, execute: func(arg string, cpu *CPU) {
		cpu.X += utils.ToInt(arg)
	}},
}

// Instruction is a simple wrapper around an OP
// and an accompanying execute argument
type Instruction struct {
	OP  OP
	arg string
}

// Execute checks if the contained OP has an execute method
// and if it does, it invokes the execute method on the given CPU
func (i Instruction) Execute(cpu *CPU) {
	if i.OP.execute != nil {
		i.OP.execute(i.arg, cpu)
	}
}

// ParseInstructions takes raw computer instructions and
// maps them to Instruction objects to interact with.
func ParseInstructions(rawInstructions []string) []Instruction {
	var instructions []Instruction
	for _, rawInstruction := range rawInstructions {
		instruction, arg, _ := strings.Cut(rawInstruction, " ")

		op, found := OPs[instruction]
		if !found {
			panic("unknown instruction: " + instruction)
		}

		instructions = append(instructions, Instruction{
			OP:  op,
			arg: arg,
		})
	}
	return instructions
}
