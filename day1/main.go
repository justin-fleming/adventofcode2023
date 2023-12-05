package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// part1()
	part2()
}

func part1() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	var sum int
	for scanner.Scan() {
		sum += getDigits(scanner.Text())
	}
	fmt.Println(sum)
}

func part2() {
	// Read file
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	// Map of text to int
	nums := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var sum int
	// iterate over lines in file
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		// Create 2 maps with indices of first and last location within the string
		// initialized to -1 to indicate not in string
		line := scanner.Text()
		firstLoc := map[int]int{
			1: -1,
			2: -1,
			3: -1,
			4: -1,
			5: -1,
			6: -1,
			7: -1,
			8: -1,
			9: -1,
		}
		lastLoc := map[int]int{
			1: -1,
			2: -1,
			3: -1,
			4: -1,
			5: -1,
			6: -1,
			7: -1,
			8: -1,
			9: -1,
		}
		// fmt.Println("****************** Processing Line: ****************")
		fmt.Println("*", line, "*")
		// fmt.Println("****************************************************")
		fmt.Println("---------- Record first & last location text match ----------")
		// iterate over text to int map and record first and last location in the given line
		for k, v := range nums {
			firstLoc[v] = strings.Index(line, k)
			lastLoc[v] = strings.LastIndex(line, k)
		}
		fmt.Println(firstLoc)
		fmt.Println(lastLoc)

		// // iterate of the line in runes to search for digits
		// // only override the first/last location if it is before/after the text finding
		fmt.Println("---------- Record/update first & last location digit match ----------")
		runes := []rune(line)
		for j, r := range runes {
			if unicode.IsDigit(r) {
				// convert rune to int
				rindex := int(r - '0')
				if firstLoc[rindex] == -1 || j < firstLoc[rindex] {
					firstLoc[rindex] = j
				}
				break
			}
		}
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				rindex := int(runes[i] - '0')
				if lastLoc[rindex] == -1 || i > lastLoc[rindex] {
					lastLoc[rindex] = i
				}
				break
			}
		}
		fmt.Println(firstLoc)
		fmt.Println(lastLoc)

		// fmt.Println("---------- Finding first and last digits for line ----------")
		// Get first and last num in the line and add to the cumulative sum
		firstIndex := len(line)
		// fmt.Println("===firstIndex:=len(lin÷e)===:", firstIndex)
		var firstDigit int
		for k, v := range firstLoc {
			if v >= 0 && v < firstIndex {
				firstDigit = k
				firstIndex = v
			}
		}
		lastIndex := 0
		var lastDigit int
		for k, v := range lastLoc {
			if v >= lastIndex {
				lastDigit = k
				lastIndex = v
			}
		}

		lineResult, err := strconv.Atoi(strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit))
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		fmt.Println(line, ": ", lineResult)

		sum += lineResult
	}
	fmt.Println(sum)
}

func getDigits(line string) int {
	runes := []rune(line)
	var lineResult string
	for _, r := range runes {
		if unicode.IsDigit(r) {
			lineResult += string(r)
			fmt.Println("Found Digit (ltr):", string(r))
			break
		}
	}

	for i := len(runes) - 1; i >= 0; i-- {
		if unicode.IsDigit(runes[i]) {
			lineResult += string(runes[i])
			fmt.Println("Found Digit (rtl):", string(runes[i]))
			break
		}
	}

	num, err := strconv.Atoi(lineResult)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println("Digits for line", line, "are", num)
	return num
}
