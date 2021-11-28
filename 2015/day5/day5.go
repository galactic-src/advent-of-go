package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func countVowels(s string) int {
	count := 0
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			count++
		}
	}
	return count
}

func gotRepeatedPair(s string) bool {
	for i := 0; i<len(s)-3; i++ {
		for j:=i+2;j<len(s)-1; j++ {
			if s[i] == s[j] && s[i+1] == s[j+1] {
				return true
			}
		}
	}
	return false
}

func gotLetterPair(s string) bool {
	for i := 0; i<len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func gotTwinLetter(s string) bool {
	for i := 0; i<len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func gotForbiddenPair(s string) bool {
	for i := 0; i<len(s)-1; i++ {
		if (s[i] == 'a' && s[i+1] == 'b') || (s[i] == 'c' && s[i+1] == 'd') || (s[i] == 'p' && s[i+1] == 'q') || (s[i] == 'x' && s[i+1] == 'y') {
			return true
		}
	}
	return false
}

func isNice1(s string) bool {
	return countVowels(s) >= 3 && gotLetterPair(s) && !gotForbiddenPair(s)
}

func isNice2(s string) bool {
	return gotRepeatedPair(s) && gotTwinLetter(s)
}

func part1() {
	lines := strings.Split(input, "\n")
	niceCount := 0
	for _, line := range lines {
		if isNice1(line) {
			niceCount++
		}
	}
	fmt.Printf("Part 1: %d\n", niceCount)
}

func part2() {
	lines := strings.Split(input, "\n")
	niceCount := 0
	for _, line := range lines {
		if isNice2(line) {
			niceCount++
		}
	}
	fmt.Printf("Part 2: %d\n", niceCount)
}

func main() {
	part1()
	part2()
}