package directions

import "github.com/GoosvandenBekerom/advent-of-code/data"

type Direction struct {
	data.Vector
}

var (
	Up        = Direction{data.Vector{X: 0, Y: -1}}
	Down      = Direction{data.Vector{X: 0, Y: 1}}
	Left      = Direction{data.Vector{X: -1, Y: 0}}
	Right     = Direction{data.Vector{X: 1, Y: 0}}
	UpLeft    = Direction{Up.Add(Left.Vector)}
	UpRight   = Direction{Up.Add(Right.Vector)}
	DownLeft  = Direction{Down.Add(Left.Vector)}
	DownRight = Direction{Down.Add(Right.Vector)}

	All           = []Direction{Up, Down, Left, Right, UpLeft, UpRight, DownLeft, DownRight}
	AllOrthogonal = []Direction{Up, Down, Left, Right}
)

func (d Direction) Opposite() Direction {
	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	case UpLeft:
		return DownRight
	case DownRight:
		return UpLeft
	case UpRight:
		return DownLeft
	case DownLeft:
		return UpRight
	default:
		panic("unknown direction")
	}
}
