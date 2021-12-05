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

type Point struct {
	x, y int
}

type Line struct {
	start, end Point
}

func parseInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		log.Fatalf("urgh")
	}
	return i
}

func parsePoint(s string) Point {
	split := strings.Split(s, ",")
	return Point{parseInt(split[0]), parseInt(split[1])}
}

func parseInput(onlyHV bool) []Line{
	lines := strings.Split(input, "\n")

	var points []Line

	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		p1 := parsePoint(coords[0])
		p2 := parsePoint(coords[1])

		if onlyHV {
			if p1.x == p2.x || p1.y == p2.y {
				points = append(points, Line{p1, p2})
			}
		} else {
			points = append(points, Line{p1, p2})
		}
	}

	return points
}

func minInt(ints ...int) int {
	if len(ints) == 0 {
		log.Fatalf("no ints for min")
	}

	min := ints[0]
	for _, i := range ints {
		if min > i {
			min = i
		}
	}
	return min
}

func maxInt(ints ...int) int {
	if len(ints) == 0 {
		log.Fatalf("no ints for max")
	}

	max := ints[0]
	for _, i := range ints {
		if max < i {
			max = i
		}
	}
	return max
}

func part1() {
	grid := make(map[Point]int)

	lines := parseInput(true)

	// figure out a big enough grid
	/*var minX, maxX, minY, maxY int
	for _, line := range lines {
		minX = minInt(minX, line.start.x, line.end.x)
		minY = minInt(minY, line.start.y, line.end.y)
		maxX = maxInt(maxX, line.start.x, line.end.x)
		maxY = maxInt(maxY, line.start.y, line.end.y)
	}*/

	for _, line := range lines {
		if line.start.x == line.end.x {
			for y:=minInt(line.start.y, line.end.y); y<=maxInt(line.start.y, line.end.y); y++ {
				p := Point{line.start.x, y}
				grid[p]++
			}
		} else {
			for x:=minInt(line.start.x, line.end.x); x<=maxInt(line.start.x, line.end.x); x++ {
				p := Point{x, line.start.y}
				grid[p]++
			}
		}

	}

	twoOrMore := 0
	for _, count := range grid {
		if count >= 2 {
			twoOrMore++
		}
	}

	fmt.Printf("Part 1: %d\n", twoOrMore)
}

func part2() {
	grid := make(map[Point]int)

	lines := parseInput(false)

	for _, line := range lines {
		x := line.start.x
		y := line.start.y
		for {
			p := Point{x, y}
			grid[p]++

			if x == line.end.x && y == line.end.y {
				break
			}

			if line.start.x < line.end.x {
				x++
			} else if line.start.x > line.end.x {
				x--
			}
			if line.start.y < line.end.y {
				y++
			} else if line.start.y > line.end.y {
				y--
			}
		}
	}

	twoOrMore := 0
	for _, count := range grid {
		if count >= 2 {
			twoOrMore++
		}
	}

	fmt.Printf("Part 2: %d\n", twoOrMore)
}

func main() {
	part1()
	part2()
}