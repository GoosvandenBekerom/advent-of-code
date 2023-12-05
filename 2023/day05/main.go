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
	return findLowestSeedLocation(lines, func(seeds []int) []int {
		return seeds
	})
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	return findLowestSeedLocation(lines, func(seedRanges []int) (seeds []int) {
		for i := 0; i < len(seedRanges); i += 2 {
			for num := seedRanges[i]; num < seedRanges[i]+seedRanges[i+1]; num++ {
				seeds = append(seeds, num)
			}
		}
		return seeds
	})
}

func findLowestSeedLocation(lines []string, seedsMapper func([]int) []int) int {
	var seeds []int
	ranges := map[string][]numberRange{
		"seedToSoil":            {},
		"soilToFertilizer":      {},
		"fertilizerToWater":     {},
		"waterToLight":          {},
		"lightToTemperature":    {},
		"temperatureToHumidity": {},
		"humidityToLocation":    {},
	}

	var currentRange string
	for _, line := range lines {
		left, right, _ := strings.Cut(line, ":")
		switch left {
		case "":
			continue
		case "seeds":
			seeds = seedsMapper(utils.Map(strings.Split(strings.TrimSpace(right), " "), func(item string) int {
				return utils.ToInt(item)
			}))
		case "seed-to-soil map":
			currentRange = "seedToSoil"
		case "soil-to-fertilizer map":
			currentRange = "soilToFertilizer"
		case "fertilizer-to-water map":
			currentRange = "fertilizerToWater"
		case "water-to-light map":
			currentRange = "waterToLight"
		case "light-to-temperature map":
			currentRange = "lightToTemperature"
		case "temperature-to-humidity map":
			currentRange = "temperatureToHumidity"
		case "humidity-to-location map":
			currentRange = "humidityToLocation"
		default:
			ranges[currentRange] = append(ranges[currentRange], toNumberRange(line))
		}
	}

	lowest := math.MaxInt
	for _, seed := range seeds {
		soil := mapToRange(seed, ranges["seedToSoil"])
		fertilizer := mapToRange(soil, ranges["soilToFertilizer"])
		water := mapToRange(fertilizer, ranges["fertilizerToWater"])
		light := mapToRange(water, ranges["waterToLight"])
		temperature := mapToRange(light, ranges["lightToTemperature"])
		humidity := mapToRange(temperature, ranges["temperatureToHumidity"])
		location := mapToRange(humidity, ranges["humidityToLocation"])
		if location < lowest {
			println("new lowest, seed", seed, "location", location)
			lowest = location
		}
	}

	return lowest
}

type numberRange struct {
	destStart, sourceStart, length int
}

func (r numberRange) contains(target int) bool {
	return target >= r.sourceStart && target < r.sourceStart+r.length
}

func toNumberRange(line string) numberRange {
	numbers := strings.Split(line, " ")
	return numberRange{
		destStart:   utils.ToInt(numbers[0]),
		sourceStart: utils.ToInt(numbers[1]),
		length:      utils.ToInt(numbers[2]),
	}
}

func mapToRange(target int, ranges []numberRange) int {
	for _, r := range ranges {
		if r.contains(target) {
			return r.destStart + (target - r.sourceStart)
		}
	}
	return target
}
