package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	w, h, _ := strings.Cut(s.Text(), " ")
	width, _ := strconv.Atoi(w)
	height, _ := strconv.Atoi(h)

	grid := make(map[pos]byte)
	var cat, mouse pos
	var cheeses []pos
	for y := 0; y < height; y++ {
		s.Scan()
		for x, char := range s.Text() {
			p := pos{x, y}
			switch char {
			case '.':
				continue
			case 'C':
				cat = p
			case 'M':
				mouse = p
			case 'X':
				cheeses = append(cheeses, p)
			}
			grid[p] = byte(char)
		}
	}

	answer := "cat"
	for _, cheese := range cheeses {
		catD := findShortestPath(grid, cat, cheese, width, height)
		mouseD := findShortestPath(grid, mouse, cheese, width, height)
		fmt.Println("cat", catD)
		fmt.Println("mouse", mouseD)
		if mouseD <= catD {
			answer = "mouse"
			break
		}
	}

	fmt.Printf("%s", answer)
}

type pos struct {
	x, y int
}

func (v pos) Add(other pos) pos {
	v.x += other.x
	v.y += other.y
	return v
}

func (d pos) Opposite() pos {
	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	default:
		panic("unknown direction")
	}
}

type distKey struct {
	position  pos
	direction pos
}

type state struct {
	cost      int
	position  pos
	direction pos
}

func (s state) LessThan(other state) bool {
	return s.cost < other.cost
}

var (
	Up            = pos{x: 0, y: -1}
	Down          = pos{x: 0, y: 1}
	Left          = pos{x: -1, y: 0}
	Right         = pos{x: 1, y: 0}
	AllOrthogonal = []pos{Up, Down, Left, Right}
)

func findShortestPath(grid map[pos]byte, from, to pos, width, height int) int {
	dist := make(map[distKey]int)
	h := new(Heap[state])

	dist[distKey{position: from, direction: Right}] = 0
	dist[distKey{position: from, direction: Down}] = 0
	heap.Push(h, state{cost: 0, position: from, direction: Right})

	for {
		s := heap.Pop(h).(state)
		fmt.Printf("%v\n", s)
		if s.position == to {
			return s.cost
		}

		k := distKey{position: s.position, direction: s.direction}
		if v, exists := dist[k]; exists && s.cost > v {
			continue
		}

		for _, dir := range AllOrthogonal {
			if dir == s.direction.Opposite() {
				continue
			}
			pos := s.position.Add(dir)
			if pos.x < 0 || pos.y < 0 || pos.x > width || pos.y > height {
				continue
			}

			cost := 1
			if grid[pos] == '#' {
				cost = 100
			}
			next := state{cost: s.cost + cost, position: pos, direction: dir}
			nextKey := distKey{position: pos, direction: dir}

			oldCost, exists := dist[nextKey]
			if !exists || next.cost < oldCost {
				heap.Push(h, next)
				dist[nextKey] = next.cost
			}
		}
	}
}

type LessThanComparer[T any] interface {
	LessThan(other T) bool
}

type Heap[T LessThanComparer[T]] []T

func (h Heap[T]) Len() int           { return len(h) }
func (h Heap[T]) Less(i, j int) bool { return h[i].LessThan(h[j]) }
func (h Heap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push uses a pointer receiver because it modifies
// the slice's length, not just its contents.
func (h *Heap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

// Pop uses a pointer receiver because it modifies
// the slice's length, not just its contents.
func (h *Heap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
