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
