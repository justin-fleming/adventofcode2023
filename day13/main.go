package main

import (
	"bufio"
	"fmt"
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

	var patterns [][]string
	var tempPattern []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			patterns = append(patterns, tempPattern)
			tempPattern = nil
			continue
		}

		tempPattern = append(tempPattern, strings.TrimSpace(line))
	}
	patterns = append(patterns, tempPattern)

	part1Answer := 0
	for c, pattern := range patterns {
		fmt.Println("\n\n*****Processing pattern", c)
		// Check for horizontal mirrors:
		fmt.Println("...Checking for horizontal mirror")
		part1Answer += 100 * part1(pattern)

		transposedPattern := make([]string, len(pattern[0]))
		for _, p := range pattern {
			lineData := strings.Split(p, "")
			for i := 0; i < len(lineData); i++ {
				transposedPattern[i] += lineData[i]
			}
		}
		for _, l := range transposedPattern {
			fmt.Println(l)
		}

		fmt.Println("...Checking for vertical mirror")
		part1Answer += part1(transposedPattern)
	}

	fmt.Println(part1Answer)
}

func part1(pattern []string) int {
	// Check rows for reflections
	for i := 0; i < len(pattern)-1; i++ {
		if pattern[i] == pattern[i+1] {
			fmt.Println("Found potential horizontal mirror... rows", i, "and", i+1, "=", pattern[i])
		innerLoop:
			for j := 1; j < len(pattern); j++ {
				fmt.Println("Checking", j, "away from rows", i, "-", i+1)
				if i-j >= 0 && i+1+j < len(pattern) {
					if pattern[i-j] != pattern[i+1+j] {
						fmt.Println("...not a mirror, pattern[", i-j, "]=", pattern[i-j], "does NOT match pattern[", i+1+j, "]=", pattern[i+1+j])
						break innerLoop
					}
				} else {
					fmt.Println("Found a mirror between", i, "and", i+1)
					return i + 1
				}
			}
		}
	}
	return 0
}

func part2() {

}
