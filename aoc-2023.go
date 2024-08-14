package main

import (
    "fmt"
    "aoc-2023-golang/day1"
)

func main() {

    lines := day1.ReadAndParseData("data/aoc_day1_input.txt")

    var value int

    for _, l := range lines {
        v := day1.ParseCalibrationValue([]byte(l))
        fmt.Println(v)
        value += v
    }
    
    fmt.Println(value)
}
