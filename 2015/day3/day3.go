package main

import (
	_ "embed"
	"fmt"
	"log"
)

//go:embed input.txt
var input string

type Point struct {
	x, y int
}

func part1() {
	var curr_x, curr_y int
	tallies := make(map[Point]int)
	tallies[Point{0,0}] = 1

	for _, runeValue := range input {
		switch runeValue {
		case '>':
			curr_x++
		case '<':
			curr_x--
		case '^':
			curr_y++
		case 'v':
			curr_y--
		default:
			log.Fatalf("unexpected char \"%s\"", string(runeValue))
		}

		curr_p := Point{curr_x, curr_y}
		tallies[curr_p] += 1
	}

	fmt.Printf("Part 1: %d\n", len(tallies))
}

func part2() {
	santa := Point{0,0}
	robo_santa := Point{0,0}

	tallies := make(map[Point]int)
	tallies[Point{0,0}] = 2

	curr_santa := &santa

	for _, runeValue := range input {
		switch runeValue {
		case '>':
			curr_santa.x++
		case '<':
			curr_santa.x--
		case '^':
			curr_santa.y++
		case 'v':
			curr_santa.y--
		default:
			log.Fatalf("unexpected char \"%s\"", string(runeValue))
		}

		tallies[*curr_santa] += 1

		if curr_santa == &santa {
			curr_santa = &robo_santa
		} else {
			curr_santa = &santa
		}
	}

	fmt.Printf("Part 2: %d\n", len(tallies))
}

func main() {
	part1()
	part2()
}