package data

import (
	"fmt"
	"math"
)

type Vector struct {
	X, Y int
}

func (v Vector) Add(other Vector) Vector {
	v.X += other.X
	v.Y += other.Y
	return v
}

func (v Vector) ManhattanDistance(to Vector) int {
	return int(math.Abs(float64(v.X-to.X))) + int(math.Abs(float64(v.Y-to.Y)))
}

func (v Vector) String() string {
	return fmt.Sprintf("%d,%d", v.X, v.Y)
}
