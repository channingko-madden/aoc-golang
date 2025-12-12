package internal

import (
	"bufio"
	"os"
)

// Data for this is a file with a alphanumeric string on each line
// Read the file and return a slice containing each string in the file
func ReadAndParseData(filename string) []string {

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var ret []string

	for fileScanner.Scan() {
		ret = append(ret, fileScanner.Text())
	}

	file.Close()
	return ret
}
