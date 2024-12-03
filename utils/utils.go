package utils

import (
	"fmt"
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

func All[T any](in []T, f func(T) bool) bool {
	for _, v := range in {
		if !f(v) {
			return false
		}
	}
	return true
}

func ReverseSlice[S any](s []S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Print(x any) {
	fmt.Printf("%#v\n", x)
}
