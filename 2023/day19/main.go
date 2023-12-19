package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/GoosvandenBekerom/advent-of-code/utils"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

type outputType int

type part map[byte]int

type workflow struct {
	field    byte
	operator byte
	value    int
	next     string
}

func (w workflow) execute(p part) string {
	if w.field == 0 {
		return w.next
	}
	if w.operator == '>' {
		if p[w.field] > w.value {
			return w.next
		}
	} else if w.operator == '<' {
		if p[w.field] < w.value {
			return w.next
		}
	}
	return ""
}

func parse(lines []string) (workflowsByName map[string][]workflow, parts []part) {
	parsingWorkflows := true
	workflowsByName = make(map[string][]workflow)
	for _, line := range lines {
		if parsingWorkflows {
			if line == "" {
				parsingWorkflows = false
				continue
			}
			var workflows []workflow
			name, remainder, _ := strings.Cut(line, "{")
			for _, raw := range strings.Split(remainder[:len(remainder)-1], ",") {
				workflows = append(workflows, parseWorkflow(raw))
			}
			workflowsByName[name] = workflows
		} else {
			parts = append(parts, parsePart(line))
		}
	}
	return workflowsByName, parts
}

func parseWorkflow(raw string) workflow {
	if !strings.ContainsAny(raw, "<>") {
		return workflow{next: raw}
	}
	colonIndex := strings.IndexByte(raw, ':')
	return workflow{
		field:    raw[0],
		operator: raw[1],
		value:    utils.ToInt(raw[2:colonIndex]),
		next:     raw[colonIndex+1:],
	}
}

func parsePart(raw string) part {
	p := make(part)
	for _, field := range strings.Split(raw[1:len(raw)-1], ",") {
		p[field[0]] = utils.ToInt(field[2:])
	}
	return p
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	workflows, parts := parse(lines)
	//fmt.Printf("workflows: %v\nparts: %v\n", workflows, parts)
	var accepted []part
	start := workflows["in"]
	for _, p := range parts {
		p := p // to avoid getting bitten by yet another loop variable bug
		current := start
	workflowLoop:
		for {
		currentLoop:
			for _, wf := range current {
				result := wf.execute(p)
				switch result {
				case "A":
					accepted = append(accepted, p)
					break workflowLoop
				case "R":
					break workflowLoop
				case "":
					continue
				default:
					current = workflows[result]
					break currentLoop
				}
			}
		}
	}
	var sum int
	for _, p := range accepted {
		for _, i := range p {
			sum += i
		}
	}
	return sum
}

func countDistinctCombinations(workflowMap map[string][]workflow, name string, bounds map[byte][2]int) int {
	if name == "A" {
		// Return product of remaining ranges.
		product := 1
		for _, bound := range bounds {
			product *= bound[1] - bound[0] + 1 // (inclusive)
		}
		return product
	}
	if name == "R" {
		return 0
	}

	workflows := workflowMap[name]
	sum := 0
	for _, wf := range workflows[:len(workflows)-1] {
		// Copy the bounds map for the true side of the condition.
		// The argument bounds will be used for the false side.
		trueBounds := make(map[byte][2]int)
		for k, v := range bounds {
			trueBounds[k] = v
		}

		if wf.operator == '<' {
			// True case: [lower, wf.value - 1]
			trueBounds[wf.field] = [2]int{bounds[wf.field][0], wf.value - 1}
			sum += countDistinctCombinations(workflowMap, wf.next, trueBounds)

			// False case: [wf.value, upper]
			bounds[wf.field] = [2]int{wf.value, bounds[wf.field][1]}
		} else {
			// True case: [wf.value + 1, upper]
			trueBounds[wf.field] = [2]int{wf.value + 1, bounds[wf.field][1]}
			sum += countDistinctCombinations(workflowMap, wf.next, trueBounds)

			// False case: [lower, wf.value]
			bounds[wf.field] = [2]int{bounds[wf.field][0], wf.value}
		}
	}

	// Finally, count the distinct combinations for the fallback workflow using leftover 'false' bounds.
	return sum + countDistinctCombinations(workflowMap, workflows[len(workflows)-1].next, bounds)
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	workflows, _ := parse(lines)
	return countDistinctCombinations(workflows, "in", map[byte][2]int{
		'x': {1, 4000},
		'm': {1, 4000},
		'a': {1, 4000},
		's': {1, 4000},
	})
}
