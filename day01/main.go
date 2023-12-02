package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func parseWordDigits(line string) string {
	digitWords := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}
	for digitWord, digitValue := range digitWords {
		if strings.Contains(line, digitWord) {
			line = strings.ReplaceAll(line, digitWord, digitValue)
		}
	}
	return line
}

func findDigits(part, line string) (firstDigit, lastDigit int) {
	firstDigit, lastDigit = -1, -1

	if part == "part2" {
		line = parseWordDigits(line)
		fmt.Println(line)
	}

	for _, r := range line {
		if unicode.IsDigit(r) {
			digit := int(r - '0')
			if firstDigit == -1 {
				firstDigit = digit
			}
			lastDigit = digit
		}
	}
	return
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	sum := 0
	sum2 := 0
	for scanner.Scan() {
		firstDigit, lastDigit := findDigits("part1", scanner.Text())
		firstDigit2, lastDigit2 := findDigits("part2", scanner.Text())

		if firstDigit != -1 && lastDigit != -1 {
			calibrationValue, err := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
			if err != nil {
				panic(err)
			}
			sum += calibrationValue
		}
		if firstDigit2 != -1 && lastDigit2 != -1 {
			fmt.Println(firstDigit2, lastDigit2)
			calibrationValue2, err := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit2, lastDigit2))
			if err != nil {
				panic(err)
			}
			sum2 += calibrationValue2
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("The result for part 1: ", sum) // 55834
	fmt.Println("The result for part 2: ", sum2)
}
