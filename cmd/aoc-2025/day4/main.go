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

func main() {

	inputFile := flag.String("input", "", "Input file")
	flag.Parse()

	if len(*inputFile) == 0 {
		println("input file is required")
		return
	}

	occupancyMap := parseLines(internal.ReadAndParseData(*inputFile))

	part1Result := part1(occupancyMap)
	part2Result := part2(occupancyMap)

	println("Part 1 result: ", part1Result)
	println("Part 2 result: ", part2Result)
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

// The forklifts can only access a roll of paper if there are fewer
// than four rolls of paper in the eight adjacent positions
func part1(occupancyMap [][]int) int {
	maxRow := len(occupancyMap) - 1
	maxCol := len(occupancyMap[0]) - 1

	accessableRolls := 0
	for rowNum, col := range occupancyMap {
		for colNum := range col {
			if occupancyMap[rowNum][colNum] == 1 {
				point := Point{
					row: rowNum,
					col: colNum,
				}

				neighbors := point.get_neighbors(maxRow, maxCol)

				neighbor_count := 0
				for _, n := range neighbors {
					neighbor_count += occupancyMap[n.row][n.col]
				}

				if neighbor_count < 4 {
					accessableRolls += 1
				}
			}

		}
	}

	return accessableRolls
}

func part2(occupancyMap [][]int) int {
	maxRow := len(occupancyMap) - 1
	maxCol := len(occupancyMap[0]) - 1

	removableRolls := 0
	rollWasRemoved := true
	//rollsToRemove := make([]Point, 0)
	for rollWasRemoved {
		rollWasRemoved = false
		for rowNum, col := range occupancyMap {
			for colNum := range col {
				if occupancyMap[rowNum][colNum] == 1 {
					point := Point{
						row: rowNum,
						col: colNum,
					}

					neighbors := point.get_neighbors(maxRow, maxCol)

					neighbor_count := 0
					for _, n := range neighbors {
						neighbor_count += occupancyMap[n.row][n.col]
					}

					if neighbor_count < 4 {
						removableRolls += 1
						rollWasRemoved = true
						// remove the roll from the map
						occupancyMap[point.row][point.col] = 0
					}
				}

			}
		}
	}

	return removableRolls
}

type Point struct {
	row int
	col int
}

func (p *Point) get_neighbors(maxRow int, maxCol int) []Point {
	row_above := p.row > 0
	row_below := p.row < maxRow

	col_left := p.col > 0
	col_right := p.col < maxCol

	neighbors := make([]Point, 0, 6)

	if row_above {
		neighbors = append(neighbors, Point{row: p.row - 1, col: p.col})

		if col_left {
			neighbors = append(neighbors, Point{row: p.row - 1, col: p.col - 1})
		}

		if col_right {
			neighbors = append(neighbors, Point{row: p.row - 1, col: p.col + 1})
		}
	}

	if row_below {
		neighbors = append(neighbors, Point{row: p.row + 1, col: p.col})

		if col_left {
			neighbors = append(neighbors, Point{row: p.row + 1, col: p.col - 1})
		}

		if col_right {
			neighbors = append(neighbors, Point{row: p.row + 1, col: p.col + 1})
		}

	}

	if col_left {
		neighbors = append(neighbors, Point{row: p.row, col: p.col - 1})

	}

	if col_right {
		neighbors = append(neighbors, Point{row: p.row, col: p.col + 1})
	}

	return neighbors
}
