package main

import (
	"testing"
)

func TestSplitInput(t *testing.T) {

	testCases := []struct {
		inputFile          string
		numberRangeOutput  []string
		singleNumberOutput []string
	}{
		{"test_input.txt", []string{"3-5", "10-14", "16-20", "12-18"}, []string{"1", "5", "8", "11", "17", "32"}},
	}

	for _, tc := range testCases {
		t.Run(tc.inputFile, func(t *testing.T) {

			actualNumberRange, actualSingleNumber := SplitInput(tc.inputFile)

			if len(actualNumberRange) != len(tc.numberRangeOutput) {
				t.Errorf(
					"number range got %d output length, want %d",
					len(actualNumberRange),
					len(tc.numberRangeOutput),
				)
			}

			if len(actualSingleNumber) != len(tc.singleNumberOutput) {
				t.Errorf(
					"single number got %d output length, want %d",
					len(actualNumberRange),
					len(tc.singleNumberOutput),
				)
			}

			for i, actualVal := range actualNumberRange {
				expectedVal := tc.numberRangeOutput[i]
				if actualVal != expectedVal {
					t.Errorf("number range got %s at index %d, want %s", actualVal, i, expectedVal)
				}
			}

			for i, actualVal := range actualSingleNumber {
				expectedVal := tc.singleNumberOutput[i]
				if actualVal != expectedVal {
					t.Errorf("number range got %s at index %d, want %s", actualVal, i, expectedVal)
				}
			}
		})
	}
}

func TestMerge(t *testing.T) {

	testCases := []struct {
		id     string
		input1 NumberRange
		input2 NumberRange
		output NumberRange
	}{
		{"case 1", NumberRange{Min: 5, Max: 10}, NumberRange{Min: 7, Max: 11}, NumberRange{Min: 5, Max: 11}},
	}
	for _, tc := range testCases {
		t.Run(tc.id, func(t *testing.T) {
			actual, err := merge(&tc.input1, &tc.input2)
			if err != nil {
				t.Errorf("Error calling merge %s", err)
			}

			if actual.Max != tc.output.Max {
				t.Errorf("Max got %d wanted %d", actual.Max, tc.output.Max)
			}

			if actual.Min != tc.output.Min {
				t.Errorf("Min got %d wanted %d", actual.Min, tc.output.Min)
			}
		},
		)
	}

}
