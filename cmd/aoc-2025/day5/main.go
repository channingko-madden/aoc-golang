package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/channingko-madden/aoc-golang/internal"
	"strconv"
	"strings"
)

func main() {

	inputFile := flag.String("input", "", "Input file")
	flag.Parse()

	if len(*inputFile) == 0 {
		println("input file is required")
		return
	}

	strRanges, strSingles := SplitInput(*inputFile)

	intSingles := make([]int, len(strSingles))
	for _, strValue := range strSingles {
		intValue, err := strconv.Atoi(strValue)
		if err != nil {
			panic(fmt.Sprintf("Done messed up converting a single number %s: %s", strValue, err))
		}
		intSingles = append(intSingles, intValue)
	}

	numberRange := CreateNumberRanges(strRanges)

	println(Part1Fresh(numberRange, intSingles))
	println(Part2Fresh(numberRange))

}

/*

Input format:
3-5
10-14
16-20
12-18

1
5
8
11
17
32

There are number ranges, a blank line, and then single numbers
*/

// Return the number ranges and then single numbers from the input file
func SplitInput(filename string) ([]string, []string) {

	lines := internal.ReadAndParseData(filename)

	number_ranges := make([]string, 0, 1)
	single_ids := make([]string, 0, 1)

	ranges_over := false
	for _, line := range lines {
		if len(line) == 0 {
			ranges_over = true
			continue
		}

		if ranges_over {
			single_ids = append(single_ids, line)
		} else {
			number_ranges = append(number_ranges, line)
		}
	}

	return number_ranges, single_ids

}

type NumberRange struct {
	Min int
	Max int
}

func (r *NumberRange) contains(value int) bool {
	return r.Min <= value && value <= r.Max
}

func overlaps(r *NumberRange, other *NumberRange) bool {
	return r.contains(other.Min) || r.contains(other.Max) || other.contains(r.Min) || other.contains(r.Max)

}

func (r *NumberRange) count() int {
	return (r.Max - r.Min) + 1
}

func merge(r1 *NumberRange, r2 *NumberRange) (*NumberRange, error) {
	if !overlaps(r1, r2) {
		return nil, errors.New("Cannot merge non-overlapping NumberRange")
	}

	newMin := min(r1.Min, r2.Min)
	newMax := max(r1.Max, r2.Max)

	return &NumberRange{Min: newMin, Max: newMax}, nil
}

func CreateNumberRanges(input []string) []NumberRange {
	output := make([]NumberRange, 0, len(input))
	for _, raw := range input {
		min, max, found := strings.Cut(raw, "-")
		if !found {
			panic(fmt.Sprintf("Done messed up cut parsing the number range %s", raw))
		}

		minInt, err := strconv.Atoi(min)
		if err != nil {
			panic(fmt.Sprintf("Done messed up parsing the number range %s: %s", raw, err))
		}

		maxInt, err := strconv.Atoi(max)
		if err != nil {
			panic(fmt.Sprintf("Done messed up parsing the number range %s: %s", raw, err))
		}

		output = append(output, NumberRange{Min: minInt, Max: maxInt})
	}

	return output
}

func Part1Fresh(numRanges []NumberRange, ingredientIDs []int) int {
	freshCount := 0
	for _, id := range ingredientIDs {
		for _, numRange := range numRanges {
			if numRange.contains(id) {
				freshCount++
				break
			}
		}
	}

	return freshCount
}

func CollapseNumberRanges(numRanges []*NumberRange) []*NumberRange {

	initialLength := len(numRanges)

	mergedRanges := make([]*NumberRange, 0)
	for _, r := range numRanges {
		mergedOccurred := false
		for i, mRange := range mergedRanges {
			if overlaps(r, mRange) {
				newMerge, err := merge(r, mRange)
				if err != nil {

					panic(fmt.Sprintf("Done messed up merging number ranges %s", err))
				}
				mergedRanges[i] = newMerge
				mergedOccurred = true
				break
			}
		}
		if !mergedOccurred {
			mergedRanges = append(mergedRanges, r)
		}

	}

	if len(mergedRanges) == initialLength {
		return mergedRanges
	}

	return CollapseNumberRanges(mergedRanges)
}

// Returns the number of IDs considered to be fresh
func Part2Fresh(numRanges []NumberRange) int {

	initialState := make([]*NumberRange, 0, len(numRanges))
	for _, r := range numRanges {
		initialState = append(initialState, &r)
	}

	collapsedRanges := CollapseNumberRanges(initialState)

	count := 0
	for _, r := range collapsedRanges {
		count += r.count()
	}
	return count
}
