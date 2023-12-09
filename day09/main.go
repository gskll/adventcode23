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

func parseSequence(line string) (seq []int) {
	for _, numString := range strings.Fields(line) {
		num, _ := strconv.Atoi(numString)
		seq = append(seq, num)
	}
	return
}

func predictNextValue(line string) (nextValue, prevValue int) {
	seq := parseSequence(line)
	counter, allZeroes := 1, false

	var firstValues []int
	for !allZeroes {
		allZeroes = true
		firstValues = append(firstValues, seq[0])
		for i := range seq[:len(seq)-counter] {
			seq[i] = seq[i+1] - seq[i]
			if seq[i] != 0 {
				allZeroes = false
			}

		}
		counter++
	}
	for _, val := range seq {
		nextValue += val
	}
	for i := len(firstValues) - 1; i >= 0; i-- {
		prevValue = firstValues[i] - prevValue
	}
	return
}

func solve(input string) {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var nextSum, prevSum int
	for scanner.Scan() {
		nextValue, prevValue := predictNextValue(scanner.Text())
		nextSum += nextValue
		prevSum += prevValue
	}
	fmt.Println("part 1 sum:", nextSum)
	fmt.Println("part 2 sum:", prevSum)
}

func main() {
	inputFlag := flag.String("input", "input.txt", "path to input file")
	flag.Parse()

	start := time.Now()
	solve(*inputFlag)
	duration := time.Since(start)
	fmt.Printf("Execution time: %v\n", duration)
}
