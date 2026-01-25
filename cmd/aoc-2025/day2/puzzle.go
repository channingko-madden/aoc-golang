package main

import "strconv"

// Part1
// Find sequences of digits repeated twice.
// Ex. 55 (5 twice), 6464 (64 twice), and 123123 (123 twice)
func isInvalidPt1(id string) bool {

	length := len(id)

	if length%2 != 0 {
		return false
	}

	half := length / 2

	return id[:half] == id[half:]
}

// Part2
// Find some sequence of digits repeated at least twice
//
// Have to break down string into the possible repeatible amounts
// Loop until length / 2
// If length % i is 0, check the substrings of size i
func isInvalidPt2(id string) bool {

	length := len(id)

	for i := 1; i <= length/2; i++ {
		if length%i == 0 {
			var subStrings []string
			for j := 0; j < length; j += i {
				subStrings = append(subStrings, id[j:j+i])
			}

			theyMatch := true
			for _, sub := range subStrings {
				theyMatch = theyMatch && (sub == subStrings[0])
			}
			if theyMatch {
				return true
			}
		}
	}

	return false
}

type IDRange struct {
	Begin int64
	End   int64
}

func (i *IDRange) getInvalidIDs() []int64 {

	output := make([]int64, 1)

	for n := i.Begin; n <= i.End; n++ {
		s := strconv.FormatInt(n, 10)

		/*
		   if isInvalidPt1(s) {
		           output = append(output, n)
		   }
		*/
		if isInvalidPt2(s) {
			output = append(output, n)
		}
	}

	return output
}
