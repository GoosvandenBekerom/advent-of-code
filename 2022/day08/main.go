package main

import (
	_ "embed"
	"fmt"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
	"strings"
)

//go:embed input
var input string

type vector struct {
	x, y int
}

func main() {
	var grid [][]int
	for y, line := range strings.Split(input, "\n") {
		grid = append(grid, make([]int, len(line)))
		for x, value := range []byte(line) {
			grid[y][x] = utils.ToInt(string(value))
		}
	}

	fmt.Println(part1(grid))
	fmt.Println(part2(grid))
}

func part1(trees [][]int) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	visible := make(map[vector]struct{}, 0)
	rows, cols := len(trees), len(trees[0])

	// borders
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if x == 0 || y == 0 || x == rows-1 || y == cols-1 {
				visible[vector{x, y}] = struct{}{}
			}
		}
	}

	// left/right
	for x := 0; x < rows; x++ {
		highestLeft := trees[x][0]
		highestRight := trees[x][cols-1]

		for y := 0; y < cols; y++ {
			leftTree := trees[x][y]
			rightTree := trees[x][cols-1-y]

			if leftTree > highestLeft {
				highestLeft = leftTree
				visible[vector{x, y}] = struct{}{}
			}

			if rightTree > highestRight {
				highestRight = rightTree
				visible[vector{x, cols - 1 - y}] = struct{}{}
			}
		}
	}

	// up/down
	for y := 0; y < cols; y++ {
		highestTop := trees[0][y]
		highestBottom := trees[rows-1][y]

		for x := 0; x < rows; x++ {
			topTree := trees[x][y]
			bottomTree := trees[rows-1-x][y]

			if topTree > highestTop {
				highestTop = topTree
				visible[vector{x, y}] = struct{}{}
			}

			if bottomTree > highestBottom {
				highestBottom = bottomTree
				visible[vector{rows - 1 - x, y}] = struct{}{}
			}
		}
	}

	return len(visible)
}

func part2(trees [][]int) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	scores := make(map[vector]int, 0)
	rows, cols := len(trees), len(trees[0])

	for x := 1; x < rows; x++ {
		for y := 1; y < cols; y++ {
			pos := vector{x, y}
			scores[pos] = calculateScenicScore(pos, trees)
		}
	}

	var max int
	for _, v := range scores {
		if v > max {
			max = v
		}
	}

	return max
}

func calculateScenicScore(v vector, trees [][]int) int {
	var top int
	for i := v.x; i >= 0; i-- {
		if trees[v.x][v.y] <= trees[i][v.y] && i != v.x {
			top++
			break
		}

		if trees[v.x][v.y] > trees[i][v.y] {
			top++
		}
	}

	var bottom int
	for i := v.x; i < len(trees); i++ {
		if trees[v.x][v.y] <= trees[i][v.y] && i != v.x {
			bottom++
			break
		}

		if trees[v.x][v.y] > trees[i][v.y] {
			bottom++
		}
	}

	var left int
	for i := v.y; i >= 0; i-- {
		if trees[v.x][v.y] <= trees[v.x][i] && i != v.y {
			left++
			break
		}

		if trees[v.x][v.y] > trees[v.x][i] {
			left++
		}
	}

	var right int
	for i := v.y; i < len(trees[0]); i++ {
		if trees[v.x][v.y] <= trees[v.x][i] && i != v.y {
			right++
			break
		}

		if trees[v.x][v.y] > trees[v.x][i] {
			right++
		}
	}

	return top * bottom * right * left
}
