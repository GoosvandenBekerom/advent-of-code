package main

import (
	_ "embed"
	"fmt"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
	"math"
	"strings"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	var seeds []int
	seedToSoil := make(map[int]int)
	soilToFertilizer := make(map[int]int)
	fertilizerToWater := make(map[int]int)
	waterToLight := make(map[int]int)
	lightToTemperature := make(map[int]int)
	temperatureToHumidity := make(map[int]int)
	humidityToLocation := make(map[int]int)

	var currentSection *map[int]int
	for _, line := range lines {
		left, right, _ := strings.Cut(line, ":")
		switch left {
		case "":
			currentSection = nil
			continue
		case "seeds":
			seeds = utils.Map(strings.Split(strings.TrimSpace(right), " "), func(item string) int {
				return utils.ToInt(item)
			})
		case "seed-to-soil map":
			currentSection = &seedToSoil
		case "soil-to-fertilizer map":
			currentSection = &soilToFertilizer
		case "fertilizer-to-water map":
			currentSection = &fertilizerToWater
		case "water-to-light map":
			currentSection = &waterToLight
		case "light-to-temperature map":
			currentSection = &lightToTemperature
		case "temperature-to-humidity map":
			currentSection = &temperatureToHumidity
		case "humidity-to-location map":
			currentSection = &humidityToLocation
		default:
			numRange := toNumberRange(line)
			section := *currentSection
			for i := 0; i < numRange.length; i++ {
				section[numRange.sourceRangeStart+i] = numRange.destinationRangeStart + i
			}
		}
	}

	lowest := math.MaxInt
	for _, seed := range seeds {
		soil, ok := seedToSoil[seed]
		if !ok {
			soil = seed
		}
		fertilizer, ok := soilToFertilizer[soil]
		if !ok {
			fertilizer = soil
		}
		water, ok := fertilizerToWater[fertilizer]
		if !ok {
			water = fertilizer
		}
		light, ok := waterToLight[water]
		if !ok {
			light = water
		}
		temp, ok := lightToTemperature[light]
		if !ok {
			temp = light
		}
		humidity, ok := temperatureToHumidity[temp]
		if !ok {
			humidity = temp
		}
		location, ok := humidityToLocation[humidity]
		if !ok {
			location = humidity
		}
		//println(seed, "maps to location", location)
		if location < lowest {
			lowest = location
		}
	}

	return lowest
}

type numberRange struct {
	destinationRangeStart, sourceRangeStart, length int
}

func toNumberRange(line string) numberRange {
	numbers := strings.Split(line, " ")
	return numberRange{
		destinationRangeStart: utils.ToInt(numbers[0]),
		sourceRangeStart:      utils.ToInt(numbers[1]),
		length:                utils.ToInt(numbers[2]),
	}
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	return -1
}
