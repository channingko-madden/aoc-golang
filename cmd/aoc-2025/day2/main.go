package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Input ranges are separated by commas (,); each range gives its first ID and last ID separated
// by a dash (-).

var testInput string = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

var input string = "24-46,124420-259708,584447-720297,51051-105889,6868562486-6868811237,55-116,895924-1049139,307156-347325,372342678-372437056,1791-5048,3172595555-3172666604,866800081-866923262,5446793-5524858,6077-10442,419-818,57540345-57638189,2143479-2274980,683602048-683810921,966-1697,56537997-56591017,1084127-1135835,1-14,2318887654-2318959425,1919154462-1919225485,351261-558210,769193-807148,4355566991-4355749498,809094-894510,11116-39985,9898980197-9898998927,99828221-99856128,9706624-9874989,119-335"

func parseInput(input string) ([]IDRange, error) {

	output := make([]IDRange, 1)

	for r := range strings.SplitSeq(input, ",") {
		values := strings.Split(r, "-")

		if len(values) != 2 {
			return nil, fmt.Errorf("Error splitting an input range, got %d values instead of 2", len(values))
		}

		begin, err := strconv.ParseInt(values[0], 10, 64)

		if err != nil {
			panic(err)
		}

		end, err := strconv.ParseInt(values[1], 10, 64)

		if err != nil {
			panic(err)
		}

		output = append(output, IDRange{Begin: begin, End: end})
	}

	return output, nil
}

func main() {

	idRanges, err := parseInput(input)

	if err != nil {
		panic(err)
	}

	var idSum int64 = 0

	for _, idRange := range idRanges {
		invalidIDs := idRange.getInvalidIDs()

		for _, id := range invalidIDs {
			idSum += id
		}
	}

	fmt.Println(idSum)
}
