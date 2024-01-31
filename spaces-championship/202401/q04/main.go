package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	s.Scan()
	colleagues := Map(strings.Split(s.Text(), " "), func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	})

	sort.Ints(colleagues)

	// dp

	lowest := colleagues[0]
	highest := colleagues[len(colleagues)-1]

	answer := math.MaxInt
	for i := lowest; i <= highest; i++ {
		x := calculateSteps(i, colleagues)
		if x < answer {
			answer = x
		}
	}

	fmt.Printf("%d", answer)
}

func calculateSteps(i int, colleagues []int) int {
	var sum int
	for _, colleague := range colleagues {
		dist := colleague - i
		if dist < 0 {
			dist *= -1
		}

		sum += dist * 2
	}

	return sum
}

func Map[I, O any](in []I, f func(item I) O) (out []O) {
	for _, item := range in {
		out = append(out, f(item))
	}
	return out
}
