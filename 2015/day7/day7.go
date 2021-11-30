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

type Wire string

type Signal uint16

// Inputs

type Input interface {
	value(Kit) Signal
}

type WireInput struct {
	wire Wire
}

func (w WireInput) value(kit Kit) Signal {
	return kit.resolve(w.wire)
}

type SignalInput struct {
	signal Signal
}

func (s SignalInput) value(_ Kit) Signal {
	return s.signal
}

type AndGate struct {
	input1 Input
	input2 Input
}

func (a AndGate) value(k Kit) Signal {
	signal1 := a.input1.value(k)
	signal2 := a.input2.value(k)
	return signal1 & signal2
}

type OrGate struct {
	input1 Input
	input2 Input
}

func (o OrGate) value(k Kit) Signal {
	signal1 := o.input1.value(k)
	signal2 := o.input2.value(k)
	return signal1 | signal2
}

type NotGate struct {
	input Input
}

func (n NotGate) value(k Kit) Signal {
	signal := n.input.value(k)
	return ^signal
}

type LShiftGate struct {
	input Input
	count int
}

func (l LShiftGate) value(k Kit) Signal {
	signal := l.input.value(k)
	return signal << l.count
}

type RShiftGate struct {
	input Input
	count int
}

func (r RShiftGate) value(k Kit) Signal {
	signal := r.input.value(k)
	return signal >> r.count
}

type Connection struct {
	input Input
	output Wire
}

type Kit struct {
	connections []Connection
	values map[Wire]Signal
}

func (kit Kit) resolve(wire Wire) Signal {
	if value, found := kit.values[wire]; found {
		return value
	} else {
		for _, connection := range kit.connections {
			if connection.output == wire {
				result := connection.input.value(kit)
				kit.values[wire] = result
				return result
			}
		}
	}
	log.Fatalf("Didn't find a connection for wire %s", wire)
	return 0
}

func parseInput() Kit {
	lines := strings.Split(input, "\n")
	kit := Kit{make([]Connection, len(lines)), make(map[Wire]Signal)}
	for _, line := range lines {

		io := strings.Split(line, " -> ")

		connInput := parseConnInput(io[0])
		connOutput := Wire(io[1])
		kit.connections = append(kit.connections, Connection{connInput,connOutput})
	}

	return kit
}

func isNumeric(s string) bool {
	for _, v := range s {
		if v < '0' || v > '9' { return false }
	}
	return true
}

func parseConnInput(rawInput string) Input {
	splitInput := strings.Split(rawInput, " ")
	switch len(splitInput) {
	case 1:
		if isNumeric(splitInput[0]) {
			value, _ := strconv.Atoi(splitInput[0])
			return SignalInput{Signal(value)}
		} else {
			return WireInput{Wire(splitInput[0])}
		}
	case 2:
		switch splitInput[0] {
		case "NOT":
				return NotGate{parseGateInput(splitInput[1])}
		default:
			log.Fatalf("len 2 unrecognised gate %s", splitInput[0])
		}
	case 3:
		switch splitInput[1] {
		case "AND":
			return AndGate{parseGateInput(splitInput[0]), parseGateInput(splitInput[2])}
		case "OR":
			return OrGate{parseGateInput(splitInput[0]), parseGateInput(splitInput[2])}
		case "LSHIFT":
			return LShiftGate{parseGateInput(splitInput[0]), parseInt(splitInput[2])}
		case "RSHIFT":
			return RShiftGate{parseGateInput(splitInput[0]), parseInt(splitInput[2])}
		default:
			log.Fatalf("what kind of gate is a %s\n", splitInput[1])
		}
	default:
		log.Fatalf("what kind of gate has %d tokens\n", len(splitInput))
	}
	return nil
}

func parseGateInput(s string) Input {
	if isNumeric(s) {
		value, _ := strconv.Atoi(s)
		return SignalInput{Signal(value)}
	} else {
		return WireInput{Wire(s)}
	}
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to parse int %s\n", s)
	}
	return i
}

func part1() {
	kit := parseInput()
	aValue := kit.resolve("a")
	fmt.Printf("Part 1: %d\n", aValue)
}

func part2() {
	kit := parseInput()
	aValue := kit.resolve("a")
	kit2 := parseInput()
	kit2.values[Wire("b")] = aValue
	aValue2 := kit2.resolve("a")
	fmt.Printf("Part 2: %d\n", aValue2)
}

func main() {
	part1()
	part2()
}