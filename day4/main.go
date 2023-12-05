package main

import (
	"bufio"
	"fmt"
	"math"
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

	// part1(f)
	part2(f)
}

func isWinner(card string) int {
	splitCard := strings.Split(card, "|")
	numCounts := make(map[int]int)
	winCount := 0

	winningNums := strings.Fields(splitCard[0])
	cardNums := strings.Fields(splitCard[1])

	fmt.Println("\n**************")
	fmt.Println("Winning Nums:", winningNums)
	fmt.Println("Card Nums:", cardNums)
	for _, num := range cardNums {
		numValue, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("ERROR:", err)
		}
		numCounts[numValue]++
		// fmt.Println(numCounts)
	}
	for _, wNum := range winningNums {
		wNumValue, err := strconv.Atoi(wNum)
		if err != nil {
			fmt.Println("ERROR:", err)
		}
		// fmt.Println(">>> Checking for winning number:", wNumValue)
		if numCounts[wNumValue] != 0 {
			// fmt.Println("Found", numCounts[wNumValue], "occurences of '", wNumValue, "'")
			winCount += numCounts[wNumValue]
		}
	}
	return winCount
}

func part1(f *os.File) {
	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		card := strings.Split(scanner.Text(), ":")[1]
		sum += int(math.Pow(2, float64(isWinner(card)-1)))
	}

	fmt.Println("ANSWER:", sum)
}

func part2(f *os.File) {
	sum := 0
	scanner := bufio.NewScanner(f)

	cardCounts := make(map[int]int, 198)
	fmt.Println("CARD COUNTS:", cardCounts)
	for scanner.Scan() {
		card := strings.Split(scanner.Text(), ":")
		cardNumber, _ := strconv.Atoi(strings.Fields(card[0])[1])

		cardCounts[cardNumber]++

		winCount := isWinner(card[1])
		fmt.Println("Card ", cardNumber, "has", winCount, "wins")
		for i := cardNumber + 1; i < cardNumber+winCount; i++ {
			fmt.Println("Adding dup of card", i)
			cardCounts[i]++
			fmt.Println("CARD COUNTS:", cardCounts)
		}
	}

	for _, cards := range cardCounts {
		sum += cards
	}

	fmt.Println("ANSWER:", sum)

}
