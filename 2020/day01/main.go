package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	input := input()
	p := findSumPair(input, 2020)
	fmt.Printf("part 1: %d\n", p.x*p.y)
	p2 := findSumVector3(input, 2020)
	fmt.Printf("part 2: %d\n", p2.x*p2.y*p2.z)
}

// ----------------------------------------
// part 1
// ----------------------------------------

type vector struct {
	x, y int64
}

func findSumPair(input []int64, sum int64) vector {
	for i := 0; i < len(input); i++ {
		for j := len(input) - 1; j > 0; j-- {
			if i == j {
				continue
			}
			if input[i]+input[j] == sum {
				return vector{input[i], input[j]}
			}
		}
	}
	return vector{}
}

// ----------------------------------------
// part 2
// ----------------------------------------

type vector3 struct {
	x, y, z int64
}

func findSumVector3(input []int64, sum int64) vector3 {
	size := len(input)
	for i := 0; i < size; i++ {
		for j := size - 1; j > 0; j-- {
			if i == j {
				continue
			}
			for k := 0; k < size; k++ {
				if j == k {
					continue
				}
				if input[i]+input[j]+input[k] == sum {
					return vector3{input[i], input[j], input[k]}
				}
			}
		}
	}
	return vector3{}
}

// ----------------------------------------
// utils
// ----------------------------------------

func input() (in []int64) {
	box := packr.New("day01", "./2020/day01")
	s, err := box.FindString("input")
	check(err)
	for _, txt := range strings.Split(strings.TrimSuffix(s, "\n"), "\n") {
		i, err := strconv.ParseInt(txt, 10, 64)
		check(err)
		in = append(in, i)
	}
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
