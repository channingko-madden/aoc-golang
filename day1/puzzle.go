package day1

// Correct answer to part 1 is 53334

import (
    "errors"
    "bufio"
    "os"
    "strconv"
    "strings"
)

// The numerical value lines up with the index of this slice
var NumberWords = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// Data for this is a file with a alphanumeric string on each line
// Read the file and return a slice containing each string in the file
func ReadAndParseData(filename string) ([]string) {

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
    return ret;
}

func IsNumber(b byte) (bool) {
    return b >= 48 && b <=57
}

func GetRightNumber(line []byte) (byte, error) {
    for i := len(line) - 1; i >= 0; i-- {
        if IsNumber(line[i]) {
            return line[i], nil
        }
    }
    return 0, errors.New("No right number found")
}

func GetLeftNumber(line []byte) (byte, error) {
    for i := 0; i < len(line); i++ {
        if IsNumber(line[i]) {
            return line[i], nil
        }
    }
    return 0, errors.New("no left number found in line " + string(line))
}

// Return the index and value of the right most numeral number
func GetRightNumberAndIndex (line []byte) (int, byte, error) {
    for i := len(line) - 1; i >= 0; i-- {
        if IsNumber(line[i]) {
            return i, line[i], nil
        }
    }
    return 0, 0, errors.New("No right number found in line " + string(line))
}

func GetLeftNumberAndIndex(line []byte) (int, byte, error) {
    for i := 0; i < len(line); i++ {
        if IsNumber(line[i]) {
            return i, line[i], nil
        }
    }
    return 0, 0, errors.New("no left number found in line " + string(line))
}

// Return the indexes and numerical value of the first word number (ex. "one", "two")
// If there are no numbers, and error is returned
func GetLeftWordNumberAndIndex(line string) (int, int, error) {
    var left_index = len(line)
    var value = -1
    for index, num := range NumberWords {
        i := strings.Index(line, num) // find index of the first instance of num in line, or -1
        if i != -1 && i < left_index {
            left_index = i
            value = index
        }
    }
    if left_index == len(line) {
        return 0, 0, errors.New("No left word found in line " + string(line))
    }
    return left_index, value, nil
}

// Return the indexes and numerical value of the last word number (ex. "one", "two")
// If there are no numbers, and error is returned
func GetRightWordNumberAndIndex(line string) (int, int, error) {
    var right_index = len(line)
    var value = -1
    for index, num := range NumberWords {
        i := strings.LastIndex(line, num) // find index of the first instance of num in line, or -1
        if i != -1 && i < right_index {
            right_index = i
            value = index
        }
    }
    if right_index == len(line) {
        return 0, 0, errors.New("No right word found in line " + string(line))
    }
    return right_index, value, nil
}

// Parse the calibration value from a line for part 1
func ParseCalibrationValue(line []byte) (int) {
    left_val, err := GetLeftNumber(line)

    if err != nil {
        panic(err)
    }

    right_val, err := GetRightNumber(line) 

    if err != nil {
        panic(err)
    }

    var value_bytes = []byte{left_val, right_val}

    value, err := strconv.Atoi(string(value_bytes))

    return value
}

// Parse the calibration value from a line for part 2
func ParseCalibrationValuePartTwo(line []byte) (int) {

    left_num_index, left_num_value, err := GetLeftNumberAndIndex(line)

    if err != nil {
        panic(err)
    }

    

    var value_bytes = []byte{left_val, right_val}

    value, err := strconv.Atoi(string(value_bytes))

    return value
}
