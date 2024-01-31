package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	fmt.Fprintf(os.Stdout, "%d", calculate(s.Text()))
}

func calculate(s string) int {
	left, right, _ := strings.Cut(s, " ")
	domain, country, _ := strings.Cut(left, ".")
	v, _ := strconv.Atoi(right)
	return v * len(country) / (len(country) + len(domain))
}
