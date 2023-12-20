package math

import "golang.org/x/exp/constraints"

func Abs[T constraints.Signed](i T) T {
	if i < 0 {
		return i * -1
	}
	return i
}
