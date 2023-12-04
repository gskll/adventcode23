package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
	"unicode"
)

type (
	number struct {
		value string
		start int
		end   int
		used  bool
	}
	numbers []number
)

func (n numbers) values() []string {
	var values []string
	for _, num := range n {
		values = append(values, num.value)
	}
	return values
}

func parseLine(line, prevLine string, prevLineNumbers numbers) (int, numbers) {
	sum := 0
	num := ""
	numStart := 0
	potentialNumbers := numbers{}
	for i, r := range line {
		if unicode.IsDigit(r) {
			if num == "" {
				numStart = i
			}
			num += string(r)
		}

		if num != "" && (!unicode.IsDigit(r) || i == len(line)-1) {
			numEnd := numStart + len(num) - 1
			if hasAdjacentSymbols(numStart, numEnd, line, prevLine) {
				numVal, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}
				sum += numVal
			} else {
				number := number{num, numStart, numEnd, false}
				potentialNumbers = append(potentialNumbers, number)
			}
			num = ""
		}

		if isSymbol(byte(r)) && len(prevLineNumbers) > 0 {
			for _, number := range prevLineNumbers {
				if !number.used && i >= number.start-1 && i <= number.end+1 {
					numVal, err := strconv.Atoi(number.value)
					if err != nil {
						panic(err)
					}
					sum += numVal
					number.used = true
				}
			}
		}
	}
	return sum, potentialNumbers
}

func isSymbol(b byte) bool {
	r := rune(b)
	return r != '.' && !unicode.IsDigit(r)
}

func checkLeft(i int, line string) bool {
	return i > 0 && isSymbol(line[i-1])
}

func checkRight(i int, line string) bool {
	return i < len(line)-1 && isSymbol(line[i+1])
}

func hasAdjacentSymbols(start, end int, currLine, prevLine string) bool {
	for i := start; i <= end; i++ {
		if i == start && checkLeft(i, currLine) {
			return true
		}
		if i == end && checkRight(i, currLine) {
			return true
		}
		if prevLine != "" {
			if isSymbol(prevLine[i]) ||
				(i == start && checkLeft(i, prevLine)) ||
				(i == end && checkRight(i, prevLine)) {
				return true
			}
		}
	}
	return false
}

func main() {
	inputFlag := flag.String("input", "./day03/input.txt", "Path to input file")
	flag.Parse()

	start := time.Now()
	calculateSum(*inputFlag) // 537832
	duration := time.Since(start)
	fmt.Printf("Execution time: %v\n", duration)
}

func calculateSum(input string) {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	totalSum := 0
	prevLine, currLine := "", ""
	prevLineNumbers := numbers{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prevLine = currLine
		currLine = scanner.Text()
		sum, potentialNumbers := parseLine(currLine, prevLine, prevLineNumbers)
		prevLineNumbers = potentialNumbers
		totalSum += sum
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Part 1 sum:", totalSum, "("+input+")")
}
