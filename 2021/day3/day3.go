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
	lines := strings.Split(input, "\n")
	gammaS, epsilonS := mostCommonString(lines)

	gamma, _ := strconv.ParseInt(gammaS, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonS, 2, 64)

	fmt.Printf("Part 1: %d\n", gamma * epsilon)
}

func part2() {
	lines := strings.Split(input, "\n")
	var oxygen int64
	for candidates, ix := lines, 0; len(candidates) > 1; ix++ {
		var filtered []string
		most, _ := mostCommonString(candidates)
		for _, line := range candidates {
			if line[ix] == most[ix] {
				filtered = append(filtered, line)
			}
		}
		if len(filtered) == 1 {
			oxygen, _ = strconv.ParseInt(filtered[0], 2, 64)
			break
		}
		candidates = filtered
	}

	var co2 int64
	for candidates2, ix := lines, 0; len(candidates2) > 1; ix++ {
		var filtered []string
		_, least := mostCommonString(candidates2)
		for _, line := range candidates2 {
			if line[ix] == least[ix] {
				filtered = append(filtered, line)
			}
		}
		if len(filtered) == 1 {
			co2, _ = strconv.ParseInt(filtered[0], 2, 64)
			break
		}
		candidates2 = filtered
	}

	fmt.Printf("Part 2: %d\n", oxygen * co2)
}

func mostCommonString(inputs []string) (string, string) {
	var bitTotals [12]int
	for _, line := range inputs {
		for ix, r := range line {
			if r == '1' {
				bitTotals[ix]++
			}
		}
	}

	var most = ""
	var least = ""
	for i:=0; i<len(bitTotals); i++ {
		if float32(bitTotals[i]) >= float32(len(inputs))/2 {
			most += "1"
			least += "0"
		} else {
			most += "0"
			least += "1"
		}
	}

	return most, least
}

func main() {
	part1()
	part2()
}