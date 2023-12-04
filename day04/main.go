package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
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
	copies := map[int]int{}
	for scanner.Scan() {
		id, points, matches := calculateLine(scanner.Text())
		totalPoints += points

		copies[id]++

		for i := id + 1; i <= id+matches; i++ {
			copies[i] += copies[id]
		}
	}

	totalCards := 0
	for _, v := range copies {
		totalCards += v
	}
	fmt.Println("Total points:", totalPoints)
	fmt.Println("Total cards:", totalCards)
}

func parseCardId(cardId string) int {
	parts := strings.Fields(cardId)
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return id
}

func calculateLine(line string) (int, int, int) {
	tmp := strings.Split(line, ": ")
	id := parseCardId(tmp[0])
	nums := strings.Split(tmp[1], " | ")
	winningNums := strings.Fields(nums[0])
	myNums := strings.Fields(nums[1])

	points := 0
	matches := 0
	for _, n := range myNums {
		if slices.Contains(winningNums, n) {
			matches++
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}
	return id, points, matches
}

func main() {
	inputFlag := flag.String("input", "./day04/input.txt", "Path to input file")
	flag.Parse()

	start := time.Now()
	caculateTotalPoints(*inputFlag)
	duration := time.Since(start)
	fmt.Printf("Execution time: %v\n", duration)
}
