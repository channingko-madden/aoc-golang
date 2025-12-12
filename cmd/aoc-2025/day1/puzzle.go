package main

import (
	"strconv"
)

// Part 1
// The dial starts by pointing at 50.
// Dial has values 0 to 99. Values wrap.
// Read input file line by line.
// R# means rotate right (increase number)
// L# means rotate left (decrease number)
// Can be hundres of rotations, there is not limit to the size
// Done!

// Part 2
// count the number of times any click causes the dial to point at 0
// regardless of whether it happens during a rotation or at the end of one.
// My approach: count how many times the dial passes 0, only need to catch
// if it lands on 0 at the very end

type Dial struct {
	Value int
}

func newDial() Dial {
	return Dial{Value: 50}
}

// Returns the number of times the dial passes the 0 tick
func (d *Dial) rotateLeft(value int) int {

	ticks := 0
	if value > 100 {
		ticks = value / 100
		value = value % 100
	}

	if d.Value < value {
		d.Value = 100 - (value - d.Value)
		ticks++ // always passes 0 to wrap to 99
	} else {
		d.Value -= value
	}

	return ticks
}

// Returns the number of times the dial passes the 0 tick
func (d *Dial) rotateRight(value int) int {
	ticks := 0
	if value > 100 {
		ticks = value / 100
		value = value % 100
	}

	if (d.Value + value) > 99 {
		d.Value = (d.Value + value) - 100

		// only passes 0 if it wraps past 0
		if d.Value > 0 {
			ticks++
		}
	} else {
		// passes 0
		if d.Value == 0 {
			ticks++
		}
		d.Value += value
	}

	return ticks
}

// Parse an input line. Ex ("R11") returns ("R", 11)
func parseInputLine(line string) (string, int) {
	direction := string(line[0])
	value, err := strconv.Atoi(line[1:])
	if err != nil {
		panic(err)
	}

	return direction, value

}
