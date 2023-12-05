package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type resourceMap struct {
	destStart int
	srcStart  int
	rangeLen  int
}

func main() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	// Process input
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		// TODO: iterate over lines of input here...
		lines = append(lines, scanner.Text())
	}

	f.Close()

	seeds := strings.Fields(strings.Split(lines[0], ":")[1])
	fmt.Println("SEEDS>>>>>>>>>>>>>>>", seeds)

	var seedToSoil []resourceMap
	var soilToFert []resourceMap
	var fertToWater []resourceMap
	var waterToLight []resourceMap
	var lightToTemp []resourceMap
	var tempToHum []resourceMap
	var humToLoc []resourceMap

	var currentSection string

	for _, line := range lines[2:] {
		if strings.Contains(line, ":") {
			currentSection = strings.Split(line, ":")[0]
		} else {
			if line == "" {
				continue
			}

			lineContents := strings.Fields(line)
			destStartInt, _ := strconv.Atoi(lineContents[0])
			srcStartInt, _ := strconv.Atoi(lineContents[1])
			rangeLenInt, _ := strconv.Atoi(lineContents[2])
			resToResMap := resourceMap{
				destStart: destStartInt,
				srcStart:  srcStartInt,
				rangeLen:  rangeLenInt,
			}

			switch section := currentSection; section {
			case "seed-to-soil map":
				seedToSoil = append(seedToSoil, resToResMap)
			case "soil-to-fertilizer map":
				soilToFert = append(soilToFert, resToResMap)
			case "fertilizer-to-water map":
				fertToWater = append(fertToWater, resToResMap)
			case "water-to-light map":
				waterToLight = append(waterToLight, resToResMap)
			case "light-to-temperature map":
				lightToTemp = append(lightToTemp, resToResMap)
			case "temperature-to-humidity map":
				tempToHum = append(tempToHum, resToResMap)
			case "humidity-to-location map":
				humToLoc = append(humToLoc, resToResMap)
			default:
				fmt.Println("WARNING>>>>> SHOULD NOT HIT THIS CASE...")
			}
		}
	}

	lowestSeedLoc := math.MaxInt
	for _, seed := range seeds {
		seedNum, _ := strconv.Atoi(seed)

		soil := convert(seedNum, seedToSoil)
		fert := convert(soil, soilToFert)
		water := convert(fert, fertToWater)
		light := convert(water, waterToLight)
		temp := convert(light, lightToTemp)
		hum := convert(temp, tempToHum)
		loc := convert(hum, humToLoc)

		// loc := processSeed(seedNum)
		if loc < lowestSeedLoc {
			lowestSeedLoc = loc
		}
	}

	fmt.Println("ANSWER:", lowestSeedLoc)

	// fmt.Println("SEEDS:", seeds)
	// fmt.Println("seedToSoil:", seedToSoil)
	// fmt.Println("soilToFert:", soilToFert)
	// fmt.Println("fertToWater:", fertToWater)
	// fmt.Println("waterToLight:", waterToLight)
	// fmt.Println("lightToTemp:", lightToTemp)
	// fmt.Println("tempToHum:", tempToHum)
	// fmt.Println("humToLoc:", humToLoc)
}

func convert(startRes int, resMaps []resourceMap) int {
	for _, resMap := range resMaps {
		if startRes >= resMap.srcStart && startRes <= resMap.srcStart+resMap.rangeLen {
			// fmt.Println("Seed", startRes, "in range", resMap)
			return resMap.destStart + (startRes - resMap.srcStart)
		}
	}
	return startRes
}

func part1() {

}

func part2() {

}
