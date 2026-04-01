package main

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	testCases := []struct {
		input  string
		output []int
	}{
		{"..@@.@@@@.", []int{0, 0, 1, 1, 0, 1, 1, 1, 1, 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {

			actual := parseLine(tc.input)
			if len(actual) != len(tc.output) {
				t.Errorf("got %d output length, want %d", len(actual), len(tc.output))
			}

			for i, actual_val := range actual {
				if actual_val != tc.output[i] {
					t.Errorf("got %d at index %d, want %d", actual_val, i, tc.output[i])
				}
			}
		})
	}

}

func TestGetNeighbors(t *testing.T) {
	testCases := []struct {
		name   string
		input  Part1Point
		output []Part1Point
	}{
		{
			"origin",
			Part1Point{row: 0, col: 0}, []Part1Point{
				{row: 1, col: 0},
				{row: 1, col: 1},
				{row: 0, col: 1},
			}},
		{
			"top edge",
			Part1Point{row: 0, col: 1}, []Part1Point{
				{row: 1, col: 1},
				{row: 1, col: 0},
				{row: 1, col: 2},
				{row: 0, col: 0},
				{row: 0, col: 2},
			}},
		{
			"top right corner",
			Part1Point{row: 0, col: 2}, []Part1Point{
				{row: 1, col: 2},
				{row: 1, col: 1},
				{row: 0, col: 1},
			}},
		{
			"right edge",
			Part1Point{row: 1, col: 2}, []Part1Point{
				{row: 0, col: 2},
				{row: 0, col: 1},
				{row: 2, col: 2},
				{row: 2, col: 1},
				{row: 1, col: 1},
			}},
		{
			"bottom right corner",
			Part1Point{row: 2, col: 2}, []Part1Point{
				{row: 1, col: 2},
				{row: 1, col: 1},
				{row: 2, col: 1},
			}},
		{
			"bottom edge",
			Part1Point{row: 2, col: 1}, []Part1Point{
				{row: 1, col: 1},
				{row: 1, col: 0},
				{row: 1, col: 2},
				{row: 2, col: 0},
				{row: 2, col: 2},
			}},
		{
			"bottom left corner",
			Part1Point{row: 2, col: 0}, []Part1Point{
				{row: 1, col: 0},
				{row: 1, col: 1},
				{row: 2, col: 1},
			}},
		{
			"left edge",
			Part1Point{row: 1, col: 0}, []Part1Point{
				{row: 0, col: 0},
				{row: 0, col: 1},
				{row: 2, col: 0},
				{row: 2, col: 1},
				{row: 1, col: 1},
			}},
		{
			"center",
			Part1Point{row: 1, col: 1}, []Part1Point{
				{row: 0, col: 1},
				{row: 0, col: 0},
				{row: 0, col: 2},
				{row: 2, col: 1},
				{row: 2, col: 0},
				{row: 2, col: 2},
				{row: 1, col: 0},
				{row: 1, col: 2},
			}},
	}

	maxRow := 2
	maxCol := 2

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.input.get_neighbors(maxRow, maxCol)

			if len(actual) != len(tc.output) {
				t.Errorf("got %d output length, want %d", len(actual), len(tc.output))
			}

			for i, a := range actual {
				if a != tc.output[i] {
					t.Errorf("got %#v at index %d, want %#v", a, i, tc.output[i])
				}
			}
		})
	}

}
