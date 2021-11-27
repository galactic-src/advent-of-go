package main

import (
	"fmt"
	"log"
)

import _ "embed"

//go:embed input.txt
var input string

func part1() {
	floor := 0

	for _, runeValue := range input {
		switch runeValue {
		case '(':
			floor += 1
		case ')':
			floor -= 1
		default:
			log.Fatalf("unexpected char \"%s\"", string(runeValue))
		}
	}

	fmt.Printf("Part 1: floor = %d\n", floor)
}

func part2() {
	floor := 0

	for ix, runeValue := range input {
		switch runeValue {
		case '(':
			floor += 1
		case ')':
			floor -= 1
		default:
			log.Fatalf("unexpected char \"%s\"", string(runeValue))
		}

		if floor == -1 {
			fmt.Printf("Part 2: instruction = %d\n", ix+1)
			break
		}
	}

}

func main() {
	part1()
	part2()
}
