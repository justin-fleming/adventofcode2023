package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	var initMap [][]string
	for scanner.Scan() {
		// TODO: iterate over lines of input here...
		row := strings.Split(scanner.Text(), "")
		initMap = append(initMap, row)
	}

	expandMap(initMap)
	// printGalaxyMap(expandedMap)
}

func printGalaxyMap(galaxyMap [][]string) {
	for _, r := range galaxyMap {
		fmt.Println(r)
	}
}

func expandMap(galaxyMap [][]string) [][]string {
	// Check for empty rows
	var galaxyInitLocs [][]int
	var emptyRows []int

	for r := 0; r < len(galaxyMap); r++ {
		rowIsEmpty := true
		for c := 0; c < len(galaxyMap[r]); c++ {
			if galaxyMap[r][c] == "#" {
				// capture locations of galaxies initial locations
				galaxyInitLocs = append(galaxyInitLocs, []int{r, c})
				rowIsEmpty = false
			} else if c == len(galaxyMap[r])-1 && rowIsEmpty {
				fmt.Println("Row", r, "has no galaxies... need to expand")
				emptyRows = append(emptyRows, r)
			}
		}
	}

	// Check for empty columns
	var emptyCols []int
	for c := 0; c < len(galaxyMap[0]); c++ {
		for r := 0; r < len(galaxyMap); r++ {
			if galaxyMap[r][c] == "#" {
				// already captured init locs above, can break out of loop to save time
				break
			} else if r == len(galaxyMap)-1 {
				fmt.Println("Col", c, "has no galaxies... need to expand")
				emptyCols = append(emptyCols, c)
			}
		}
	}

	var dists []int

	for startGalaxy := 0; startGalaxy < len(galaxyInitLocs); startGalaxy++ {
		for endGalaxy := startGalaxy + 1; endGalaxy < len(galaxyInitLocs); endGalaxy++ {
			// part 1 expansion rate = 1
			// part 2 expansion rate = 999999
			dists = append(dists, part1(galaxyInitLocs[startGalaxy], galaxyInitLocs[endGalaxy], emptyRows, emptyCols, 999999))
		}
	}

	total := 0
	for _, d := range dists {
		total += d
	}
	fmt.Println("TOTAL OF DISTS:", total)
	return galaxyMap
}

func part1(startGalaxy []int, endGalaxy []int, emptyRows []int, emptyCols []int, expansionRate int) int {
	// fmt.Println("********** getting distance from", startGalaxy, "to", endGalaxy)
	rowDiff := int(math.Abs(float64((startGalaxy[0] - endGalaxy[0]))))
	colDiff := int(math.Abs(float64((startGalaxy[1] - endGalaxy[1]))))

	// fmt.Println("initRowDiff:", rowDiff)
	// fmt.Println("initColDiff:", colDiff)

	for _, er := range emptyRows {
		if (er > startGalaxy[0] && er < endGalaxy[0]) || (er < startGalaxy[0] && er > endGalaxy[0]) {
			rowDiff += expansionRate
		}
	}

	for _, cr := range emptyCols {
		if (cr > startGalaxy[1] && cr < endGalaxy[1]) || (cr < startGalaxy[1] && cr > endGalaxy[1]) {
			colDiff += expansionRate
		}
	}

	// fmt.Println("adjusted RowDiff:", rowDiff)
	// fmt.Println("adjusted ColDiff:", colDiff)

	sum := rowDiff + colDiff

	// fmt.Println("returning sum:", sum)
	return sum
}

func part2() {

}
