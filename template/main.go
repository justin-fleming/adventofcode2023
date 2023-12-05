package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// TODO: iterate over lines of input here...
	}
}

func part1() {

}

func part2() {

}
