package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func part1() {
	lines := strings.Split(input, "\n")
	x, y := 0, 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		distance, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "forward": x += distance
		case "up": y -= distance
		case "down": y += distance
		default: log.Fatalf("urgh")
		}
	}

	fmt.Printf("Part 1: %d\n", x * y)
}

func part2() {
	lines := strings.Split(input, "\n")
	x, y, aim := 0, 0, 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		distance, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "forward":
			x += distance
			y += aim * distance
		case "up": aim -= distance
		case "down": aim += distance
		default: log.Fatalf("urgh")
		}
	}

	fmt.Printf("Part 2: %d\n", x * y)
}

func main() {
	part1()
	part2()
}