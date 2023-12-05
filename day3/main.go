package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	grid := map[int]string{}
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid[count] = line
		count++
	}
	// part1(grid)
	part2(grid)
}

func adjacentToSymbol(grid map[int]string, row int, columnStart int, columnEnd int) bool {
	if columnStart > 0 {
		columnStart = columnStart - 1
		if checkSymbol(string(grid[row][columnStart])) {
			return true
		}
	}
	if columnEnd < len(grid[row])-1 {
		if checkSymbol(string(grid[row][columnEnd])) {
			return true
		}
		columnEnd = columnEnd + 1
	}

	if row > 0 {
		if checkSymbol(grid[row-1][columnStart:columnEnd]) {
			return true
		}
	}
	if row < len(grid)-1 {
		if checkSymbol(grid[row+1][columnStart:columnEnd]) {
			return true
		}
	}
	return false
}

func checkSymbol(s string) bool {
	re := regexp.MustCompile("[!@#$%^&*()=+/-]")
	if re.FindStringIndex(s) != nil {
		return true
	}
	return false
}

func part1(grid map[int]string) {
	sum := 0
	for i, content := range grid {
		fmt.Println("\n\n***** Line no.", i, ":", content, "*****")
		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllStringIndex(content, -1)

		for _, num := range nums {
			if adjacentToSymbol(grid, i, num[0], num[1]) {
				numValue, err := strconv.Atoi(content[num[0]:num[1]])
				if err != nil {
					fmt.Println("ERROR:", err)
				}
				sum += int(numValue)
			}
		}
	}

	fmt.Println(sum)
}

func part2(grid map[int]string) {
	sum := 0

	for i, content := range grid {
		fmt.Println("\n\n***** Line no.", i, ":", content, "*****")
		re := regexp.MustCompile("[*]")
		stars := re.FindAllStringIndex(content, -1)
		// fmt.Println("Found", len(stars), "stars")
		for _, star := range stars {
			fmt.Println("Checking star at", star)
			sum += gearRatio(grid, i, star)
		}
	}
	fmt.Println(sum)
}

func getNums(line string) [][]int {
	re := regexp.MustCompile("[0-9]+")
	nums := re.FindAllStringIndex(line, -1)
	return nums
}

func gearRatio(grid map[int]string, row int, star []int) int {
	var numsTouching []int
	starRowNums := getNums(grid[row])
	starRange := make([]int, len(star))
	copy(starRange, star)
	fmt.Println("Starting star range:", starRange)
	if star[0] > 0 {
		starRange[0] -= 1
		fmt.Println("Star not on left edge, new star range:", starRange)
		// Check left edge
		for _, num := range starRowNums {
			if num[1] == star[0] {
				numValue, err := strconv.Atoi(grid[row][num[0]:num[1]])
				if err != nil {
					fmt.Println("ERROR:", err)
				}
				numsTouching = append(numsTouching, numValue)
				// fmt.Println("FOUND NUM TO LEFT OF STAR:", numValue)
			}
		}
	}
	if star[0] < len(grid[row])-1 {
		// Check right edge
		starRange[1] += 1
		fmt.Println("Star not on right edge, new star range:", starRange)
		for _, num := range starRowNums {
			if num[0] == star[0]+1 {
				numValue, err := strconv.Atoi(grid[row][num[0]:num[1]])
				if err != nil {
					fmt.Println("ERROR:", err)
				}
				numsTouching = append(numsTouching, numValue)
				// fmt.Println("FOUND NUM TO RIGHT OF STAR:", numValue)
			}
		}
	}
	if row > 0 {
		// Check row above
		rowAboveNums := getNums(grid[row-1])
		fmt.Println("rowAboveNums=", rowAboveNums)
		for _, num := range rowAboveNums {
			if (num[0] >= starRange[0] && num[0] < starRange[1]) || (num[1]-1 >= starRange[0] && num[1]-1 < starRange[1]) {
				// fmt.Println("DEBUG>>>>>>>> grid[row-1][num[0]:num[1]]=", grid[row-1][num[0]:num[1]])
				numValue, err := strconv.Atoi(grid[row-1][num[0]:num[1]])
				if err != nil {
					fmt.Println("ERROR:", err)
				}
				numsTouching = append(numsTouching, numValue)
				// fmt.Println("FOUND NUM ABOVE OF STAR:", numValue)
			}
		}
	}
	if row < len(grid)-1 {
		// Check row below
		rowBelowNums := getNums(grid[row+1])
		// fmt.Println("rowBelowNums=", rowBelowNums)
		for _, num := range rowBelowNums {
			// fmt.Println("checking num", num)
			if (num[0] >= starRange[0] && num[0] < starRange[1]) || (num[1]-1 >= starRange[0] && num[1]-1 < starRange[1]) {
				// fmt.Println("DEBUG>>>>>>>> grid[row+1][num[0]:num[1]]=", grid[row+1][num[0]:num[1]])
				numValue, err := strconv.Atoi(grid[row+1][num[0]:num[1]])
				if err != nil {
					fmt.Println("ERROR:", err)
				}
				numsTouching = append(numsTouching, numValue)
				fmt.Println("FOUND NUM BELOW OF STAR:", numValue)
			}
		}
	}

	fmt.Println(numsTouching)

	if len(numsTouching) == 2 {
		fmt.Println("...........FOUND GEAR, 2 NUMS TOUCHING:", numsTouching)
		return numsTouching[0] * numsTouching[1]
	} else {
		return 0
	}
}
