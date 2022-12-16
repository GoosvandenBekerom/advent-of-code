package main

import (
	_ "embed"
	"fmt"
	"github.com/GoosvandenBekerom/advent-of-code/data"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
	"math"
	"regexp"
	"strings"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines, 2_000_000))
	fmt.Println(part2(lines, 4_000_000))
}

type sensorToBeacon map[data.Vector]data.Vector
type vectorMap map[data.Vector]struct{}

func part1(lines []string, checkY int) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	sensors, beacons := parseInput(lines)

	noBeacons := make(vectorMap)
	for sensor, closestBeacon := range sensors {
		distance := sensor.ManhattanDistance(closestBeacon)

		distanceToRow := int(math.Abs(float64(sensor.Y) - float64(checkY)))
		if distanceToRow > distance {
			continue
		}

		amount := distance - distanceToRow
		for x := sensor.X - amount; x <= sensor.X+amount; x++ {
			v := data.Vector{X: x, Y: checkY}
			if _, isBeacon := beacons[v]; isBeacon {
				continue
			}
			noBeacons[v] = struct{}{}
		}
	}

	return len(noBeacons)
}

func part2(lines []string, maxXY int) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	for x := 0; x < maxXY; x++ {
		for y := 0; y < maxXY; y++ {
			// todo: this would be way to brute forcy, find smarter way
		}
	}

	var beaconPosition data.Vector // TODO find remaining

	return beaconPosition.X*maxXY + beaconPosition.Y
}

var regex = regexp.MustCompile(`Sensor at x=([-\d]+), y=([-\d]+): closest beacon is at x=([-\d]+), y=([-\d]+)`)

func parseInput(lines []string) (sensors sensorToBeacon, beacons vectorMap) {
	sensors = make(sensorToBeacon)
	beacons = make(vectorMap)
	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)
		sensor := data.Vector{X: utils.ToInt(matches[1]), Y: utils.ToInt(matches[2])}
		beacon := data.Vector{X: utils.ToInt(matches[3]), Y: utils.ToInt(matches[4])}
		sensors[sensor] = beacon
		beacons[beacon] = struct{}{}
	}
	return sensors, beacons
}

func printMap(sensors sensorToBeacon, beacons vectorMap) {
	m := make(map[data.Vector]byte)
	for vector := range sensors {
		m[vector] = 'S'
	}
	for vector := range beacons {
		m[vector] = 'B'
	}
	var minx, maxx, miny, maxy int
	minx = math.MaxInt
	for position := range m {
		if position.X < minx {
			minx = position.X
		}
		if position.X > maxx {
			maxx = position.X
		}
		if position.Y < miny {
			miny = position.Y
		}
		if position.Y > maxy {
			maxy = position.Y
		}
	}

	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if v, ok := m[data.Vector{X: x, Y: y}]; ok {
				fmt.Printf("%c", v)
			} else {
				print(".")
			}
		}
		println()
	}
}
