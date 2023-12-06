package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func calculateWinningPossibilities(time, distance int) int {
	// distance = x(time-x) ==> x^2 - time*x + distance = 0

	t, d := float64(time), float64(distance)
	discriminant := t*t - 4*d
	if discriminant > 0 {
		root1 := (t - math.Sqrt(discriminant)) / 2
		root2 := (t + math.Sqrt(discriminant)) / 2

		lowerBound := int(math.Ceil(root1))
		upperBound := int(math.Floor(root2))
		// if the bounds are integers we need to adjust
		// we want > distance not >=
		if math.Floor(root1) == root1 {
			lowerBound++
		}
		if math.Floor(root2) == root2 {
			upperBound--
		}
		return upperBound - lowerBound + 1
	} else if discriminant == 0 {
		return 1
	} else {
		return 0
	}
}

func calculateAllWinningPossibilities(times, distance []int) int {
	results := 1
	for i := range times {
		res := calculateWinningPossibilities(times[i], distance[i])
		results *= res
	}
	return results
}

func parseInputLine(line string) ([]int, []string) {
	strs := strings.Fields(line)[1:]
	ints := make([]int, len(strs))
	for i := range strs {
		intVal, _ := strconv.Atoi(strs[i])
		ints[i] = intVal
	}
	return ints, strs
}

func solve(input string) {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines [][]int
	var linesStrings [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numLine, stringLine := parseInputLine(scanner.Text())
		lines = append(lines, numLine)
		linesStrings = append(linesStrings, stringLine)
	}
	times := lines[0]
	distances := lines[1]
	result := calculateAllWinningPossibilities(times, distances)
	fmt.Println("The result for part 1:", result)

	timeStr := strings.Join(linesStrings[0], "")
	distanceStr := strings.Join(linesStrings[1], "")
	time, _ := strconv.Atoi(timeStr)
	distance, _ := strconv.Atoi(distanceStr)
	result = calculateWinningPossibilities(time, distance)
	fmt.Println("The result for part 2:", result)
}

func main() {
	inputFlag := flag.String("input", "day06/input.txt", "Path to input file")
	flag.Parse()

	start := time.Now()
	solve(*inputFlag)
	duration := time.Since(start)
	fmt.Printf("Execution time: %v\n", duration)
}
