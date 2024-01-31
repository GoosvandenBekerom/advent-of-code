package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Sample Input 0

11 3
1 2 5
Sample Output 0

3
*/

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	numPackages := strings.Split(s.Text(), " ")[0]
	s.Scan()
	bikeTypes := Map(strings.Split(s.Text(), " "), func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	})

	sort.Ints(bikeTypes)

	fmt.Printf("%d", answer)
}

func Map[I, O any](in []I, f func(item I) O) (out []O) {
	for _, item := range in {
		out = append(out, f(item))
	}
	return out
}
