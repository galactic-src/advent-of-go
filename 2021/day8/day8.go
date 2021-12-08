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

type LCD struct {
	config map[string]int
	value [4]string
}

func parseInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		log.Fatalf("urgh")
	}
	return i
}

func (lcd LCD) intValue() int {
	//fmt.Println(lcd.config, " -> ", lcd.value)
	intS := fmt.Sprintf("%d%d%d%d", lcd.config[lcd.value[0]], lcd.config[lcd.value[1]], lcd.config[lcd.value[2]], lcd.config[lcd.value[3]])
	//println(intS)
	return parseInt(intS)
}

func invert(s string) string {
	result := ""
	for r:='a'; r<'h'; r++ {
		found := false
		for i:=0; i<len(s); i++ {
			if r == int32(s[i]) {
				found = true
				break
			}
		}

		if !found {
			result += string(r)
		}
	}
	return result
}

func canonical(s string) string {
	result := ""
	for r:='a'; r<'h'; r++ {
		found := false
		for i:=0; i<len(s); i++ {
			if r == int32(s[i]) {
				found = true
				break
			}
		}

		if found {
			result += string(r)
		}
	}
	return result
}

func parseConfig(configs []string) map[string]int {
	byLength := make(map[int] []string)
	for _, entry := range configs {
		byLength[len(entry)] = append(byLength[len(entry)], canonical(entry))
	}

	configRev := make(map[int]string)

	configRev[1] = byLength[2][0]
	configRev[4] = byLength[4][0]
	configRev[7] = byLength[3][0]
	configRev[8] = byLength[7][0]

	// 2, 3, 5
	for _, entry := range byLength[5] {
		inverted := invert(entry)
		//println(entry, " -> ", inverted)
		//fmt.Println(configRev)

		// 2 = acdeg => b,f
		// 4 = bcdf
		mightBeTwo := true
		for i:=0; i<len(inverted); i++ {
			found := false
			for j:=0; j<4; j++ {
				if inverted[i] == configRev[4][j] {
					found = true
				}
			}
			if !found {
				mightBeTwo = false
				break
			}
		}
		if mightBeTwo {
			configRev[2] = entry
			continue
		}

		mightBeThree := true
		for i:=0; i<2; i++ {
			found := false
			for j:=0; j<len(entry); j++ {
				if entry[j] == configRev[1][i] {
					found = true
				}
			}
			if !found {
				mightBeThree = false
				break
			}
		}
		if mightBeThree {
			configRev[3] = entry
			continue
		}

		configRev[5] = entry
	}

	// 0, 6, 9
	for _, entry := range byLength[6] {
		off := invert(entry)[0]

		couldBeSix := false
		for i:=0; i<2; i++ {
			if off == configRev[1][i] {
				couldBeSix = true
				break
			}
		}
		if couldBeSix {
			configRev[6] = entry
			continue
		}

		couldBeZero := false
		for i:=0; i<len(configRev[2]); i++ {
			if off == configRev[2][i] {
				couldBeZero = true
				break
			}
		}
		if !couldBeZero {
			configRev[9] = entry
			continue
		}

		couldBeZero = false
		for i:=0; i<len(configRev[3]); i++ {
			if off == configRev[3][i] {
				couldBeZero = true
				break
			}
		}
		if !couldBeZero {
			configRev[9] = entry
			continue
		}

		couldBeZero = false
		for i:=0; i<len(configRev[3]); i++ {
			if off == configRev[3][i] {
				couldBeZero = true
				break
			}
		}
		if couldBeZero {
			configRev[0] = entry
		} else {
			configRev[9] = entry
		}
	}

	config := make(map[string]int)
	for k, v := range configRev {
		config[v] = k
	}
	//fmt.Println(configs, " -> ", config)
	return config
}

func parseInput() []LCD {
	var displays []LCD
	config := make(map[string]int)
	var value [4]string

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " | ")

		config = parseConfig(strings.Split(parts[0], " "))

		for ix, valueEntry := range strings.Split(parts[1], " ") {
			value[ix] = canonical(valueEntry)
		}

		displays = append(displays, LCD{config, value})
	}

	return displays
}

func isUnique(s string) bool {
	return len(s) == 2 || len(s) == 3 || len(s) == 4 || len(s) == 7
}

func part1() {
	displays := parseInput()
	total := 0
	for _, display := range displays {
		for i:=0; i<4; i++ {
			if isUnique(display.value[i]) {
				total++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", total)
}

func part2() {
	var total uint64

	for _, lcd := range parseInput() {
		total += uint64(lcd.intValue())
	}

	fmt.Printf("Part 2: %d\n", total)
}

// 56703 low

func main() {
	part1()
	part2()
}
