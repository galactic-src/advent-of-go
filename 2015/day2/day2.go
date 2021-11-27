package main

import (
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Box struct {
	x, y, z int
}

func (b Box) paperRequired() int {
	return 3*b.x*b.y + 2*b.y*b.z + 2*b.x*b.z
}

func (b Box) ribbonRequired() int {
	return 2*(b.x+b.y) + b.x*b.y*b.z
}

func parseInput() []Box {
	lines := strings.Split(input, "\n")
	boxes := make([]Box, len(lines))
	for i := 0; i<len(lines); i++ {
		line := lines[i]
		dims := strings.Split(line, "x")
		x, err := strconv.Atoi(dims[0])
		if err != nil {
			log.Fatalf("failed to parse x")
		}
		y, err := strconv.Atoi(dims[1])
		if err != nil {
			log.Fatalf("failed to parse y")
		}
		z, err := strconv.Atoi(dims[2])
		if err != nil {
			log.Fatalf("failed to parse z")
		}

		sorted_dims := []int{x, y, z}
		sort.Ints(sorted_dims)

		boxes = append(boxes, Box{sorted_dims[0], sorted_dims[1], sorted_dims[2]})
	}

	return boxes
}

func part1() {
	boxes := parseInput()
	total := 0
	for i:=0; i< len(boxes); i++ {
		box := boxes[i]
		total += box.paperRequired()
	}

	fmt.Printf("Part 1: paper required = %d\n", total)
}

func part2() {
	boxes := parseInput()
	total := 0
	for i:=0; i< len(boxes); i++ {
		box := boxes[i]
		total += box.ribbonRequired()
	}

	fmt.Printf("Part 2: ribbon required = %d\n", total)
}

func main() {
	part1()
	part2()
}