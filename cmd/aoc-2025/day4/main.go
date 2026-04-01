package main

import (
	"flag"
	"github.com/channingko-madden/aoc-golang/internal"
)

/*
'@' is a roll of paper,
'.' is an empty space

The 8 adjacent positions in the grid are the spaces around a point in the 2D grid.
A A A
A P A
A A A

- Some positions on the edges of the grid will have less than 8 adjacent spaces.

Corners have 3

P A
A A

Edges have 5.

A A
P A
A A

*/

/*
Idea:
Parse input into a 2D slice

Create a struct to encapsulate accessing adjacent positions?

Count number of positions with less than 4 adjacent rolls.

*/

func main() {

	inputFile := flag.String("input", "", "Input file")
	flag.Parse()

	if len(*inputFile) == 0 {
		println("input file is required")
		return
	}

	occupancy_map := parseLines(internal.ReadAndParseData(*inputFile))

	maxRow := len(occupancy_map) - 1
	maxCol := len(occupancy_map[0]) - 1

	accessableRolls := 0
	for rowNum, col := range occupancy_map {
		for colNum := range col {
			if occupancy_map[rowNum][colNum] == 1 {
				point := Part1Point{
					row: rowNum,
					col: colNum,
				}

				neighbors := point.get_neighbors(maxRow, maxCol)

				neighbor_count := 0
				for _, n := range neighbors {
					neighbor_count += occupancy_map[n.row][n.col]
				}

				if neighbor_count < 4 {
					accessableRolls += 1
				}
			}

		}
	}

	println(accessableRolls)
}

func parseLine(line string) []int {
	row := make([]int, 0, len(line))
	for _, r := range line {
		if r == '@' {
			row = append(row, 1)
		} else {
			row = append(row, 0)
		}
	}
	return row
}

// Returns a 2D occupancy map
func parseLines(lines []string) [][]int {
	// create a row
	output := make([][]int, 0, len(lines))
	for _, line := range lines {
		output = append(output, parseLine(line))
	}
	return output
}

func printMap(occupancyMap [][]int) {

	for _, line := range occupancyMap {
		for _, p := range line {
			print(p)
		}
		println()
	}

}

type Part1Point struct {
	row int
	col int
}

func (p *Part1Point) get_neighbors(maxRow int, maxCol int) []Part1Point {
	row_above := p.row > 0
	row_below := p.row < maxRow

	col_left := p.col > 0
	col_right := p.col < maxCol

	neighbors := make([]Part1Point, 0, 6)

	if row_above {
		neighbors = append(neighbors, Part1Point{row: p.row - 1, col: p.col})

		if col_left {
			neighbors = append(neighbors, Part1Point{row: p.row - 1, col: p.col - 1})
		}

		if col_right {
			neighbors = append(neighbors, Part1Point{row: p.row - 1, col: p.col + 1})
		}
	}

	if row_below {
		neighbors = append(neighbors, Part1Point{row: p.row + 1, col: p.col})

		if col_left {
			neighbors = append(neighbors, Part1Point{row: p.row + 1, col: p.col - 1})
		}

		if col_right {
			neighbors = append(neighbors, Part1Point{row: p.row + 1, col: p.col + 1})
		}

	}

	if col_left {
		neighbors = append(neighbors, Part1Point{row: p.row, col: p.col - 1})

	}

	if col_right {
		neighbors = append(neighbors, Part1Point{row: p.row, col: p.col + 1})
	}

	return neighbors
}
