package main

import (
	"fmt"
	"testing"
)

func TestInvert(t *testing.T) {
	var tests = []struct {
		input string
		expected string
	}{
		{"abc", "defg"},
		{"abcefg", "d"},
		{"", "abcdefg"},
		{"abcdefg", ""},
		{"a", "bcdefg"},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%s->%s", test.input, test.expected)
		t.Run(testName, func(t *testing.T) {
			if invert(test.input) != test.expected {
				t.Fatalf("expected %d vowels", test.expected)
			}
		})
	}
}