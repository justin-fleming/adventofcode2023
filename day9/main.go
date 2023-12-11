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
	sum := 0
	revSum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		var seq []int
		for _, n := range line {
			nVal, _ := strconv.Atoi(n)
			seq = append(seq, nVal)
		}
		// fmt.Println("***** Working on seq:", seq)
		nextValue := extrapolate(seq)
		nextValueRev := extrapolateReverse(seq)
		sum += nextValue
		revSum += nextValueRev
	}
	fmt.Println("\tNext values sum = ", sum)
	fmt.Println("\tNext values REVERSE sum = ", revSum)
}

func part1(line string) {

}

func extrapolate(seq []int) int {
	var newSeq []int
	allSame := true
	for i := 0; i < len(seq)-1; i++ {
		newSeq = append(newSeq, seq[i+1]-seq[i])
		if seq[i+1] != seq[i] {
			allSame = false
		}
	}
	if !allSame {
		newVal := extrapolate(newSeq)
		return newVal + seq[len(seq)-1]
	} else {
		return seq[len(seq)-1]
	}
}

func extrapolateReverse(seq []int) int {
	var newSeq []int
	allSame := true
	for i := len(seq) - 1; i > 0; i-- {
		newSeq = append([]int{seq[i] - seq[i-1]}, newSeq...)
		if seq[i] != seq[i-1] {
			allSame = false
		}
	}
	if !allSame {
		newVal := extrapolateReverse(newSeq)
		return seq[0] - newVal
	} else {
		return seq[0]
	}
}
