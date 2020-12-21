package main

import (
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	println(part1())
	println(part2())
}

type stringSet map[string]string

func (s stringSet) add(str string) {
	s[str] = ""
}

func (s stringSet) contains(str string) bool {
	_, ok := s[str]
	return ok
}

func (s stringSet) intersect(other stringSet) (intersection stringSet) {
	intersection = make(stringSet)
	for s1 := range s {
		for s2 := range other {
			if s1 == s2 {
				intersection.add(s1)
			}
		}
	}
	return
}

func (s stringSet) keys() []string {
	ks := make([]string, len(s))
	var i int
	for k := range s {
		ks[i] = k
		i++
	}
	return ks
}

func (s stringSet) first() string {
	for k := range s {
		return k
	}
	panic("faka lege set")
}

type food struct {
	allergens   stringSet
	ingredients stringSet
}

func (f food) containsAllergen(want string) bool {
	for allergen := range f.allergens {
		if allergen == want {
			return true
		}
	}
	return false
}

func part1() int {
	counts := make(map[string]int)
	allAllergens := make(stringSet)
	var foods []food

	matcher := regexp.MustCompile(`(.+)+ \(contains ([a-z, ]+)\)`)
	for line := range input() {
		matches := matcher.FindStringSubmatch(line)
		ingredients := strings.Split(matches[1], " ")
		for _, word := range ingredients {
			counts[word]++
		}
		allergens := strings.Split(strings.ReplaceAll(matches[2], " ", ""), ",")
		a := stringSet{}
		for _, allergen := range allergens {
			a.add(allergen)
			allAllergens.add(allergen)
		}
		i := stringSet{}
		for _, ingredient := range ingredients {
			i.add(ingredient)
		}
		foods = append(foods, food{
			allergens:   a,
			ingredients: i,
		})
	}

	found := make(map[string]string)

	for len(found) != len(allAllergens) {
		for allergen := range allAllergens {
			if _, ok := found[allergen]; ok {
				continue
			}
			var containing []stringSet
			for _, f := range foods {
				if f.containsAllergen(allergen) {
					fc := make(stringSet)
					for ingredient := range f.ingredients {
						contains := false
						for _, v := range found {
							if ingredient == v {
								contains = true
								break
							}
						}
						if !contains {
							fc.add(ingredient)
						}
					}
					containing = append(containing, fc)
				}
			}
			if containing == nil {
				panic("no food contains " + allergen)
			}
			c := containing[0]
			for _, c2 := range containing {
				c = c.intersect(c2)
			}

			if len(c) == 1 {
				found[allergen] = c.first()
			}
		}
	}

	sum := 0

	for ingredient, count := range counts {
		exists := false
		for _, i2 := range found {
			if ingredient == i2 {
				exists = true
				break
			}
		}
		if !exists {
			sum += count
		}
	}

	return sum
}

func part2() string {
	counts := make(map[string]int)
	allAllergens := make(stringSet)
	var foods []food

	matcher := regexp.MustCompile(`(.+)+ \(contains ([a-z, ]+)\)`)
	for line := range input() {
		matches := matcher.FindStringSubmatch(line)
		ingredients := strings.Split(matches[1], " ")
		for _, word := range ingredients {
			counts[word]++
		}
		allergens := strings.Split(strings.ReplaceAll(matches[2], " ", ""), ",")
		a := stringSet{}
		for _, allergen := range allergens {
			a.add(allergen)
			allAllergens.add(allergen)
		}
		i := stringSet{}
		for _, ingredient := range ingredients {
			i.add(ingredient)
		}
		foods = append(foods, food{
			allergens:   a,
			ingredients: i,
		})
	}

	found := make(stringSet)

	for len(found) != len(allAllergens) {
		for allergen := range allAllergens {
			if _, ok := found[allergen]; ok {
				continue
			}
			var containing []stringSet
			for _, f := range foods {
				if f.containsAllergen(allergen) {
					fc := make(stringSet)
					for ingredient := range f.ingredients {
						contains := false
						for _, v := range found {
							if ingredient == v {
								contains = true
								break
							}
						}
						if !contains {
							fc.add(ingredient)
						}
					}
					containing = append(containing, fc)
				}
			}
			if containing == nil {
				panic("no food contains " + allergen)
			}
			c := containing[0]
			for _, c2 := range containing {
				c = c.intersect(c2)
			}

			if len(c) == 1 {
				found[allergen] = c.first()
			}
		}
	}

	keys := found.keys()
	sort.Strings(keys)

	result := ""
	for _, k := range keys {
		result += found[k] + ","
	}
	return strings.TrimSuffix(result, ",")
}

// ----------------------------------------
// utils
// ----------------------------------------

func input() (raw chan string) {
	box := packr.New("day21", "./2020/day21")
	s, err := box.FindString("input")
	check(err)
	raw = make(chan string, 1)
	go func() {
		defer close(raw)
		for _, line := range strings.Split(strings.TrimSuffix(s, "\n"), "\n") {
			raw <- line
		}
	}()
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
