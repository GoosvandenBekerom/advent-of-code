package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	println(part1())
	println(part2())
}

func input() (players map[int]*deck) {
	box := packr.New("day22", "./2020/day22")
	s, err := box.FindString("input")
	check(err)
	players = make(map[int]*deck)
	for _, block := range strings.Split(strings.TrimSuffix(s, "\n"), "\n\n") {
		lines := strings.Split(block, "\n")
		id := toInt(string(lines[0][len(lines[0])-2]))
		players[id] = &deck{}
		for _, raw := range lines[1:] {
			players[id].enqueue(toInt(raw))
		}
	}
	return
}

type deck struct {
	cards []int
}

func (d *deck) len() int {
	return len(d.cards)
}

func (d *deck) enqueue(i int) {
	d.cards = append(d.cards, i)
}

func (d *deck) dequeue() int {
	i := d.cards[0]
	d.cards = d.cards[1:]
	return i
}

func (d *deck) calculate() (result int) {
	for i, value := range d.cards {
		result += (len(d.cards) - i) * value
	}
	return
}

func (d *deck) snapshot() (result []int) {
	result = make([]int, len(d.cards))
	for i, value := range d.cards {
		result[i] = value
	}
	return
}

func (d *deck) subset(amount int) (result *deck) {
	result = &deck{make([]int, amount)}
	for i := 0; i < amount; i++ {
		result.cards[i] = d.cards[i]
	}
	return
}

func part1() int {
	println("--- PART 1 ---------------------------------")
	players := input()

	fmt.Printf("amount of cards player 1 %v\n", players[1].len())
	fmt.Printf("amount of cards player 2 %v\n", players[2].len())

	round := 1
	var solution int
	for {
		fmt.Printf("player 1 %v\n", players[1].cards)
		fmt.Printf("player 2 %v\n", players[2].cards)
		if players[1].len() == 0 {
			println("player 2 wins")
			solution = players[2].calculate()
			break
		}
		if players[2].len() == 0 {
			println("player 1 wins")
			solution = players[1].calculate()
			break
		}
		p1 := players[1].dequeue()
		p2 := players[2].dequeue()
		if p1 > p2 {
			println("player 1 wins round", round)
			players[1].enqueue(p1)
			players[1].enqueue(p2)
		}
		if p2 > p1 {
			println("player 2 wins round", round)
			players[2].enqueue(p2)
			players[2].enqueue(p1)
		}
		round++
	}
	return solution
}

func part2() int {
	println("--- PART 2 ---------------------------------")
	players := input()
	winner := play(players)
	return players[winner].calculate()
}

func play(players map[int]*deck) int {
	var history1 [][]int
	var history2 [][]int
	for {
		fmt.Printf("player 1 %v\n", players[1].cards)
		fmt.Printf("player 2 %v\n", players[2].cards)

		if players[1].len() == 0 {
			println("player 2 wins game")
			return 2
		}
		if players[2].len() == 0 {
			println("player 1 wins game")
			return 1
		}

		if isRepeat(history1, players[1]) || isRepeat(history2, players[2]) {
			println("history repeats itself, player 1 wins game")
			return 1
		}

		history1 = append(history1, players[1].snapshot())
		history2 = append(history2, players[2].snapshot())

		p1 := players[1].dequeue()
		p2 := players[2].dequeue()

		var roundWinner int

		if p1 <= players[1].len() && p2 <= players[2].len() {
			println("start subgame")
			roundWinner = play(map[int]*deck{
				1: players[1].subset(p1),
				2: players[2].subset(p2),
			})
		} else {
			if p1 > p2 {
				roundWinner = 1
			}
			if p2 > p1 {
				roundWinner = 2
			}
		}

		if roundWinner == 1 {
			println("player 1 wins round")
			players[1].enqueue(p1)
			players[1].enqueue(p2)
		} else {
			println("player 2 wins round")
			players[2].enqueue(p2)
			players[2].enqueue(p1)
		}
	}
}

func isRepeat(history [][]int, d *deck) bool {
	if len(history) == 0 {
		return false
	}
	for _, snapshot := range history {
		if len(snapshot) != d.len() {
			continue
		}
		same := true
		for i, card := range d.cards {
			if snapshot[i] != card {
				same = false
				break
			}
		}
		if same {
			return true
		}
	}
	return false
}

// ----------------------------------------
// utils
// ----------------------------------------

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func toInt(s string) int {
	v, err := strconv.Atoi(s)
	check(err)
	return v
}
