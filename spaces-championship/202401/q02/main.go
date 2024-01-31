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
	n, _ := strconv.Atoi(s.Text())
	var count int
	for i := 0; i < n; i++ {
		s.Scan()
		if check(s.Text()) {
			count++
		}
	}
	fmt.Fprintf(os.Stdout, "%d", count)
}

func check(s string) bool {
	user, pass, _ := strings.Cut(s, " ")
	if len(user) > len(pass) {
		return true
	}

	lastFound := 0
	for _, char := range user {
		idx := strings.IndexRune(pass[lastFound:], char)
		if idx == -1 {
			return true
		}
		lastFound = idx + lastFound + 1
	}

	return false
}
