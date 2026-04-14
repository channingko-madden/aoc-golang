package internal

import (
	"bufio"
	"os"
)

// Read the file and return a slice containing each line in the file
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
