package main

import (
	"fmt"
	"testing"
)

func TestCountVowels(t *testing.T) {
	var tests = []struct {
		input string
		expected int
	}{
		{"bhpkc", 0},
		{"aeiou", 5},
		{"afterwards", 3},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%s,%d", test.input, test.expected)

		t.Run(testName, func(t *testing.T) {
			if countVowels(test.input) != test.expected {
				t.Fatalf("expected %d vowels", test.expected)
			}
		})
	}
}

func TestGotRepeatedPair(t *testing.T) {
	var tests = []struct {
		input string
		expected bool
	}{
		{"kckc", true},
		{"kcookc", true},
		{"adfgasdd", false},
		{"aaab", false},
		{"aokbok", true},
		{"xyxyab", true},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%s,%t", test.input, test.expected)

		t.Run(testName, func(t *testing.T) {
			if gotRepeatedPair(test.input) != test.expected {
				t.Fatalf("expected repeated pairs to be %t", test.expected)
			}
		})
	}
}