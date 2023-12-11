package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards          []int
	optimizedCards []int
	handType       int
	optHandType    int
	handNum        int
	bid            int
}

func main() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	cardMap := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}

	var hands []Hand
	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		tempCards := make([]int, 5)
		tempCardsOpt := make([]int, 5)
		var tempType int

		for i, c := range strings.Split(line[0], "") {
			tempCards[i] = cardMap[c]
			if c == "J" {
				tempCardsOpt[i] = 1
			} else {
				tempCardsOpt[i] = cardMap[c]
			}
			tempType = getHandType(tempCards)
		}

		tempBid, _ := strconv.Atoi(line[1])

		hand := Hand{
			cards:    tempCards,
			handType: tempType,
			handNum:  count,
			bid:      tempBid,
		}

		hands = append(hands, hand)
		count++
	}
	fmt.Println("Unsorted Hands")
	printHands(hands)

	fmt.Println("Sorted Hands")
	sort.Sort(HandByType(hands))

	printHands(hands)

	answer := 0
	for i, h := range hands {
		answer += h.bid * (i + 1)
	}
	fmt.Println("***** Part 1 Answer:", answer)
}

// Implement the Sort interface to sort the hands by type (ie. 3 of a kind is better than a pair)
type HandByType []Hand

func (h HandByType) Len() int {
	return len(h)
}

func (h HandByType) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h HandByType) Less(i, j int) bool {
	if h[i].handType != h[j].handType {
		return h[i].handType < h[j].handType
	}
	return breakTie(h[i].cards, h[j].cards)
}

func breakTie(h1 []int, h2 []int) bool {
	for c := 0; c < 5; c++ {
		if h1[c] < h2[c] {
			return true
		} else if h2[c] < h1[c] {
			return false
		}
	}
	return false
}

func getHandType(hand []int) int {
	cardCounts := map[int]int{
		14: 0,
		13: 0,
		12: 0,
		11: 0,
		10: 0,
		9:  0,
		8:  0,
		7:  0,
		6:  0,
		5:  0,
		4:  0,
		3:  0,
		2:  0,
		1:  0,
	}

	for _, card := range hand {
		cardCounts[card]++
	}

	/* Return an int that corresponds to this map:
	"five of a kind": 	7
	"four of a kind": 	6
	"full house": 		5
	"three of a kind": 	4
	"two pair": 		3
	"pair":				2
	"high card": 		1
	*/

	threeFound := false
	twoFound := false

	// TODO: add logic for optimal hand type here...
	for _, v := range cardCounts {
		if v == 5 {
			return 7
		}
		if v == 4 {
			return 6
		}
		if v == 3 {
			threeFound = true
		}
		if v == 2 {
			if twoFound {
				return 3
			}
			twoFound = true
		}
	}

	if threeFound && twoFound {
		return 5
	} else if threeFound {
		return 4
	} else if twoFound {
		return 2
	}

	return 1
}

func printHands(hands []Hand) {
	for _, h := range hands {
		fmt.Println("Hand no.", h.handNum, ":", h.cards, "Type:", h.handType)
	}
}

// func part2() {

// }
