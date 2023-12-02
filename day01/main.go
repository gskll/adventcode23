package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func findDigits(line string) (firstDigit, lastDigit int) {
	firstDigit, lastDigit = -1, -1

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
	for scanner.Scan() {
		firstDigit, lastDigit := findDigits(scanner.Text())
		if firstDigit != -1 && lastDigit != -1 {
			calibrationValue, err := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
			if err != nil {
				panic(err)
			}
			sum += calibrationValue
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum) // 55834
}
