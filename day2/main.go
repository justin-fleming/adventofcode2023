package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var sum int
	for scanner.Scan() {
		// second arg is part number..
		sum += doWork(scanner.Text(), 2)
	}
	fmt.Println(sum)
}

func doWork(game string, partNum int) int {
	gameAndRounds := strings.Split(game, ":")
	roundsString := gameAndRounds[1]
	splitRounds := strings.Split(roundsString, ";")

	redTotal := 12
	greenTotal := 13
	blueTotal := 14

	redMin := 0
	greenMin := 0
	blueMin := 0

	for _, round := range splitRounds {
		colors := strings.Split(round, ",")
		for _, color := range colors {
			countColor := strings.Split(color, " ")
			if countColor[2] == "red" {
				redCount, err := strconv.Atoi(countColor[1])
				if err != nil {
					fmt.Println("ERROR:", err)
				}
				// fmt.Println("\tRed Count:", redCount)
				if partNum == 1 {
					if redCount > redTotal {
						fmt.Println("Round not possible, TOO MANY RED:", redCount)
						return 0
					}
				} else {
					if redCount > redMin {
						redMin = redCount
					}
				}

			}
			if countColor[2] == "green" {
				greenCount, err := strconv.Atoi(countColor[1])
				if err != nil {
					fmt.Println("ERROR:", err)
				}
				// fmt.Println("\tGreen Count:", greenCount)
				if partNum == 1 {
					if greenCount > greenTotal {
						fmt.Println("Round not possible, TOO MANY GREEN:", greenCount)
						return 0
					}
				} else {
					if greenCount > greenMin {
						greenMin = greenCount
					}
				}
			}
			if countColor[2] == "blue" {
				blueCount, err := strconv.Atoi(countColor[1])
				if err != nil {
					fmt.Println("ERROR:", err)
				}
				// fmt.Println("\tBlue Count:", blueCount)
				if partNum == 1 {
					if blueCount > blueTotal {
						fmt.Println("Round not possible, TOO MANY GREEN:", blueCount)
						return 0
					}
				} else {
					if blueCount > blueMin {
						blueMin = blueCount
					}
				}
			}
		}
	}
	if partNum == 1 {
		gameNum, err := strconv.Atoi(strings.Split(gameAndRounds[0], " ")[1])
		if err != nil {
			fmt.Println("ERROR:", err)
		}
		fmt.Println("Round", gameNum, "possible")
		return gameNum
	} else {
		power := redMin * greenMin * blueMin
		fmt.Println("Power:", power)
		return power
	}
}

// func part2() {

// }
