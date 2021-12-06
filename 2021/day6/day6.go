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

type FishCounts [9]uint64

func (fishCounts *FishCounts) nextGen() {
	births := fishCounts[0]
	for i:=1; i<9; i++ {
		fishCounts[i-1] = fishCounts[i]
	}
	fishCounts[8] = births
	fishCounts[6] += births
}

func (fishCounts FishCounts) total() uint64 {
	var total uint64
	for i:=0; i<9; i++ {
		total += fishCounts[i]
	}
	return total
}

func parseInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		log.Fatalf("urgh")
	}
	return i
}

func parseInput() FishCounts {
	var fishCounts FishCounts

	for _, val := range strings.Split(input, ",") {
		fishCounts[parseInt(val)]++
	}

	return fishCounts
}

func part1() {
	fishCounts := parseInput()

	for i:=0; i<80; i++ {
		fishCounts.nextGen()
	}

	fmt.Printf("Part 1: %d\n", fishCounts.total())
}

func part2() {
	fishCounts := parseInput()

	for i:=0; i<256; i++ {
		fishCounts.nextGen()
	}

	fmt.Printf("Part 2: %d\n", fishCounts.total())
}

func main() {
	part1()
	part2()
}