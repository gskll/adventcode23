package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	mapValues [][][3]int
)

func parseSeeds(line string) (seeds []int) {
	split := strings.Fields(line)
	for _, s := range split[1:] {
		seed, _ := strconv.Atoi(s)
		seeds = append(seeds, seed)
	}
	return
}

func parseMap(mapValues mapValues, line string) mapValues {
	if strings.Contains(line, "map:") {
		mapValues = append(mapValues, [][3]int{})
	} else if line == "" {
		return mapValues
	} else {
		lastMap := &mapValues[len(mapValues)-1]
		values := [3]int{}
		strValues := strings.Fields(line)
		for i, s := range strValues {
			val, _ := strconv.Atoi(s)
			values[i] = val
		}
		*lastMap = append(*lastMap, values)
	}
	return mapValues
}

func findLocations(seeds []int, maps mapValues) (locations []int) {
	for _, value := range seeds {
		for _, mapValues := range maps {
			for _, values := range mapValues {
				if value >= values[1] && value < values[1]+values[2] {
					value = value - values[1] + values[0]
					break
				}
			}
		}
		locations = append(locations, value)
	}
	return
}

func findMinLocation(locations []int) int {
	min := locations[0]
	for _, l := range locations {
		if l < min {
			min = l
		}
	}
	return min
}

func solve(input string) {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	seeds := []int{}
	mapValues := mapValues{}
	for scanner.Scan() {
		if len(seeds) == 0 {
			seeds = parseSeeds(scanner.Text())
		} else {
			mapValues = parseMap(mapValues, scanner.Text())
		}
	}
	locations := findLocations(seeds, mapValues)
	minLocation := findMinLocation(locations)
	fmt.Println("Part 1 - Min location:", minLocation)
}

func main() {
	inputFlag := flag.String("input", "day05/input.txt", "Path to input file")
	flag.Parse()

	start := time.Now()
	solve(*inputFlag)
	duration := time.Since(start)
	fmt.Printf("Execution time: %v\n", duration)
}
