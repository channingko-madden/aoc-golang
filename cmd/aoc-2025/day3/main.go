package main

import (
	"flag"
	"strconv"

	"github.com/channingko-madden/aoc-golang/internal"
)

func main() {
	inputFile := flag.String("input", "", "Input file")
	flag.Parse()

	if len(*inputFile) == 0 {
		println("input file is required")
		return
	}

	inputLines := internal.ReadAndParseData(*inputFile)

	joltage := 0

	for _, line := range inputLines {
		//joltage += Part1Joltage(line)
		joltage += Part2Joltage(line)
	}

	println(joltage)
}

// Two characters from the input string to form an integer
// Want to max integer that can be formed.
// The order of characters cannot be changed
func Part1Joltage(batteries string) int {
	maxJoltage := 0

	for i := 0; i < len(batteries)-1; i++ {
		leadRune := batteries[i]
		for j := i + 1; j < len(batteries); j++ {
			runes := []byte{leadRune, batteries[j]}
			joltage, err := strconv.Atoi(string(runes))

			if err != nil {
				panic(err)
			}

			if joltage > maxJoltage {
				maxJoltage = joltage
			}
		}

		// if the next leadRune is less than the current leadRune, just skip it
		if batteries[i+1] < leadRune {
			i++
		}
	}
	return maxJoltage
}

// Twelve characters from the input string to form an integer
// Want the max integer that can be formed.
// The order of characters cannot be changed
//
// Solution is the FindMaxRun recursive function
func Part2Joltage(batteries string) int {
	remainingBatts := 12
	result := make([]byte, 0, 12)
	batteriesBytes := []byte(batteries)
	result = FindMaxRune(batteriesBytes, result, remainingBatts)

	joltage, err := strconv.Atoi(string(result))
	if err != nil {
		panic(err)
	}
	return joltage
}

const maxDigit byte = byte(9)

// Find the earliest max value from the source slice, and return the dest slice
// with this value appended to it
func FindMaxRune(source []byte, dest []byte, remainingBatts int) []byte {

	if remainingBatts == 0 {
		return dest
	}

	maxRune := byte(0)
	index := 0
	// The max value of i is determined by how many batteries remain to be found
	// because there must always be enough remaining for the next recursive call
	for i := 0; i < len(source)-(remainingBatts-1); i++ {
		if source[i] > maxRune {
			index = i
			maxRune = source[i]

			if maxRune == maxDigit {
				break
			}
		}

	}

	dest = append(dest, maxRune)

	return FindMaxRune(source[index+1:], dest, remainingBatts-1)
}
