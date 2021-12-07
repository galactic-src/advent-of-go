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

func locations() []int {
	parts := strings.Split(input, ",")
	var locations []int
	for _, part := range parts {
		partI := parseInt(part)
		locations = append(locations, partI)
	}
	return locations
}

func parseInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		log.Fatalf("urgh")
	}
	return i
}

func bestFuel (locations []int, fuelCost func([]int, int) uint64) uint64 {
	var best = fuelCost2(locations, 0)
	for i:=1; i<1000; i++ {
		fuel := fuelCost(locations, i)
		if best > fuel {
			best = fuel
		}
	}
	return best
}

func fuelCost(locations []int, alignAt int) uint64 {
	var total uint64 = 0
	for _, location := range locations {
		if location > alignAt {
			total += uint64(location - alignAt)
		} else {
			total += uint64(alignAt - location)
		}
	}

	return total
}

func fuelCost2(locations []int, alignAt int) uint64 {
	var total uint64 = 0

	for _, location := range locations {
		if location > alignAt {
			total += (uint64(location - alignAt) * uint64(location - alignAt + 1)) / 2
		} else {
			total += (uint64(alignAt - location) * uint64(alignAt - location + 1)) / 2
		}
	}
	return total
}

func part1() {
	fmt.Printf("Part 1: %d\n", bestFuel(locations(), fuelCost))
}

func part2() {
	fmt.Printf("Part 2: %d\n", bestFuel(locations(), fuelCost2))
}

func main() {
	part1()
	part2()
}