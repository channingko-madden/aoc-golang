package main

import (
	"testing"
)

func TestPart2Joltage(t *testing.T) {

	testCases := []struct {
		input  string
		output int
	}{
		{"987654321111111", 987654321111},
		{"811111111111119", 811111111119},
		{"234234234234278", 434234234278},
		{"818181911112111", 888911112111},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {

			actual := Part2Joltage(tc.input)
			if actual != tc.output {
				t.Errorf("got %d, want %d", actual, tc.output)
			}
		})
	}

}
