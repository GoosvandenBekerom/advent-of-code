package data

import "fmt"

type Vector struct {
	X, Y int
}

func (v Vector) Add(other Vector) Vector {
	v.X += other.X
	v.Y += other.Y
	return v
}

func (v Vector) String() string {
	return fmt.Sprintf("%d-%d", v.X, v.Y)
}
