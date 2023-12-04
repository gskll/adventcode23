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
	symbol struct {
		value byte
		index int
		line  int
	}
	symbols []symbol
	number  struct {
		value string
		start int
		end   int
		used  bool
	}
	numbers []number
	gears   map[string][]int
)

func (n numbers) values() []string {
	var values []string
	for _, num := range n {
		values = append(values, num.value)
	}
	return values
}

func (s symbols) gears() symbols {
	var gears symbols
	for _, symbol := range s {
		if symbol.value == '*' {
			gears = append(gears, symbol)
		}
	}
	return gears
}

func parseLine(
	line, prevLine string,
	prevLineNumbers numbers,
	potentialGears gears,
	lineCount int,
) (int, numbers, gears) {
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
			adjacentSymbols := adjacentSymbols(numStart, numEnd, line, prevLine, lineCount)
			if len(adjacentSymbols) > 0 {
				numVal, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}
				sum += numVal

				gears := adjacentSymbols.gears()
				if len(gears) > 0 {
					for _, sym := range gears {
						gearKey := fmt.Sprintf("%d-%d", sym.line, sym.index)
						potentialGears[gearKey] = append(potentialGears[gearKey], numVal)
					}
				}
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
	return sum, potentialNumbers, potentialGears
}

func isSymbol(b byte) bool {
	r := rune(b)
	return r != '.' && !unicode.IsDigit(r)
}

func adjacentSymbols(start, end int, currLine, prevLine string, lineCount int) symbols {
	symbols := symbols{}
	for i := start; i <= end; i++ {
		if i == start && i > 0 && isSymbol(currLine[i-1]) {
			symbols = append(symbols, symbol{currLine[i-1], i - 1, lineCount})
		}
		if i == end && i < len(currLine)-1 && isSymbol(currLine[i+1]) {
			symbols = append(symbols, symbol{currLine[i+1], i + 1, lineCount})
		}
		if prevLine != "" {
			if isSymbol(prevLine[i]) {
				symbols = append(symbols, symbol{prevLine[i], i, lineCount - 1})
			}
			if i == start && i > 0 && isSymbol(prevLine[i-1]) {
				symbols = append(symbols, symbol{prevLine[i-1], i - 1, lineCount - 1})
			}
			if i == end && i < len(prevLine)-1 && isSymbol(prevLine[i+1]) {
				symbols = append(symbols, symbol{prevLine[i+1], i + 1, lineCount - 1})
			}
		}
	}

	return symbols
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

	lineCount := 0
	totalSum := 0
	prevLine, currLine := "", ""
	prevLineNumbers := numbers{}  // numbers that have no matching symbols on the current/prev line but may on the next line
	prevPotentialGears := gears{} // all '*' symbols with their adjacent numbers
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prevLine = currLine
		currLine = scanner.Text()
		sum, potentialNumbers, potentialGears := parseLine(
			currLine,
			prevLine,
			prevLineNumbers,
			prevPotentialGears,
			lineCount,
		)
		prevPotentialGears = potentialGears
		prevLineNumbers = potentialNumbers
		totalSum += sum
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	gearSum := 0
	for _, nums := range prevPotentialGears {
		if len(nums) == 2 {
			gearSum += nums[0] * nums[1]
		}
	}

	fmt.Println("Part 1 sum:", totalSum, "("+input+")")
	fmt.Println("Part 2 sum:", gearSum, "("+input+")")
}
