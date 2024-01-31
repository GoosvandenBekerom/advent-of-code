package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	Min := math.MaxInt
	Max := 0
	for i := 0; i < n; i++ {
		s.Scan()
		points := score(s.Text())
		if points > Max {
			Max = points
		}
		if points < Min {
			Min = points
		}
	}
	fmt.Printf("%d %d", Max, Min)
}

func score(s string) int {
	sizeS, countS, _ := strings.Cut(s, " ")

	count, _ := strconv.Atoi(countS)
	size, _ := strconv.Atoi(sizeS)
	sum := 1
	current := 1
	for i := 0; i < count; i++ {
		current += size
		sum += current
	}

	return sum
}
