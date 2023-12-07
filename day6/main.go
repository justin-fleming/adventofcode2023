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

	var times []float64
	var dists []float64

	var p2Time float64
	var p2Dist float64

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		if line[0] == "Time" {
			for _, time := range strings.Fields(line[1]) {
				timeVal, _ := strconv.ParseFloat(time, 64)
				times = append(times, timeVal)
			}
			p2Time, _ = strconv.ParseFloat(strings.Join(strings.Fields(line[1]), ""), 64)
		}

		if line[0] == "Distance" {
			for _, dist := range strings.Fields(line[1]) {
				distVal, _ := strconv.ParseFloat(dist, 64)
				dists = append(dists, distVal)
			}
			p2Dist, _ = strconv.ParseFloat(strings.Join(strings.Fields(line[1]), ""), 64)
		}
	}

	fmt.Println("Found the times:", times)
	fmt.Println("Found the dists:", dists)

	answer := 1

	for i := 0; i < len(times); i++ {
		winners := findWinners(times[i], dists[i])
		fmt.Println("Game", i, " (time:", times[i], ", dist:", dists[i], "):", len(winners), "winners\n\t", winners)
		answer *= len(winners)
	}

	fmt.Println("***** PT 1 ANSWER:", answer)

	winners := findWinners(p2Time, p2Dist)
	fmt.Println("***** PT 2 ANSWER: (time:", p2Time, ", dist:", p2Dist, "):", len(winners))

}

func findWinners(time float64, dist float64) []int {
	halfPoint := math.Ceil((time - 1) / 2)
	fmt.Println("Halfpoint for time", time, ":", halfPoint)

	var winners []int
	for i := halfPoint; i > 0; i-- {
		if i*(time-i) > dist {
			winners = append(winners, int(i))
			if i == halfPoint && int(time)%2 == 0 {
				continue
			}
			winners = append(winners, int(time-i))
		}
	}
	return winners
}
