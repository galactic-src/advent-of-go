package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func part1() {
	depths := strings.Split(input, "\n")
	prev := 0
	decreases := 0
	for _, depth := range depths {
		d, _ := strconv.Atoi(depth)
		if prev != 0 && d > prev {
			decreases++
		}
		prev = d
	}
	fmt.Printf("Part 1: %d\n", decreases)
}

//440

func depth(s string) int {
	d, _ := strconv.Atoi(s)
	return d
}

func part2() {
	depths := strings.Split(input, "\n")
	prev1 := 0
	prev2 := 0
	prev3 := 0
	decreases := 0
	for i:=0; i<len(depths); i++ {
		d := depth(depths[i])
		if prev1 != 0 && prev2 != 0 && prev3 != 0 && d > prev1 {
			decreases++
		}
		prev1 = prev2
		prev2 = prev3
		prev3 = d
	}
	fmt.Printf("Part 2: %d\n", decreases)
}

func main() {
	part1()
	part2()
}