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

type Grid [1000][1000]int

func part1() {
	lines := strings.Split(input, "\n")
	var grid Grid
	var cmd string
	for _, line := range lines {
		remainder := line
		if strings.HasPrefix(remainder, "toggle") {
			cmd = "toggle"
			remainder = line[len("toggle "):]
		} else if strings.HasPrefix(remainder, "turn on") {
			cmd = "turn on"
			remainder = line[len("turn on "):]
		} else if strings.HasPrefix(remainder, "turn off") {
			cmd = "turn off"
			remainder = line[len("turn off "):]
		} else {
			log.Fatalf("command %s", line)
		}

		coords := strings.Split(remainder, " through ")
		coord1 := strings.Split(coords[0], ",")
		coord2 := strings.Split(coords[1], ",")

		startX, _ := strconv.Atoi(coord1[0])
		startY, _ := strconv.Atoi(coord1[1])
		endX, _ := strconv.Atoi(coord2[0])
		endY, _ := strconv.Atoi(coord2[1])

		for x:=startX; x<=endX; x++ {
			for y:=startY; y<=endY; y++ {
				if cmd == "turn off" {
					grid[x][y] = 0
				} else if cmd == "turn on" {
					grid[x][y] = 1
				} else if cmd == "toggle" {
					if grid[x][y] == 1 {
						grid[x][y] = 0
					} else {
						grid[x][y] = 1
					}
				}
			}
		}
	}

	count := 0
	for x:=0; x<1000; x++ {
		for y:=0; y<1000; y++ {
			if grid[x][y] > 0 {
				count++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", count)
}

func part2() {
	lines := strings.Split(input, "\n")
	var grid Grid
	var cmd string
	for _, line := range lines {
		remainder := line
		if strings.HasPrefix(remainder, "toggle") {
			cmd = "toggle"
			remainder = line[len("toggle "):]
		} else if strings.HasPrefix(remainder, "turn on") {
			cmd = "turn on"
			remainder = line[len("turn on "):]
		} else if strings.HasPrefix(remainder, "turn off") {
			cmd = "turn off"
			remainder = line[len("turn off "):]
		} else {
			log.Fatalf("command %s", line)
		}

		coords := strings.Split(remainder, " through ")
		coord1 := strings.Split(coords[0], ",")
		coord2 := strings.Split(coords[1], ",")

		startX, _ := strconv.Atoi(coord1[0])
		startY, _ := strconv.Atoi(coord1[1])
		endX, _ := strconv.Atoi(coord2[0])
		endY, _ := strconv.Atoi(coord2[1])

		for x:=startX; x<=endX; x++ {
			for y:=startY; y<=endY; y++ {
				if cmd == "turn off" && grid[x][y] > 0 {
					grid[x][y]--
				} else if cmd == "turn on" {
					grid[x][y]++
				} else if cmd == "toggle" {
					grid[x][y]+=2
				}
			}
		}
	}

	total := 0
	for x:=0; x<1000; x++ {
		for y:=0; y<1000; y++ {
			total+=grid[x][y]
		}
	}

	fmt.Printf("Part 2: %d\n", total)
}

func main() {
	part1()
	part2()
}