package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func caculateTotalPoints(input string) {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	totalPoints := 0
	for scanner.Scan() {
		points := calculateLinePoints(scanner.Text())
		totalPoints += points
	}
	fmt.Println("Total points:", totalPoints)
}

func calculateLinePoints(line string) int {
	tmp := strings.Split(line, ": ")
	tmp = strings.Split(tmp[1], " | ")
	winningNums := strings.Fields(tmp[0])
	myNums := strings.Fields(tmp[1])

	points := 0
	for _, n := range myNums {
		if slices.Contains(winningNums, n) {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}
	return points
}

func main() {
	inputFlag := flag.String("input", "./day04/input.txt", "Path to input file")
	flag.Parse()

	start := time.Now()
	caculateTotalPoints(*inputFlag)
	duration := time.Since(start)
	fmt.Printf("Execution time: %v\n", duration)
}
