package utils

import (
	"strconv"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ToInt(s string) int {
	v, err := strconv.Atoi(s)
	Check(err)
	return v
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Map[I, O any](in []I, f func(item I) O) (out []O) {
	for _, item := range in {
		out = append(out, f(item))
	}
	return out
}
