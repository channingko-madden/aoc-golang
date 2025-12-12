package main

import (
	"flag"
	"fmt"
	"github.com/channingko-madden/aoc-golang/internal"
)

func main() {

	inputFile := flag.String("input", "", "Input file")
	flag.Parse()

	if len(*inputFile) == 0 {
		println("input file is required")
	}

	inputLines := internal.ReadAndParseData(*inputFile)

	dial := newDial()
	var part1Output int = 0
	var part2Output int = 0

	for _, line := range inputLines {
		direction, value := parseInputLine(line)

		switch direction {
		case "R":
			part2Output += dial.rotateRight(value)
		case "L":
			part2Output += dial.rotateLeft(value)
		default:
			fmt.Println("What is this input? ", line)
		}

		if dial.Value < 0 || dial.Value > 99 {
			panic(fmt.Sprintf("Done messed up %d %s %d", dial.Value, direction, value))
		}

		if dial.Value == 0 {
			part1Output++
		}
	}

	// Possible final marker landing on 0
	if dial.Value == 0 {
		part2Output++
	}

	fmt.Println("Part 1 Output: ", part1Output)
	fmt.Println("Part 2 Output: ", part2Output)

}
