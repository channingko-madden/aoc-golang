package main

import (
	"fmt"
	"github.com/channingko-madden/aoc-golang/internal"
)

func main() {

	lines := internal.ReadAndParseData("aoc_day1_input.txt")

	var value int

	for _, l := range lines {
		v := ParseCalibrationValue([]byte(l))
		fmt.Println(v)
		value += v
	}

	fmt.Println(value)
}
