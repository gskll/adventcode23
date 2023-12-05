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
	mappings   [][][3]int
	seedRanges [][2]int
)

func parseSeeds(line string) (seeds []int) {
	split := strings.Fields(line)
	for _, s := range split[1:] {
		seed, _ := strconv.Atoi(s)
		seeds = append(seeds, seed)
	}
	return
}

func parseSeedRanges(seeds []int) (seedRanges seedRanges) {
	for i := 0; i < len(seeds)-1; i += 2 {
		seedRanges = append(seedRanges, [2]int{seeds[i], seeds[i+1]})
	}
	return
}

func buildMapFromLine(mapValues mappings, line string) mappings {
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

func findLocationsByRange(seedRanges seedRanges, mappings mappings) []int {
	var locations []int
	for _, sdRange := range seedRanges {
		ranges := [][2]int{{sdRange[0], sdRange[0] + sdRange[1]}}
		for _, mapping := range mappings {
			var transformed [][2]int
			for _, mapFilter := range mapping {
				dst, src, srcEnd := mapFilter[0], mapFilter[1], mapFilter[1]+mapFilter[2]
				var notTransformed [][2]int
				for _, r := range ranges {
					before := [2]int{r[0], min(r[1], src)}
					inter := [2]int{max(r[0], src), min(srcEnd, r[1])}
					after := [2]int{max(srcEnd, r[0]), r[1]}

					if before[0] < before[1] {
						notTransformed = append(notTransformed, before)
					}
					if inter[0] < inter[1] {
						transformedRange := [2]int{inter[0] - src + dst, inter[1] - src + dst}
						transformed = append(transformed, transformedRange)
					}
					if after[0] < after[1] {
						notTransformed = append(notTransformed, after)
					}
				}
				ranges = notTransformed
			}
			ranges = append(ranges, transformed...)
		}
		minLocation := findMinLocationRanges(ranges)
		locations = append(locations, minLocation)
	}
	return locations
}

func findLocations(seeds []int, mappings mappings) (locations []int) {
	for _, seedVal := range seeds {
		for _, mapping := range mappings {
			for _, mapFilter := range mapping {
				dst, src, sz := mapFilter[0], mapFilter[1], mapFilter[2]
				if seedVal >= src && seedVal < src+sz {
					seedVal = seedVal - src + dst
					break
				}
			}
		}
		locations = append(locations, seedVal)
	}
	return
}

func findMinLocationRanges(locations [][2]int) int {
	min := locations[0][0]
	for _, l := range locations {
		if l[0] < min {
			min = l[0]
		}
	}
	return min
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(input string) {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	seeds := []int{}
	mappings := mappings{}
	for scanner.Scan() {
		if len(seeds) == 0 {
			seeds = parseSeeds(scanner.Text())
		} else {
			mappings = buildMapFromLine(mappings, scanner.Text())
		}
	}
	locations := findLocations(seeds, mappings)
	minLocation := findMinLocation(locations)
	fmt.Println("Part 1 - Min location:", minLocation)

	seedRanges := parseSeedRanges(seeds)
	locations = findLocationsByRange(seedRanges, mappings)
	minLocation = findMinLocation(locations)
	fmt.Println("Part 2 - Min location:", minLocation)
}

func main() {
	inputFlag := flag.String("input", "day05/input.txt", "Path to input file")
	flag.Parse()

	start := time.Now()
	solve(*inputFlag)
	duration := time.Since(start)
	fmt.Printf("Execution time: %v\n", duration)
}
