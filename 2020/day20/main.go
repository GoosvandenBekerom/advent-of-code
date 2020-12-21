package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	println(part1())
	println(part2())
}

func part1() int {
	images := input()
	// borders maps image id's to their borders
	borders := make(map[int][]uint16)
	// pairedBorders maps a border (10 bit number) to a list of images id's that have that border
	pairedBorders := make(map[uint16][]int)
	for id, img := range images {
		borders[id] = getBorders(img)
		for _, border := range borders[id] {
			if _, ok := pairedBorders[border]; ok {
				pairedBorders[border] = append(pairedBorders[border], id)
			} else {
				pairedBorders[border] = []int{id}
			}
		}
	}

	for border, ids := range pairedBorders {
		if len(ids) != 2 {
			delete(pairedBorders, border)
		}
	}

	neighbors := make(map[int]map[int]struct{})

	for border, ids := range pairedBorders {
		fmt.Printf("%d: %v\n", border, ids)
		i1 := ids[0]
		i2 := ids[1]
		found := false
		for _, b := range borders[i1] {
			if border == b {
				found = true
				break
			}
		}
		if !found {
			panic("no match found for border 1")
		}
		o2 := -1
		for side, b := range borders[ids[1]] {
			if border == b {
				o2 = side
				break
			}
		}
		if o2 < 0 {
			panic("no match found for orientation 2")
		}

		if _, ok := neighbors[i1]; !ok {
			neighbors[i1] = map[int]struct{}{}
		}
		if _, ok := neighbors[i2]; !ok {
			neighbors[i2] = map[int]struct{}{}
		}
		neighbors[i1][i2] = struct{}{}
		neighbors[i2][i1] = struct{}{}
	}

	fmt.Printf("neighbors %v\n", neighbors)

	var corners []int
	// a corner should only have 2 neighbors
	for id, neighbors := range neighbors {
		if len(neighbors) == 2 {
			corners = append(corners, id)
		}
	}

	fmt.Printf("corners %v\n", corners)

	if len(corners) > 4 {
		panic("found more then 4 corners")
	}

	solution := 1
	for _, corner := range corners {
		solution *= corner
	}
	return solution
}

func part2() (sum int) {
	return
}

// a border is a 10 bit number in which every bit represents if a pixel is . or #
func getBorders(img image) (borders []uint16) {
	borders = make([]uint16, 8)
	var s string
	// top + reversed
	for i := 0; i < 10; i++ {
		s += toBitString(img[0][i])
	}
	borders = append(borders, toBorder(s))
	s = string(reverse([]byte(s)))
	borders = append(borders, toBorder(s))
	s = ""
	// right + reversed
	for i := 0; i < 10; i++ {
		s += toBitString(img[i][9])
	}
	borders = append(borders, toBorder(s))
	s = string(reverse([]byte(s)))
	borders = append(borders, toBorder(s))
	s = ""
	// bottom + reversed
	for i := 0; i < 10; i++ {
		s += toBitString(img[9][i])
	}
	borders = append(borders, toBorder(s))
	s = string(reverse([]byte(s)))
	borders = append(borders, toBorder(s))
	s = ""
	// left + reversed
	for i := 0; i < 10; i++ {
		s += toBitString(img[i][0])
	}
	borders = append(borders, toBorder(s))
	s = string(reverse([]byte(s)))
	borders = append(borders, toBorder(s))
	return
}

func toBitString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func reverse(s []byte) []byte {
	a := make([]byte, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func toBorder(bits string) uint16 {
	i, err := strconv.ParseUint(bits, 2, 10)
	check(err)
	return uint16(i)
}

// ----------------------------------------
// domain
// ----------------------------------------

type image [][]bool

func (img image) print() {
	for _, row := range img {
		for _, col := range row {
			c := "."
			if col {
				c = "#"
			}
			print(c, " ")
		}
		println()
	}
}

// ----------------------------------------
// utils
// ----------------------------------------

func input() (images map[int]image) {
	box := packr.New("day20", "./2020/day20")
	s, err := box.FindString("input")
	check(err)
	matcher := regexp.MustCompile(`\d+`)
	images = make(map[int]image)
	for _, raw := range strings.Split(strings.TrimSuffix(s, "\n"), "\n\n") {
		rows := strings.Split(raw, "\n")
		idStr := matcher.FindString(rows[0])
		id, err := strconv.Atoi(idStr)
		check(err)
		images[id] = make(image, len(rows)-1)
		for y, columns := range rows[1:] {
			images[id][y] = make([]bool, len(columns))
			for x, c := range columns {
				images[id][y][x] = c == '#'
			}
		}
	}
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
