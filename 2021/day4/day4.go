package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Card [5][5]int

func ( c Card) wins(called map[int]bool) bool {
	for row:=0; row<5; row++ {
		allFound := true
		for col:=0; col<5; col++ {
			if _, ok := called[c[row][col]]; !ok {
				allFound = false
			}
		}
		if allFound {
			return true
		}
	}

	// winning cols
	for col:=0; col<5; col++ {
		allFound := true
		for row:=0; row<5; row++ {
			if _, ok := called[c[row][col]]; !ok {
				allFound = false
			}
		}
		if allFound {
			return true
		}
	}

	return false
}

func parseInput() ([]int, []Card){
	removeDoubleSpaces := strings.Replace(input, "  ", " ", -1)
	removeLeadingSpaces := strings.Replace(removeDoubleSpaces, "\n ", "\n", -1)
	lines := strings.Split(removeLeadingSpaces, "\n")
	callsS := strings.Split(lines[0], ",")
	var calls []int
	for _, num := range callsS {
		i, _ := strconv.Atoi(num)
		calls = append(calls, i)
	}

	var cards []Card
	for ix:=2; ix<len(lines); ix+=6 {
		card := Card{[5]int{},[5]int{},[5]int{},[5]int{},[5]int{}}

		for i:=0; i<5; i++ {
			line := lines[ix+i]
			numsS := strings.Split(line, " ")
			for j:=0; j<5; j++ {
				k, _ := strconv.Atoi(numsS[j])
				card[i][j] = k
			}
		}

		cards = append(cards, card)
	}

	return calls, cards
}

func winningCard(calls []int, cards []Card) (Card, int, map[int]bool){
	called := make(map[int]bool)
	for _, call := range calls {
		called[call] = true
		for _, card := range cards {
			if card.wins(called) {
				return card, call, called
			}
		}
	}

	return Card{}, 0, make(map[int]bool)
}

func losingCard(calls []int, cards []Card) (Card, int, map[int]bool){
	remaining := cards
	called := make(map[int]bool)
	for _, call := range calls {
		called[call] = true

		if len(remaining) > 1 {
			var nextRemaining []Card
			for _, card := range remaining {
				if !card.wins(called) {
					nextRemaining = append(nextRemaining, card)
				}
			}
			remaining = nextRemaining
		} else {
			wins := remaining[0].wins(called)
			if wins {
				return remaining[0], call, called
			}
		}
	}

	return Card{}, 0, make(map[int]bool)
}

func part1() {
	calls, cards := parseInput()
	card, winningCall, winningCalls := winningCard(calls, cards)

	cardTotal := 0
	for col:=0; col<5; col++ {
		for row:=0; row<5; row++ {
			num := card[row][col]
			if _, marked := winningCalls[num]; !marked {
				cardTotal += num
			}
		}
	}

	fmt.Printf("Part 1: %d\n", winningCall * cardTotal)
}

//19740

func part2() {
	calls, cards := parseInput()
	card, losingCall, winningCalls := losingCard(calls, cards)

	cardTotal := 0
	for col:=0; col<5; col++ {
		for row:=0; row<5; row++ {
			num := card[row][col]
			if _, marked := winningCalls[num]; !marked {
				cardTotal += num
			}
		}
	}

	fmt.Printf("Part 2: %d\n", losingCall * cardTotal)
}

func main() {
	part1()
	part2()
}